package main

import (
  "fmt"
  "io"
  "log"
  "net"
  "time"

  pb "grpc-sample/pb/chat"
  "github.com/pkg/errors"
  "google.golang.org/grpc"
)

const port = ":50051"

type ServerBidirectional struct {
  pb.UnimplementedChatServer
}

func request(stream pb.Chat_ChatServer, message string) error {
  reply := fmt.Sprintf("received%s", message)
  return stream.Send(&pb.ChatReply{
    Message: reply,
  })
}

func (s *ServerBidirectional) Chat(stream pb.Chat_ChatServer) error {
  for {
    in, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      return err
    }
    message := in.GetMessage()
    fmt.Println("get:", message)
    if err := request(stream, message); err != nil {
      return err
    }
    time.Sleep(time.Second)
  }
}

func set() error {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    return errors.Wrap(err, "port failed")
  }
  s := grpc.NewServer()
  var server ServerBidirectional
  pb.RegisterChatServer(s, &server)
  if err := s.Serve(lis); err != nil {
    return errors.Wrap(err, "server boot failed")
  }
  return nil 
}
 
func main(){
  fmt.Println("boot")
  if err := set(); err != nil {
    log.Fatalf("%v", err)
  }
}
