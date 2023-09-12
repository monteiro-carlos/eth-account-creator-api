package queue

import (
	"encoding/json"
	"eth-account-creator-api/internal/log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Client struct {
	log *log.Logger
}

func NewClient(
	logger *log.Logger,
) (*Client, error) {
	return &Client{
		log: logger,
	}, nil
}

func (c *Client) SendMessage(myMessage map[string]interface{}) error {
	qURL := os.Getenv("QUEUE_URL")

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})

	svc := sqs.New(sess)
	body, _ := json.Marshal(myMessage)
	messageGroupId := "mymessagegroupid"
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:    aws.String(string(body)),
		MessageGroupId: &messageGroupId,
		QueueUrl:       &qURL,
	})

	return err
}
