package fitness

import (
	"fmt"
	"gateway_service/internal/ports/grpc/interseptors"
	"time"

	fitness "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api  fitness.FitnessClient
	conn *grpc.ClientConn
}

// Constructor FitnessService
func New(
	log *logrus.Logger,
	address string, //Connection address host:port
	timeout time.Duration, //time alive
	retriesCount int,
) (*Client, error) {
	const op = "FitnessClient"
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
		),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	grpcClient := fitness.NewFitnessClient(conn)
	return &Client{
		api:  grpcClient,
		conn: conn,
	}, nil
}
func (c *Client) Close() error {
	return c.conn.Close()
}
