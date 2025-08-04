package pubsub

type Publisher struct {
	Topic string
}

func (t *Publisher) PublishMessage(message string) {
	b := GetMessageBroker()
	b.channel <- BrokerMessage{topic: t.Topic, payload: message}
}
