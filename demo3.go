package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 15:44
* @Description: 读取kv
**/

func main() {
	var(
		config clientv3.Config
		client *clientv3.Client
		err error
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

	// 用于操作etcdkv
	kv = clientv3.NewKV(client)

	if getResp, err := kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(getResp.Kvs)
	}

	// 只返回count
	if getResp, err := kv.Get(context.TODO(), "/cron/jobs/job1", clientv3.WithCountOnly()); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(getResp.Count)
	}
}