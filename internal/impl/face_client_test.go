package impl

import (
	"context"
	"flag"
	"testing"

	"github.com/funxdata/baidu/internal/biz"
	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func Test_face_detect(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	// 连接
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()
	c := pb.NewFaceClient(conn)
	req := &pb.DetectOption{Image: "/home/wanglei/Pictures/c.jpg",
		ImageType: biz.FaceImageTypeBase64}
	res, err := c.Detect(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	logrus.Println(res)
}
