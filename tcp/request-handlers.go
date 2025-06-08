package tcp

import (
	"fmt"

	"github.com/Dylan-Oleary/cachigo/store"
)

func HandleRequest(req *Request, res *Response) *Response {
	c := store.GetCache()

	switch req.Data.Command {
	case "get":
		v, err := c.Get(req.Data.Key)

		if err != nil {
			return handleRequestError(res, err)
		}

		res.Success = true
		res.Message = v
	case "set":
		v, err := c.Set(req.Data.Key, req.Data.Value)

		if err != nil {
			return handleRequestError(res, err)
		}

		res.Success = v
		res.Message = "Value Set"
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
