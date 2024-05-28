package client

import (
	"crypto/tls"
	"emperror.dev/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

func newClient[V any](newClient func(cc grpc.ClientConnInterface) V, serverAddr string, tlsConfig *tls.Config, opts ...grpc.DialOption) (V, io.Closer, error) {
	if tlsConfig != nil {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		var n V
		return n, nil, errors.Wrapf(err, "cannot connect to %s", serverAddr)
	}

	client := newClient(conn)
	return client, conn, nil
}
