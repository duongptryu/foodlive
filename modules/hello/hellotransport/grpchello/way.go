package grpchello

import (
	"context"
	pb "foodlive/gen/proto"
	"foodlive/modules/hello/hellobiz"
)

func (s *ServerHello) Way(ctx context.Context, req *pb.WayRequest) (*pb.WayResponse, error) {
	id := req.GetId()
	biz := hellobiz.NewHelloBiz()

	name, err := biz.WayBiz(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return &pb.WayResponse{
		Name: name,
	}, nil
}
