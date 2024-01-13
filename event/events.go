package event

import (
	"encoding/json"

	"github.com/DecodeWorms/messaging-protocol/pulse"
	"sv.notify/email"
	"sv.notify/model"
)

type Events struct {
	pulsar pulse.Message
}

func NewEvents(p pulse.Message) *Events {
	return &Events{
		pulsar: p,
	}
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
		return err
	}
	return nil
}
