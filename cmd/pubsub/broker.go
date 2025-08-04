package pubsub

type BrokerMessage struct {
	topic   string
	payload string
}

type Broker struct {
	channel chan BrokerMessage
}

var b = Broker{
	channel: make(chan BrokerMessage),
}

func GetMessageBroker() *Broker {
	return &b
}

func (b *Broker) listen() {
	for msg := range b.channel {
		ConsumeMessage(msg)
	}
}

func StartMessageBroker() {
	b.listen()
}
