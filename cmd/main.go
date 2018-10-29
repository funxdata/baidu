package main

import (
	"net"

	"github.com/funxdata/baidu/internal/impl"
	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	debug   bool
	rootCmd = cobra.Command{
		Use: "baidu-cli",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			listen, err := net.Listen("tcp", ":50052")
			if err != nil {
				logrus.Info(err)
				return
			}
			s := grpc.NewServer()
			bs := impl.NewBodyTrackServer(" ", " ")
			fs := impl.NewFaceServer(" ", " ")
			pb.RegisterBodyTrackServer(s, bs)
			pb.RegisterFaceServer(s, fs)
			s.Serve(listen)
		},
	}
)

func main() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "log debug message.")

	rootCmd.Execute()
}
