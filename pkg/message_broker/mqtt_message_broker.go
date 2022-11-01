package message_broker

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const AT_MOST_ONCE byte = 0
const AT_LEAST_ONCE byte = 1
const ONLY_ONCE byte = 2

type MQTTMessageBroker struct {
	client  mqtt.Client
	quitter chan struct{}
}

func New(broker string, port int, clientId string, username string, password string, waitTimeout time.Duration) *MQTTMessageBroker {
	mb := MQTTMessageBroker{}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)

	mb.client = mqtt.NewClient(opts)

	token := mb.client.Connect()
	token.WaitTimeout(waitTimeout)

	return &mb
}

func (mb *MQTTMessageBroker) Publish(topic string, recFq byte, retain bool, data string, waitTimeout time.Duration) {
	token := mb.client.Publish(topic, recFq, retain, data)
	token.WaitTimeout(waitTimeout)

}

func (mb *MQTTMessageBroker) Subscribe(topic string, recFq byte, waitTimeout time.Duration, do func(msg []byte)) {
	token := mb.client.Subscribe(topic, recFq, func(c mqtt.Client, m mqtt.Message) {
		do(m.Payload())
	})
	token.WaitTimeout(waitTimeout)
}

func (mb *MQTTMessageBroker) SyncSubscribe(topic string, recFq byte, do func(msg []byte)) {
	for {
		token := mb.client.Subscribe(topic, recFq, func(c mqtt.Client, m mqtt.Message) {
			do(m.Payload())
		})
		token.Wait()
	}
}

func (mb *MQTTMessageBroker) Close() {
	mb.client.Disconnect(1000)
	mb.quitter <- struct{}{}
}
