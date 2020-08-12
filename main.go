package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log"
	"time"
)

func CognitoHandler(ctx context.Context) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)
}

func LongRunningHandler(ctx context.Context) (string, error) {

	deadline, _ := ctx.Deadline()
	deadline = deadline.Add(-100 * time.Millisecond)
	timeoutChannel := time.After(time.Until(deadline))

	for {
		select {
			case <- timeoutChannel:
				return "Finished before timing out.", nil

		default:
			log.Print("hello!")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	lambda.Start(LongRunningHandler)
}
