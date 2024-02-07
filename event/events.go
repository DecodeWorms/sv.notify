package event

import (
	"encoding/json"
	"log"

	mesmessagingprotocol "github.com/DecodeWorms/messaging-protocol"
	"github.com/DecodeWorms/messaging-protocol/pulse"
	"sv.notify/email"
	"sv.notify/model"
)

type Events struct {
	pulsar *pulse.Message
}

func NewEvents(pulsarUrl, serviceName string) (*Events, error) {
	opt := mesmessagingprotocol.Options{
		Address:     pulsarUrl,
		ServiceName: serviceName,
	}
	pMessage, err := pulse.NewMessage(opt)
	if err != nil {
		return nil, err
	}
	return &Events{
		pulsar: pMessage,
	}, nil
}

func (e *Events) SendWelcomeMessage(topic string) error {
	msg, err := e.pulsar.Subscriber(topic)
	if err != nil {
		return err
	}
	var welcome model.Welcome
	if err = json.Unmarshal(msg, &welcome); err != nil {
		return err
	}

	// Send welcome message to the player
	if err = email.SendWelcomeMessage(welcome); err != nil {
		log.Printf("The error is %v", err)
		return err
	}
	return nil
}
