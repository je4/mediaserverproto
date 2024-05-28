package client

import (
	"crypto/tls"
	"emperror.dev/errors"
	pb "github.com/je4/mediaserverproto/v2/pkg/mediaserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

func NewDatabaseClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.DatabaseClient, io.Closer, error) {
	if tlsConfig != nil {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot connect to %s", serverAddr)
	}

	client := pb.NewDatabaseClient(conn)
	return client, conn, nil
}
