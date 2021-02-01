package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 21:15
* @Description: 租约操作
**/

func main() {
	var(
		config clientv3.Config
		client *clientv3.Client
		err error
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		kv clientv3.KV
	)

	config = clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5 * time.Second,
	}

	//创建客户端
	if client, err = clientv3.New(config); err != nil{
		fmt.Println(err)
		return
	}

	lease = clientv3.NewLease(client)
	// 申请续约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil{
		fmt.Println(err)
		return
	}
	leaseId = leaseGrantResp.ID

	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil{
		fmt.Println(err)
		return
	}
	go func() {
		for {
			select {
			case keepResp = <- keepRespChan:
				if keepResp == nil{
					fmt.Println("租约失效")
					return
				}else{
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}
			}
		}
	}()

	kv = clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时的看一下key过期了没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}