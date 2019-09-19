package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"

	"github.com/gogo/status"
	pb "github.com/q3k/is-even/proto/is-even"
	opb "github.com/q3k/is-even/proto/is-odd"
)

var (
	flagListen string
	flagOdd    string
)

type server struct {
	odd opb.IsOddClient
}

func main() {
	flag.StringVar(&flagListen, "listen", "0.0.0.0:2138", "Address to listen at")
	flag.StringVar(&flagOdd, "odd", "127.0.0.1:2137", "Address of odd microservice")
	flag.Parse()

	lis, err := net.Listen("tcp", flagListen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := grpc.Dial(flagOdd, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial odd: %v", err)
	}
	defer conn.Close()

	odd := opb.NewIsOddClient(conn)

	srv := &server{
		odd: odd,
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterIsEvenServer(s, srv)

	log.Printf("listening at %v", flagListen)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) IsEven(ctx context.Context, req *pb.IsEvenRequest) (*pb.IsEvenResponse, error) {
	req2 := &opb.IsOddRequest{
		Number: req.Number,
	}

	res2, err := s.odd.IsOdd(ctx, req2)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "odd service: %v", err)
	}

	res := &pb.IsEvenResponse{}
	switch res2.Result {
	case opb.IsOddResponse_RESULT_ODD:
		res.Result = pb.IsEvenResponse_RESULT_NON_EVEN
	case opb.IsOddResponse_RESULT_NON_ODD:
		res.Result = pb.IsEvenResponse_RESULT_EVEN
	}

	return res, nil
}
