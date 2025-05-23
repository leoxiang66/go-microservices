package main

import (
    "time"
    "go-microservices/microservices"
)





func main() {
    // 假设 B 在 50052，A 在 50051
    go microservices.RunServiceB(":50052")
    // 等 B 启动好
    time.Sleep(500 * time.Millisecond)
    go microservices.RunServiceA("localhost:50052", ":50051")

    // 阻塞主线程
    select {}
}
