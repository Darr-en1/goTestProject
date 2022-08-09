package trip

import (
	"context"
	trippb "goTestProject/grpc-demo/gen/go"
)

// Service gPRC 生成 go 代码为什么会有 UnimplementedServer 和 mustEmbedUnimplementedServer？
// https://blog.csdn.net/Canon_in_D_Major/article/details/108135724
type Service struct {
	*trippb.UnimplementedTripServiceServer
}

func (s Service) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:   200,
				Longtitude: 300,
			},
			Status: trippb.TripStatus_NOT_STARTED,
		},
	}, nil
}
