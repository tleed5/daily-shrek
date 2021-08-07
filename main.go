package main

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"math/rand"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ReadShrekLine() string {
    content, err := ioutil.ReadFile("include/shrek.txt")

	if err != nil {
		log.Fatal(err)
	}

	sentences := strings.Split(string(content), ".")
	randomIndex := rand.Intn(len(sentences))
	pick := sentences[randomIndex]

	log.Println(pick)
	return pick;
}

func Handler(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
		return err
	}

	client := sns.NewFromConfig(cfg)

	message := ReadShrekLine()
	phoneNumber := string("todo read from config")
	output, err := client.Publish(ctx, &sns.PublishInput{
		Message: &message,
		PhoneNumber: &phoneNumber,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Print(output)
	return nil
}

func main() {
	lambda.Start(Handler)
}