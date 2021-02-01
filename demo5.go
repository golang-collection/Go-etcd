package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 21:09
* @Description: 删除操作
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
	if deleteResp, err := kv.Delete(context.TODO(), "/cron/jobs/job1",clientv3.WithPrevKV()); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(deleteResp.PrevKvs)
		fmt.Println(deleteResp.Deleted)
	}
}