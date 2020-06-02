package Worker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"video-analysis-worker/Http"
	"video-analysis-worker/Models"
	"video-analysis-worker/VideoAI"
	"video-analysis-worker/grpc_client"
)

type Worker struct {
	RabbitMQ *RabbitMqConnection
	env *Models.Env
	mediaMetadataGrpcClient *grpc_client.MediaMetadataClient
	mediaDowLoader *Http.MediaDownloader
	videoAI *VideoAI.VideoAI
}

func (worker *Worker) Work()  {
	forever := make(chan bool)
	go func() {
		for d := range worker.RabbitMQ.msgs {
			log.Printf("Received a message: %s", d.Body)

			rabbiMQMessageAnalysis := &Models.RabbitMQMessageAnalysis{}
			err := json.Unmarshal(d.Body, rabbiMQMessageAnalysis)
			if err != nil{
				log.Println(err)
			}

			// fmt.Println(rabbiMQMessageAnalysis)
			mediaMetadata, err := worker.mediaMetadataGrpcClient.GetMediaMetadata(rabbiMQMessageAnalysis.MediaId)
			if err != nil{
				log.Println(err)
			}
			// fmt.Println(mediaMetadata)

			fileUrl := worker.env.AwsStorageUrl + "v1/awsStorage/media/" + mediaMetadata.AwsBucketWholeMedia + "/" + mediaMetadata.AwsStorageNameWholeMedia
			err = worker.mediaDowLoader.DownloadFile("./assets/" + mediaMetadata.AwsStorageNameWholeMedia, fileUrl)
			if err != nil {
				log.Println(err)
			}
			labelsStringArray, err := worker.videoAI.Label("./assets/" + mediaMetadata.AwsStorageNameWholeMedia)
			if err != nil {
				log.Println(err)
			}
			// fmt.Println()
			// fmt.Println(labelsStringArray)

			_, err = worker.mediaMetadataGrpcClient.UpdateMediaKeywords(mediaMetadata.MediaId, labelsStringArray)

			if err != nil{
				log.Println(err)
			}
			// fmt.Println(updatedMetadata)

			worker.removeFile("./assets/" + mediaMetadata.AwsStorageNameWholeMedia)

			log.Printf("Done")
			_ = d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}



func (worker *Worker) removeFile(path string)  {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}

func InitWorker() *Worker  {
	return &Worker{
		RabbitMQ: 					initRabbitMqConnection(Models.GetEnvStruct()),
		env:      					Models.GetEnvStruct(),
		mediaMetadataGrpcClient: 	grpc_client.InitMediaMetadataGrpcClient(),
		mediaDowLoader:				&Http.MediaDownloader{},
		videoAI: 					&VideoAI.VideoAI{},
	}
}

