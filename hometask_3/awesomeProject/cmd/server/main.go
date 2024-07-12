package main

import (
	"awesomeProject/proto"
	"awesomeProject/accounts"

	"fmt"
	"net"
	"context"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAccountServiceServer
}

var accountHandler = accounts.New()

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	if account, err := accountHandler.GetAccount(req); err != nil {
		return &proto.GetAccountResponse{Name: req.Name, Amount: 0, Success: false}, err
	} else {
		return &proto.GetAccountResponse{Name: account.Name, Amount: account.Amount, Success: true}, nil
	}
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CommonAccountResponse, error) {
	if err := accountHandler.CreateAccount(req); err != nil {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "create", Success: false}, err
	} else {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "create", Success: true}, nil
	}
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.CommonAccountResponse, error) {
	if err := accountHandler.DeleteAccount(req); err != nil {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "Delete", Success: false}, err
	} else {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "Delete", Success: true}, nil
	}
}

func (s *server) ChangeAccount(ctx context.Context, req *proto.ChangeAccountRequest) (*proto.CommonAccountResponse, error) {
	if err := accountHandler.ChangeAccount(req); err != nil {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "Change", Success: false}, err
	} else {
		return &proto.CommonAccountResponse{Name: req.NewName, Operation: "Change", Success: true}, nil
	}
}

func (s *server) PatchAccount(ctx context.Context, req *proto.PatchAccountRequest) (*proto.CommonAccountResponse, error) {
	if err := accountHandler.PatchAccount(req); err != nil {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "Patch", Success: false}, err
	} else {
		return &proto.CommonAccountResponse{Name: req.Name, Operation: "Patch", Success: true}, nil
	}
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    proto.RegisterAccountServiceServer(s, &server{}) 
    if err := s.Serve(lis); err != nil {
        panic(err)
    }
}