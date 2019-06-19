package main

import (
	"context"
	"fmt"
	"io"
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

func (ChatService) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageRes, error) {
	fmt.Println("sucsess!")
	return nil, nil
}

func (ChatService) InRoom(stream pb.Chat_InRoomServer) error {
	// 無限ループ
	hoge := &pb.InRoomRes{
		RoomID: "",
	}
	for {
		req, err := stream.Recv()
		fmt.Println(req.RoomID)
		err = stream.Send(hoge)
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
	}
}
