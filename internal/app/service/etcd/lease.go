package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	keepTime   = 100 * time.Second // 续租时长
	cancelFunc func()              // 租约撤销上下文回调

	err           error
	leaseResp     *clientv3.LeaseGrantResponse
	keepaliveResp *clientv3.LeaseKeepAliveResponse
	keepRespChan  <-chan *clientv3.LeaseKeepAliveResponse
)

func (e *EtcdClient) LeasesList() (*clientv3.LeaseLeasesResponse, error) {
	return e.Cli.Leases(e.Cli.Ctx())
}

// ApplyLease 申请租约
func (e *EtcdClient) ApplyLease(ttl int64) error {
	// 设置租约
	leaseResp, err = e.LeaseCli.Grant(e.Cli.Ctx(), ttl)
	if err != nil {
		return err
	}
	e.LeaseID = leaseResp.ID
	// 设置续租，定期发送续租请求
	ctx, cancel := context.WithCancel(e.Cli.Ctx())
	cancelFunc = cancel
	keepRespChan, err = e.LeaseCli.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}
	//监听续租相应chan
	go e.listenLeaseRespChan()
	return nil
}

// ListenLeaseRespChan 监听续租情况
func (e *EtcdClient) listenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-keepRespChan:
			// 如果关闭续租或续租到期，则退出监听
			if leaseKeepResp == nil {
				goto END
				return
			}
		}
	}
END:
}

// RevokeLease 撤销租约
func (e *EtcdClient) RevokeLease() error {
	cancelFunc()
	time.Sleep(2 * time.Second)
	_, err = e.LeaseCli.Revoke(e.Cli.Ctx(), leaseResp.ID)
	return err
}
