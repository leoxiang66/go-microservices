package microservices

import (
    "context"

    "log"
    "net"

    "google.golang.org/grpc"
	pbB "go-microservices/microservices/pb/api/serviceB"
)


// --- 微服务 B 部分 ---

type serviceBServer struct {
    pbB.UnimplementedServiceBServer
}

func (s *serviceBServer) BarB(ctx context.Context, req *pbB.ReqB) (*pbB.RespB, error) {
    // B 只做自己的业务
    return &pbB.RespB{Data: "Hello from B: " + req.Param}, nil
}

func RunServiceB(bPort string) {
    lis, err := net.Listen("tcp", bPort)
    if err != nil {
        log.Fatalf("B Listen 失败: %v", err)
    }
    grpcServer := grpc.NewServer()
    pbB.RegisterServiceBServer(grpcServer, &serviceBServer{})
    log.Printf("Service B listening on %s", bPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("B Serve 失败: %v", err)
    }
}