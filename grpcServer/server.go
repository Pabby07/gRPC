package main

import ("log"
		"net"
		// "github.com/grpc-go-course/hello/hellopb"
		"../grpcpb"
		"google.golang.org/grpc"
		"context"
)

type server struct {

}

func (*server) GreetSal(ctx context.Context, request *grpcpb.Request) (*grpcpb.Response, error) {
	name := request.Name
	response := &grpcpb.Response{
		Salutaion: "Hello " + name,
	}
	return response, nil
}

func main() {
	log.Println("started grpc")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err, "error initializing the server")
	}

	grpcser := grpc.NewServer()

	grpcpb.RegisterSatuationGreetServer(grpcser, &server{})
	grpcser.Serve(lis)
}