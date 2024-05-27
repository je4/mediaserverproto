package client

import (
	"crypto/tls"
	"emperror.dev/errors"
	pb "github.com/je4/mediaserverproto/v2/pkg/mediaserverdeleter/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

func CreateClient(serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (pb.DeleterControllerClient, io.Closer, error) {
	if tlsConfig != nil {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.NewClient(serverAddr, opts...)
	//conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot connect to %s", serverAddr)
	}

	client := pb.NewDeleterControllerClient(conn)
	return client, conn, nil
}
