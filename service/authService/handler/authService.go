package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	authService "authService/proto"
)

type AuthService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *AuthService) Call(ctx context.Context, req *authService.Request, rsp *authService.Response) error {
	log.Info("Received AuthService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AuthService) Stream(ctx context.Context, req *authService.StreamingRequest, stream authService.AuthService_StreamStream) error {
	log.Infof("Received AuthService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&authService.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AuthService) PingPong(ctx context.Context, stream authService.AuthService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&authService.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
