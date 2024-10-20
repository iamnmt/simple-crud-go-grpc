package main

import (
	"context"
	"log"

	"ms/server/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:6969"
)

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc::NewClient %v", err)
	}

	defer conn.Close()

	client := types.NewMsgClient(conn)
	ctx := context.Background()

	log.Println("CREATE:")
	// Success CREATE
	{

		r, err := client.Create(ctx, &types.MsgCreate{
			Id:      69,
			Content: "420",
		})
		if err != nil {
			log.Fatalf("client::Create %v", err)
		}
		log.Printf("\tStatus: %t", r.GetStatus())
	}

	log.Println("READ:")

	// Success READ
	{
		r, err := client.Read(ctx, &types.MsgRead{
			Id: 69,
		})
		if err != nil {
			log.Fatalf("client::Read %v", err)
		}
		log.Printf("\tContent: %s", r.GetContent())
	}

	// Fail READ
	{
		_, err := client.Read(ctx, &types.MsgRead{
			Id: 68,
		})
		if err != nil {
			log.Printf("\tclient::Read %v", err)
		}
	}

	log.Println("UPDATE:")
	// Success UPDATE
	{
		r, err := client.Update(ctx, &types.MsgUpdate{
			Id:      69,
			Content: "421",
		})
		if err != nil {
			log.Fatalf("client::Update %v", err)
		}
		log.Printf("\tStatus: %t", r.GetStatus())
	}

	// READ
	{
		r, err := client.Read(ctx, &types.MsgRead{
			Id: 69,
		})
		if err != nil {
			log.Fatalf("client::Read %v", err)
		}
		log.Printf("\tContent: %s", r.GetContent())
	}

	// Fail UPDATE
	{
		_, err := client.Update(ctx, &types.MsgUpdate{
			Id:      68,
			Content: "421",
		})
		if err != nil {
			log.Printf("\tclient::Update %v", err)
		}
	}

	log.Println("DELETE:")
	// Success DELETE
	{
		r, err := client.Delete(ctx, &types.MsgDelete{
			Id: 69,
		})
		if err != nil {
			log.Fatalf("client::Delete %v", err)
		}
		log.Printf("\tStatus: %t", r.GetStatus())
	}

	// Fail READ
	{
		_, err := client.Read(ctx, &types.MsgRead{
			Id: 69,
		})
		if err != nil {
			log.Printf("\tclient::Read %v", err)
		}
	}

	// Fail DELETE
	{
		_, err := client.Delete(ctx, &types.MsgDelete{
			Id: 69,
		})
		if err != nil {
			log.Printf("\tclient::Delete %v", err)
		}
	}
}
