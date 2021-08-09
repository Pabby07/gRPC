package main

import ("log"
		// "net"
		// "github.com/grpc-go-course/hello/hellopb"
		"../grpcpb"
		"google.golang.org/grpc"
		"context"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	grpcClient := grpcpb.NewSatuationGreetClient(cc)

	req := &grpcpb.Request{Name: "pabby ji"}
	resp, err := grpcClient.GreetSal(context.Background(), req)
	if err != nil {
		log.Fatal("Error recieving the response")
	}
	log.Println("ddrrreeecccttt from grpc response",resp.Salutaion)
}