package client

import (
	"crypto/tls"
	pb "github.com/je4/mediaserverproto/v2/pkg/mediaserver/proto"
	"google.golang.org/grpc"
	"io"
)

func NewDeleterClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.DeleterClient, io.Closer, error) {
	return newClient[pb.DeleterClient](pb.NewDeleterClient, serverAddr, tlsConfig, opts...)
}
