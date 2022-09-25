package grpc

import (
	"cpfd-back/internal/core/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataGen struct {
	*conf
}

func New() *DataGen {
	dg := DataGen{}
	dg.conf = newConf()

	return &dg
}

func (dg *DataGen) Request(req *pb.LogRequest) (*pb.LogResponse, error) {
	conn, err := dg.connToGrpc()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	client := pb.NewDataGeneratorClient(conn)
	return client.GenerateAllLog(dg.ctx, req)
}

func (dg *DataGen) connToGrpc() (*grpc.ClientConn, error) {
	return grpc.Dial(dg.getDSN(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
