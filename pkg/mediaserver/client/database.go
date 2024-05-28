package client

import (
	"crypto/tls"
	pb "github.com/je4/mediaserverproto/v2/pkg/mediaserver/proto"
	"google.golang.org/grpc"
	"io"
)

func NewDatabaseClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.DatabaseClient, io.Closer, error) {
	return newClient[pb.DatabaseClient](pb.NewDatabaseClient, serverAddr, tlsConfig, opts...)
}
