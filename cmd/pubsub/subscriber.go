package pubsub

import (
	"fmt"
	"net"
)

type Subscribers struct {
	connections map[string][]net.Conn
}

var subscribers = Subscribers{connections: make(map[string][]net.Conn)}

func ConsumeMessage(msg BrokerMessage) error {
	conns := subscribers.connections[msg.topic]

	if len(conns) == 0 {
		fmt.Println("No connections found for topic", msg.topic)
		return nil
	}

	for _, c := range conns {
		c.Write([]byte(msg.payload))
	}

	return nil
}

func SubscribeToTopic(topic string, conn net.Conn) {
	subscribers.connections[topic] = append(subscribers.connections[topic], conn)
}
