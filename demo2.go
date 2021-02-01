package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 15:28
* @Description: 写入kv
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
	if putResp, err := kv.Put(context.TODO(), "/cron/jobs/job1", "hello1"); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("revision", putResp.Header.Revision)
	}
	//获取之前revision的值
	if putResp, err := kv.Put(context.TODO(), "/cron/jobs/job1", "hello2", clientv3.WithPrevKV()); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("revision", putResp.Header.Revision)
		if putResp.PrevKv != nil{
			fmt.Println("prevValue", string(putResp.PrevKv.Value))
		}
	}

}