package microservices

import (
	"context"
	"log"
	"net"
	pbA "go-microservices/microservices/pb/api/serviceA"
	pbB "go-microservices/microservices/pb/api/serviceB"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// --- 微服务 A 部分 ---

type serviceAServer struct {
    pbA.UnimplementedServiceAServer
    bClient pbB.ServiceBClient  // 用于调用 B
}

func (s *serviceAServer) FooA(ctx context.Context, req *pbA.ReqA) (*pbA.RespA, error) {
    // A 处理自身逻辑，然后调用 B 的 BarB
    respB, err := s.bClient.BarB(ctx, &pbB.ReqB{Param: req.Param})
    if err != nil {
        return nil, err
    }
    return &pbA.RespA{Result: "A got: " + respB.Data}, nil
}



func RunServiceA(bAddr string, aPort string) {
    // 1. 建立到 B 的 gRPC 连接（Client）
    connB, err := grpc.NewClient(bAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("A 连接 B 失败: %v", err)
    }
    defer connB.Close()
    bClient := pbB.NewServiceBClient(connB)

    // 2. 在本地启动 gRPC Server 暴露 A 的接口
    lis, err := net.Listen("tcp", aPort)
    if err != nil {
        log.Fatalf("A Listen 失败: %v", err)
    }
    grpcServer := grpc.NewServer()
    pbA.RegisterServiceAServer(grpcServer, &serviceAServer{bClient: bClient})
    log.Printf("Service A listening on %s", aPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("A Serve 失败: %v", err)
    }
}