package client

import (
	"crypto/tls"
	pb "github.com/je4/mediaserverproto/v2/pkg/mediaserver/proto"
	"google.golang.org/grpc"
	"io"
)

func NewActionDispatcherClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.ActionDispatcherClient, io.Closer, error) {
	return newClient[pb.ActionDispatcherClient](pb.NewActionDispatcherClient, serverAddr, tlsConfig, opts...)
}

func NewActionClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.ActionClient, io.Closer, error) {
	return newClient[pb.ActionClient](pb.NewActionClient, serverAddr, tlsConfig, opts...)
}
