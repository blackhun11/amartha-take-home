package pubsub

import (
	"context"
	"time"

	pubsubPkg "cloud.google.com/go/pubsub"
)

type pubsub struct {
	*pubsubPkg.Client
}

type Pubsub interface {
	Publish(ctx context.Context, data []byte, topicID string) (string, error)
}

func NewPubsubClient(client *pubsubPkg.Client) Pubsub {

	return &pubsub{
		Client: client,
	}
}

func (p pubsub) Publish(ctx context.Context, data []byte, topicID string) (string, error) {
	return p.Client.Topic(topicID).Publish(ctx, &pubsubPkg.Message{
		Data:        data,
		PublishTime: time.Now(),
	}).Get(ctx)
}
