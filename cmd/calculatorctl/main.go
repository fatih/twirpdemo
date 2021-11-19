package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fatih/twirpdemo/internal/rpc"
)

func main() {
	// Always use Protobuf Client. It's compact and faster to serialize.
	client := rpc.NewCalculatorProtobufClient("http://localhost:8080", http.DefaultClient)

	resp, err := client.Add(context.Background(), &rpc.AddReq{
		A: 10,
		B: 5,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result = %+v\n", resp.GetResult())
}
