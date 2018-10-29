package impl

import (
	"context"
	"flag"
	"testing"
	"time"

	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = ":50052"
)

func Test_body_client(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()
	c := pb.NewBodyTrackClient(conn)
	req := &pb.TrackOption{CaseId: time.Now().Unix(),
		CaseInit: true,
		Image:    "/home/wanglei/Pictures/obm.jpg",
		Dynamic:  false}
	res, err := c.Track(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	logrus.Println(res)
}
