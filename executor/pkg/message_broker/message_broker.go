package message_broker

type IMessageBroker interface {
	Subscribe(topic string)
}
