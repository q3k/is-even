package main

import (
	"context"
	"testing"

	pb "github.com/q3k/is-even/proto/is-even"
	opb "github.com/q3k/is-even/proto/is-odd"
	"google.golang.org/grpc"
)

type oddFake struct{}

func (o *oddFake) IsOdd(ctx context.Context, req *opb.IsOddRequest, _ ...grpc.CallOption) (*opb.IsOddResponse, error) {
	res := &opb.IsOddResponse{
		Result: opb.IsOddResponse_RESULT_NON_ODD,
	}
	if req.Number%2 != 0 {
		res.Result = opb.IsOddResponse_RESULT_ODD
	}
	return res, nil
}

func TestIsEven(t *testing.T) {
	ctx := context.Background()
	s := &server{odd: &oddFake{}}

	tests := []struct {
		number int64
		want   pb.IsEvenResponse_Result
	}{
		{1337, pb.IsEvenResponse_RESULT_NON_EVEN},
		{-1337, pb.IsEvenResponse_RESULT_NON_EVEN},
		{42, pb.IsEvenResponse_RESULT_EVEN},
		{-42, pb.IsEvenResponse_RESULT_EVEN},
		{0, pb.IsEvenResponse_RESULT_EVEN},
	}

	for i, test := range tests {
		req := &pb.IsEvenRequest{
			Number: test.number,
		}
		res, err := s.IsEven(ctx, req)
		if err != nil {
			t.Fatalf("Case %d: error from service: %v", i, err)
		}

		if want, got := test.want, res.Result; want != got {
			t.Errorf("Case %d: want %v, got %v", i, want, got)
		}
	}
}
