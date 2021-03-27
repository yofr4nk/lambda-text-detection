package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/yofr4nk/lambda-text-detection/sessions"
)

func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {

	fmt.Printf("This is your event: " + event.Records[0].EventName + "\n")

	file := event.Records[0].S3.Object.Key
	bucket := event.Records[0].S3.Bucket.Name

	awsSession := sessions.CreateAwsSession()

	// Create a Rekognition client
	svc := rekognition.New(awsSession)

	output, err := svc.DetectText(&rekognition.DetectTextInput{
		Filters: nil,
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: &bucket,
				Name:   &file,
			},
		},
	})

	if err != nil {
		fmt.Print(err)

		return "", err
	}

	for _, text := range output.TextDetections {
		fmt.Printf("This is your text detected in your image: %s!", *text.DetectedText+"\n")
	}

	return fmt.Sprintf("This is your detection %s!", output.GoString()), nil
}

func main() {
	lambda.Start(HandleRequest)
}
