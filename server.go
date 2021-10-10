// @Time    : 2021/10/9 7:22 下午
// @Author  : HuYuan
// @File    : main.go
// @Email   : huyuan@virtaitech.com

package main

import (
	"github.com/mdlayher/vsock"
	"io"
	"log"
)

// 服务器端只能运行在宿主机上
func main()  {
	// 这个 Listen 的端口是 vsock 的端口, 和 tcp 或者 udp 这些端口没关系. 即便端口相同也不会冲突
	l, err := vsock.Listen(1024)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()

		// 将客户端发送的消息回显给客户端(即vm)
		if _, err := io.Copy(c, c); err != nil {
			log.Fatal(err)
		}
	}
}
