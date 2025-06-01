package tcp

import (
	"encoding/json"
	"fmt"
	"net"
)

type GetRequest struct {
	Command string `json:"command"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func GetClient(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("Failed to connect to TCP server at %s\n", addr)
		return nil, err
	}

	return conn, nil
}

func SendRequest(conn net.Conn, req *GetRequest) (*Response, error) {
	payload, err := json.Marshal(req)

	if err != nil {
		fmt.Println("Failed to process request payload")
		return nil, err
	}

	_, err = conn.Write(payload)

	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 1024)
	size, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Could not decode response", err)
		return nil, err
	}

	data := Response{}
	err = json.Unmarshal(buffer[:size], &data)

	if err != nil {
		fmt.Println("Could not decode response", err)
		return nil, err
	}

	return &data, nil
}
