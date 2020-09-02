package etcd

/**
 * Copyright (C) 2007-2020 56X.NET,All rights reserved.
 *
 * name : etcd_registry
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2020-09-02 15:38
 * description :
 * history :
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"hash/crc32"
	"log"
	"time"
)

var prefix = "/registry/server/"

type Registry interface {
	// 创建租期/注册服务,返回租期ID和错误
	Register(addr string)(int64,error)
	// 撤销租期/注销服务
	Revoke(LeaseID int64)error
	UnRegister()
}
var _ Registry = new(registryServer)
type registryServer struct {
	cli        *clientv3.Client
	stop       chan bool
	isRegistry bool
	leaseID    clientv3.LeaseID // 租约ID
	service    string
	ttl        int64 // 租约时间
}
type Node struct {
	Id   uint32 `json:"id"`
	Addr string `json:"addr"`
}

// 创建服务注册, ttl租约时间
func NewRegistry(service string, ttl int64, config clientv3.Config) (Registry, error) {
	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	return &registryServer{
		service:    service,
		ttl:        ttl,
		stop:       make(chan bool),
		isRegistry: false,
		cli:        cli,
	}, nil
}
func (s *registryServer) Register(addr string)(leaseId int64, err error){
	if s.isRegistry {
		panic("only one node can be registered")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.ttl)*time.Second)
	defer cancel()
	// 创建租约
	grant, err := s.cli.Grant(context.Background(), s.ttl)
	if err != nil {
		return -1,err
	}
	var node = Node{
		Id:   s.HashKey(addr),
		Addr: addr,
	}
	nodeVal, err := s.GetVal(node)
	if err != nil {
		return -1,err
	}
	//　存储键值,注册服务
	_, err = s.cli.Put(ctx, s.GetKey(node), nodeVal, clientv3.WithLease(grant.ID))
	if err != nil {
		return -1,err
	}
	s.leaseID = grant.ID
	s.isRegistry = true
	go s.KeepAlive()
	return int64(s.leaseID),nil
}
func (s *registryServer) UnRegister() {
	if s.isRegistry{
		s.stop <- true
	}
}

// 注销服务
func (s *registryServer) Revoke(leaseId int64) error {
	return s.revoke(clientv3.LeaseID(leaseId))
}

// 注销服务
func (s *registryServer) revoke(leaseID clientv3.LeaseID) error {
	// 撤销租期
	_, err := s.cli.Revoke(context.TODO(),leaseID)
	if err != nil {
		log.Printf("[Revoke] err : %s", err.Error())
	}
	s.isRegistry = false
	return err
}

// 监听服务
func (s *registryServer) KeepAlive() error {
	keepAliveCh, err := s.cli.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[KeepAlive] err : %s", err.Error())
		return err
	}
	for {
		select {
		case <-s.stop:
			_ = s.revoke(s.leaseID)
			return nil
		case _, ok := <-keepAliveCh:
			if !ok {
				_ = s.revoke(s.leaseID)
				return nil
			}
		}
	}
}
func (s *registryServer) GetKey(node Node) string {
	return fmt.Sprintf("%s%s/%d", prefix, s.service, s.HashKey(node.Addr))
}
func (s *registryServer) GetVal(node Node) (string, error) {
	data, err := json.Marshal(&node)
	return string(data), err
}
func (s *registryServer) HashKey(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}
