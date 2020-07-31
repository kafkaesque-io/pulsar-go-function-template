package pf

import (
	"context"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/kafkaesque-io/pulsar-go-function-template/src/logutil"
	pb "github.com/kafkaesque-io/pulsar-go-function-template/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// instance-control is the control gRPC interface to interact with
// the management Java process

type server struct {
	pb.UnimplementedInstanceControlServer
}

func (s *server) GetFunctionStatus(ctx context.Context, empt *empty.Empty) (*pb.FunctionStatus, error) {

	return &pb.FunctionStatus{}, nil
}

func (s *server) GetAndResetMetrics(ctx context.Context, empt *empty.Empty) (*pb.MetricsData, error) {
	return &pb.MetricsData{}, nil
}
func (s *server) ResetMetrics(ctx context.Context, empt *empty.Empty) (*empty.Empty, error) {
	return nil, nil

}
func (s *server) GetMetrics(ctx context.Context, empt *empty.Empty) (*pb.MetricsData, error) {
	return &pb.MetricsData{}, nil

}
func (s *server) HealthCheck(ctx context.Context, empt *empty.Empty) (*pb.HealthCheckResult, error) {
	return &pb.HealthCheckResult{
		Success: true,
	}, nil
}

func startController(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterInstanceControlServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Fatal(err)
	}

}