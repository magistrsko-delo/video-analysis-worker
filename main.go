
// Sample video_quickstart uses the Google Cloud Video Intelligence API to label a video.
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/ptypes"

	video "cloud.google.com/go/videointelligence/apiv1"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
)

func main() {
	/*ctx := context.Background()

	// Creates a client.
	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: "gs://cloud-samples-data/video/cat.mp4",
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start annotation job: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to annotate: %v", err)
	}

	// Only one video was processed, so get the first result.
	// fmt.Println(resp.GetAnnotationResults())
	result := resp.GetAnnotationResults()[0]
	for _, annotation := range result.SegmentLabelAnnotations {
		fmt.Printf("Description: %s\n", annotation.Entity.Description)

		for _, category := range annotation.CategoryEntities {
			fmt.Printf("\tCategory: %s\n", category.Description)
		}

		for _, segment := range annotation.Segments {
			start, _ := ptypes.Duration(segment.Segment.StartTimeOffset)
			end, _ := ptypes.Duration(segment.Segment.EndTimeOffset)
			fmt.Printf("\tSegment: %s to %s\n", start, end)
			fmt.Printf("\tConfidence: %v\n", segment.Confidence)
		}
	}*/

	_ = label("sample.mp4")
}

func label(file string) error {
	ctx := context.Background()
	client, err := video.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("video.NewClient: %v", err)
	}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
		InputContent: fileBytes,
	})
	if err != nil {
		return fmt.Errorf("AnnotateVideo: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Wait: %v", err)
	}

	printLabels := func(labels []*videopb.LabelAnnotation) {
		for _, label := range labels {
			fmt.Printf( "\tDescription: %s\n", label.Entity.Description)
			for _, category := range label.CategoryEntities {
				fmt.Printf("\t\tCategory: %s\n", category.Description)
			}
			for _, segment := range label.Segments {
				start, _ := ptypes.Duration(segment.Segment.StartTimeOffset)
				end, _ := ptypes.Duration(segment.Segment.EndTimeOffset)
				fmt.Printf("\t\tSegment: %s to %s\n", start, end)
				fmt.Printf("\tConfidence: %v\n", segment.Confidence)
			}
		}
	}

	// A single video was processed. Get the first result.
	result := resp.AnnotationResults[0]

	fmt.Printf("SegmentLabelAnnotations:")
	printLabels(result.SegmentLabelAnnotations)
	fmt.Printf( "ShotLabelAnnotations:")
	printLabels(result.ShotLabelAnnotations)
	fmt.Printf( "FrameLabelAnnotations:")
	printLabels(result.FrameLabelAnnotations)

	return nil
}