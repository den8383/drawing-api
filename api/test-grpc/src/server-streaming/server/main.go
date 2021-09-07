package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-sample/pb/notification"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"


type ServerServerSide struct {
  pb.UnimplementedNotificationServer
}

func (s *ServerServerSide) Notification(req *pb.NotificationRequest, stream pb.Notification_NotificationServer) error {
  fmt.Println("get request")
  for i := int32(0); i < req.GetNum(); i++ {
    message := fmt.Sprintf("%d", i)
    if err := stream.Send(&pb.NotificationReply{
      Message: message,
    }); err != nil{
      return err
    }
    time.Sleep(time.Second)
  }
  return nil
}

func set() error {
  lis,err := net.Listen("tcp", port)
  if err != nil {
    return errors.Wrap(err, "port failed")
  }
  s := grpc.NewServer()
  var server ServerServerSide
  pb.RegisterNotificationServer(s, &server)
  if err := s.Serve(lis); err != nil {
    return errors.Wrap(err, "server boot failed")
  }
  return nil
}

func main(){
  fmt.Println("boot")
  if err := set(); err != nil {
    log.Fatal("%v", err)
  }
}
