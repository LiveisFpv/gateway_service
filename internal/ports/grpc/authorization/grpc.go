package authorization

import (
	"fmt"
	"gateway_service/internal/ports/grpc/interseptors"
	"time"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api sso.AuthClient
}

// func for create new connections
func New(
	log *logrus.Logger,
	address string, //adress to connect auth_service
	timeout time.Duration, //time for try
	retriesCount int, //retry count
) (*Client, error) {
	const op = "grpc.New"
	//Interseptor options
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	//Logging options for interseptor grpclog
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(
			grpclog.PayloadReceived,
			grpclog.PayloadSent,
		),
	}
	// TODO check logger, logs NF
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(interseptors.InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//Create grpc-client auth_service
	grpcClient := sso.NewAuthClient(conn)
	return &Client{
		api: grpcClient,
	}, nil
}
