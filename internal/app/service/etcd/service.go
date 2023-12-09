package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"time"
)

// PutService 注册服务
func (e *EtcdClient) PutService(key, val string) error {
	ctx, cancel := context.WithTimeout(e.Cli.Ctx(), 2*time.Second)
	_, err := e.KVCli.Put(ctx, key, val)
	cancel()
	return err
}

// DelService 删除服务
func (e *EtcdClient) DelService(key string) error {
	ctx, cancel := context.WithTimeout(e.Cli.Ctx(), 2*time.Second)
	_, err := e.KVCli.Delete(ctx, key)
	cancel()
	return err
}

// GetService 服务发现
func (e *EtcdClient) GetService(key string) ([]*mvccpb.KeyValue, error) {
	ctx, cancel := context.WithTimeout(e.Cli.Ctx(), time.Second)
	resp, err := e.KVCli.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) > 0 {
		return resp.Kvs, nil
	}
	return nil, err
}
