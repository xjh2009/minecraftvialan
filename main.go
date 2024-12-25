package main

import (
    "net"
    "os"
    "time"
)

func main() {
    // 检查是否提供了参数
	if len(os.Args) < 3 {
        println("你需要提供端口号和隧道名称")
        os.Exit(1)
    }

    // 获取端口号
    port := os.Args[2]
    pname := os.Args[1]

    // 设置广播地址和端口
    addr := net.UDPAddr{
        IP:   net.ParseIP("224.0.2.60"),
        Port: 4445,
    }

    // 创建UDP连接
    conn, err := net.DialUDP("udp", nil, &addr)
    if err != nil {
        os.Exit(1)
    }
    defer conn.Close()

    // 广播内容
    message := "[MOTD]默连局域网 - "+pname+"[/MOTD][AD]" + port + "[/AD]"

    // 每隔1.5秒发送一次消息
    ticker := time.NewTicker(1500 * time.Millisecond) // 使用毫秒表示
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            _, err = conn.Write([]byte(message))
            if err != nil {
                os.Exit(1)
            }
            // 输出发送的内容（可选）
            //println("广播发送:", message)
        }
    }
}
