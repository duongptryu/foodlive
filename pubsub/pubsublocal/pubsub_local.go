package pubsublocal

import (
	"context"
	"fmt"
	"fooddelivery/common"
	"fooddelivery/pubsub"
	"sync"
)

type localPubSub struct {
	locker            *sync.RWMutex
	messageQueue      chan *pubsub.Message
	mapTopicSubscribe map[pubsub.Topic][]chan *pubsub.Message
}

func NewPubSubLocal() *localPubSub {
	pb := &localPubSub{
		locker:            new(sync.RWMutex),
		messageQueue:      make(chan *pubsub.Message, 10000),
		mapTopicSubscribe: make(map[pubsub.Topic][]chan *pubsub.Message),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetTopic(topic)

	go func() {
		defer common.AppRecovery()
		ps.messageQueue <- data
		fmt.Println("New event published:", data.String(), "With data", data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapTopicSubscribe[topic]; ok {
		val = append(ps.mapTopicSubscribe[topic], c)
		ps.mapTopicSubscribe[topic] = val
	} else {
		ps.mapTopicSubscribe[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()
	fmt.Println("Added subscriber - ", topic)

	return c, func() {
		fmt.Println("Unsubscribe")

		if chans, ok := ps.mapTopicSubscribe[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					//remove element at index
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapTopicSubscribe[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	fmt.Println("PubSub started")

	go func() {
		for {
			mess := <-ps.messageQueue
			fmt.Println("Message dequeue:", mess)

			if subs, ok := ps.mapTopicSubscribe[mess.Topic()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
