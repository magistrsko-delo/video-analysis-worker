package VideoAI

import (
	video "cloud.google.com/go/videointelligence/apiv1"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
	"io/ioutil"
)

type VideoAI struct {

}


func (videoAI *VideoAI) Label(file string) ([]string, error) {
	labelsStringArray := []string{}
	ctx := context.Background()
	client, err := video.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("video.NewClient: %v", err)
	}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
		InputContent: fileBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("AnnotateVideo: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("Wait: %v", err)
	}

	printLabels := func(labels []*videopb.LabelAnnotation) {
		for _, label := range labels {
			fmt.Printf( "\tDescription: %s\n", label.Entity.Description)
			labelsStringArray = append(labelsStringArray, label.Entity.Description)
			for _, category := range label.CategoryEntities {
				fmt.Printf("\t\tCategory: %s\n", category.Description)
				labelsStringArray = append(labelsStringArray, category.Description)
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

	return labelsStringArray, nil
}