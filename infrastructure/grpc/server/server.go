package server

import (
	"context"
	"go-challenger/infrastructure/grpc/functions"
	"go-challenger/infrastructure/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	funcGrpc functions.CRUDServiceServer
	port     string
	timeout  time.Duration
}

func NewServer(funcGrpc functions.CRUDServiceServer, port string, timeout time.Duration) Server {
	return Server{funcGrpc: funcGrpc, timeout: timeout, port: port}
}

func (s *Server) Listen(ctx context.Context, wg *sync.WaitGroup) {

	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}

	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,
		MaxConnectionAge:      s.timeout,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               1 * time.Second,
	}

	g := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))

	hs := health.NewServer()
	hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	healthpb.RegisterHealthServer(g, hs)
	functions.RegisterCRUDServiceServer(g, s.funcGrpc)

	logger.Infof("Server started at :" + s.port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := g.Serve(lis); err != nil {
				logger.Fatal("Failed to serve: %v", err)
			}
		}()


	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		logger.Infof("Graceful shutdown...")
		g.GracefulStop()
		logger.Infof("Exiting server")
	}()

}
