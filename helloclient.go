package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"hello"
)

func handleClient(client hello.Hello) error {
	str, err := client.HelloString("你好呀，Hello moto")
	fmt.Println(str)
	return err
}

func main() {
	/**
	  与服务端同理

	  1. 选择协议分层中的  数据封装协议、传输协议、
	*/
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	addr := "localhost:8888"

	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		panic(err)
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()

	/**
	  2. 建立TCP连接
	*/
	if err := transport.Open(); err != nil {
		panic(err)
	}

	/**
	  3. 构造服务的引用
	*/
	service := hello.NewHelloClientFactory(transport, protocolFactory)

	/**
	  4. 根据自己的业务，调用所需的服务
	*/
	str, _ := service.HelloString("你好呀，Hello moto")
	fmt.Println(str)
}
