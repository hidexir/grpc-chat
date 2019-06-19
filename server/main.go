package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/hidexir/grpc-chat/proto"

	"google.golang.org/grpc"
)

type ChatService struct {
}

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	chatService := &ChatService{}
	pb.RegisterChatServer(server, chatService)
	server.Serve(listenPort)
}

func (ChatService) SendMessage(ctx context.Context, in *SendMessageReq) (*pb.SendMessageRes, error) {
	fmt.Println("sucsess!")
	return nil, nil
}
func InRoom(stream pb.Chat_InRoomServer) error {
}
