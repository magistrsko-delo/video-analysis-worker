package grpc_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"video-analysis-worker/Models"
	pbMediaMetadata "video-analysis-worker/proto/media_metadata"
)


type MediaMetadataClient struct {
	Conn *grpc.ClientConn
	client pbMediaMetadata.MediaMetadataClient
}

func (mediaMetadataClient *MediaMetadataClient) GetMediaMetadata(mediaId int) (*pbMediaMetadata.MediaMetadataResponse, error)  {

	response, err := mediaMetadataClient.client.GetMediaMetadata(context.Background(), &pbMediaMetadata.GetMediaMetadataRequest{
		MediaId:              int32(mediaId),
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mediaMetadataClient *MediaMetadataClient) UpdateMediaKeywords(mediaId int32, keywords []string) (*pbMediaMetadata.MediaMetadataResponse, error)  {

	response, err := mediaMetadataClient.client.UpdateMediaKeywords(context.Background(), &pbMediaMetadata.UpdateMediaKeywords{
		MediaId:              mediaId,
		Keywords:             keywords,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func InitMediaMetadataGrpcClient() *MediaMetadataClient  {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING")
	conn, err := grpc.Dial(env.MediaMetadataGrpcServer + ":" + env.MediaMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION")

	client := pbMediaMetadata.NewMediaMetadataClient(conn)
	return &MediaMetadataClient{
		Conn:    conn,
		client:  client,
	}

}
