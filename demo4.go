package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 15:51
* @Description: 获取目录结构下的所有字符串
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

	if getResp, err := kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(getResp.Kvs)
	}
}