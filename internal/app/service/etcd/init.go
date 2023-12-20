package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdClient struct {
	Cli      *clientv3.Client
	KVCli    clientv3.KV
	LeaseCli clientv3.Lease   // 租约句柄
	LeaseID  clientv3.LeaseID // 租约ID
}

var etcdLocalMap = []string{"127.0.0.1:2379"}

var dialTimeout = 5 * time.Second

// NewClient 创建 etcd client 句柄
func NewClient() EtcdClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            etcdLocalMap, // etcd 的多个节点服务地址
		AutoSyncInterval:     0,            // 对应设置的 etcd 节点服务最新更新的时间间隔，0禁用自动同步。默认情况下，自动同步处于禁用状态。
		DialTimeout:          dialTimeout,  // 创建 client 的首次连接超时时间，这里传了 5 秒，如果 5 秒都没有连接成功就会返回 err；一旦 client 创建成功，我们就不用再关心后续底层连接的状态了，client 内部会重连
		DialKeepAliveTime:    0,            // 是客户端 ping 服务器以查看传输是否处于活动状态的时间。
		DialKeepAliveTimeout: dialTimeout,  // 是客户端等待保持活动探测的响应的时间。如果此时未收到响应，则连接将关闭。
		MaxCallSendMsgSize:   0,            // 是以字节为单位的客户端请求发送限制。如果为0，则默认为2.0 MiB（2*1024*1024）确保“MaxCallSendMsgSize” 大于服务器端默认发送/接收限制。（“--max request bytes” 标记为 etcd 或 “embed.Config.MaxRequestBytes”）
		MaxCallRecvMsgSize:   0,            // 是客户端响应接收限制。如果为0，则默认为 “math.MaxInt32”，因为范围响应很容易超过请求发送限制。确保“MaxCallRecvMsgSize” >= 服务器端默认发送/接收限制。（“--max request bytes”标记为etcd或“embed.Config.MaxRequestBytes”）
		TLS:                  nil,          // 设置客户端安全凭据（如果有的话）。
		Username:             "",           // 是用于身份验证的用户名。
		Password:             "",           // 是用于身份验证的密码。
		RejectOldCluster:     false,        // 当设置时，将拒绝针对过时的集群创建客户端。默认为不拒绝
		DialOptions:          nil,          // 是grpc客户端的拨号选项列表（例如，用于拦截器）。 例如，将“grpc.WithBlock（）”传递到block，直到底层连接启动。 否则，Dial会立即返回，并在后台连接服务器。
		Context:              nil,          // 是默认的客户端上下文；它可以用于取消grpc拨出和其他没有显式上下文的操作。
		Logger:               nil,          // 设置客户端日志，如果为空，默认使用 LogConfig
		LogConfig:            nil,          // 配置客户端日志，如果为零，则使用默认记录器
		PermitWithoutStream:  false,        // 设置时将允许客户端在没有任何活动流（RPC）的情况下向服务器发送保持活动的ping。
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("[etcd] connect %#v to success... \r\n", etcdLocalMap)

	return EtcdClient{
		Cli:      cli,
		KVCli:    clientv3.NewKV(cli),
		LeaseCli: clientv3.NewLease(cli),
	}
}
