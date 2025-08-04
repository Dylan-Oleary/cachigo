package tcp

import (
	"fmt"
	"net"

	"github.com/Dylan-Oleary/cachigo/cmd/pubsub"
	"github.com/Dylan-Oleary/cachigo/store"
)

func HandleRequest(req *Request, res *Response, conn net.Conn) *Response {
	c := store.GetCache()

	switch req.Data.Command {
	case "del":
		c.Remove(req.Data.Key, true)
		res.Success = true
		res.Message = "Success"
	case "get":
		v, err := c.Get(req.Data.Key)

		if err != nil {
			return handleRequestError(res, err)
		}

		res.Success = true
		res.Message = v
	case "keys":
		for _, k := range c.ListKeys() {
			res.Message += k + "\n"
		}

		res.Success = true
	case "set":
		v, err := c.Set(req.Data.Key, req.Data.Value, true)

		if err != nil {
			return handleRequestError(res, err)
		}

		res.Success = v
		res.Message = "Value Set"
	case "pub":
		p := pubsub.Publisher{Topic: req.Data.Key}
		p.PublishMessage(req.Data.Value)

		res.Success = true
		res.Message = "Message Sent to broker"
	case "sub":
		pubsub.SubscribeToTopic(req.Data.Key, conn)

		res.Success = true
		res.Message = "Subscribed to topic"
	default:
		res.Message = "Default Command"
		res.Success = true
	}

	return res
}

func handleRequestError(res *Response, err error) *Response {
	fmt.Println("Error:", err)

	res.Success = false
	res.Message = err.Error()

	return res
}
