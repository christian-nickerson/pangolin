package embeddings

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/christian-nickerson/pangolin/pangolin/internal/proto"
)

var Client proto.EmbeddingsClient
var Conn *grpc.ClientConn

func Connect(address string) {
	var err error

	Conn, err = grpc.NewClient(address)
	if err != nil {
		log.Fatal("Failed to connect to client:", err)
	}

	Client = proto.NewEmbeddingsClient(Conn)
}

func Inference(text *[]string, modelName string) []*proto.Vector {
	// timeout after 5 mins
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// call model
	response, err := Client.Inference(ctx, &proto.InferenceRequest{Text: *text, ModelName: modelName})
	if err != nil {
		log.Fatalf("Client.Inference failed on model %v: %v", modelName, err)
	}

	return response.Embeddings
}

func ModelList() []string {
	// timeout after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// call model
	response, err := Client.ModelList(ctx, &proto.ModelListRequest{})
	if err != nil {
		log.Fatalf("Client.ModelList failed: %v", err)
	}

	return response.ModelNames
}
