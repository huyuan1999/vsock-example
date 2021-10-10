// @Time    : 2021/10/9 8:59 下午
// @Author  : HuYuan
// @File    : client.go
// @Email   : huyuan@virtaitech.com

package main

import (
	"bufio"
	"fmt"
	"github.com/mdlayher/vsock"
	"log"
	"os"
)



// 客户端只能运行在虚拟机上
func main()  {
	// 这个 Listen 的端口是 vsock 的端口, 和 tcp 或者 udp 这些端口没关系. 即便端口相同也不会冲突
	c, err := vsock.Dial(vsock.Host, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		fmt.Printf("请输入你想给服务器发送的消息: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("读取输入错误: ", err.Error())
		}
		// 发送消息给服务器端
		if _, err := c.Write([]byte(input)); err != nil {
			log.Fatal(err)
		}

		// 接受服务端返回的消息
		b := make([]byte, 16)
		n, err := c.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b[:n]))
	}
}
