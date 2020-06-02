package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

type Env struct {
	RabbitUser string
	RabbitPassword string
	RabbitQueue string
	RabbitHost string
	RabbitPort string
	AwsStorageUrl string
	Env string
	MediaMetadataGrpcServer string
	MediaMetadataGrpcPort string
}

func InitEnv()  {
	envStruct = &Env{
		RabbitUser:       			os.Getenv("RABBIT_USER"),
		RabbitPassword:   			os.Getenv("RABBIT_PASSWORD"),
		RabbitQueue:      			os.Getenv("RABBIT_QUEUE"),
		RabbitHost:       			os.Getenv("RABBIT_HOST"),
		AwsStorageUrl:   			os.Getenv("AWS_STORAGE_URL"),
		RabbitPort: 				os.Getenv("RABBIT_PORT"),
		Env: 			  			os.Getenv("ENV"),
		MediaMetadataGrpcServer: 	os.Getenv("MEDIA_METADATA_GRPC_SERVER"),
		MediaMetadataGrpcPort:   	os.Getenv("MEDIA_METADATA_GRPC_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}