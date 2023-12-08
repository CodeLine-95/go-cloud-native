package test

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/etcd"
	"testing"
)

func TestEtcd(t *testing.T) {
	etcdClient := etcd.NewClient()

	//// 申请租约
	//err := etcdClient.ApplyLease(60)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(etcdClient.LeaseID)

	// 查看租约
	leaseList, err := etcdClient.LeasesList()
	if err != nil {
		panic(err)
	}
	fmt.Println(leaseList.Leases)

	//err := etcdClient.PutService("test2", "127.0.0.1:4000")
	//if err != nil {
	//	panic(err)
	//}
	//
	//list, err := etcdClient.GetService("test")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(list)
	//
	//list2, err := etcdClient.GetService("test2")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(list2)
}
