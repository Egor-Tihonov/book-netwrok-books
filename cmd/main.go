package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"

	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/config"
	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/handlers"
	pb "github.com/Egor-Tihonov/book-netwrok-books.git/pkg/pb"
	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/repository"
	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	InitLog()
	
	c, err := config.LoadConfig()

	if err != nil {
		logrus.Fatalf("book service: error load configs: %w", err)
	}

	dbP, err := repository.New(c.DBUrl)
	if err != nil {
		logrus.Fatalf("book service: error connecting to db, %w", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		logrus.Fatalln("book service: Failed to listing:", err)
	}

	logrus.Info("------ START SERVER ON ", c.Port, " ------")

	s := service.New(dbP)
	h := handlers.New(s)

	grpcServer := grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Book service: Failed to serve:", err)
	}
}

func InitLog() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	})

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
