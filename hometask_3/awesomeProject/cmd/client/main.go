package main 

import (
	"awesomeProject/proto"

	"context"
	"flag"
	"fmt"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	command := flag.String("cmd", "", "type of command (get, create, delete, change, patch)")
	name := flag.String("name", "", "name of the account")
	amount := flag.Int("amount", 0, "amount of the account")
	newName := flag.String("new_name", "", "new name of the account")
	newAmount := flag.Int("new_amount", 0, "new amount for the account")

	flag.Parse()

	conn, err := grpc.NewClient("localhost:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	client := proto.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
		
	switch *command {
	case "get":
		res, err := client.GetAccount(ctx, &proto.GetAccountRequest{
			Name:   *name,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("name: %s, amount: %d\n", res.Name, res.Amount)
	case "create":
		res, err := client.CreateAccount(ctx, &proto.CreateAccountRequest{
			Name:   *name,
			Amount: int32(*amount),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Account %s has been created successfully\n", res.Name)
	case "delete":
		res, err := client.DeleteAccount(ctx, &proto.DeleteAccountRequest{
			Name:   *name,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Account %s has been deleted successfully\n", res.Name)
	case "change":
		res, err := client.ChangeAccount(ctx, &proto.ChangeAccountRequest{
			Name:   *name,
			NewName: *newName,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Account name has been changed to %s successfully\n", res.Name)
	case "patch":
		res, err := client.PatchAccount(ctx, &proto.PatchAccountRequest{
			Name:   *name,
			NewAmount: int32(*newAmount),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Account %s has patched successfully\n", res.Name)
	}
}