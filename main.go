package main

import "main/mosaic_server"

func main() {
	server := mosaic_server.Server{}
	server.Load()
}

/*
func client() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := server.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &server.Message{Text: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Text)
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := server.Server{}
	grpcServ := grpc.NewServer()
	server.RegisterChatServiceServer(grpcServ, &serv)
	ch := make(chan bool)
	go func() {
		if err := grpcServ.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
		ch <- true
	}()
	go client()
	<-ch
}
*/
