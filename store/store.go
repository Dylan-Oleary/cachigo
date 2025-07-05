package store

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Store interface {
	get(key string) (string, error)
	listKeys() ([]string, error)
	set(key string, value string) (bool, error)
	remove(key string)
}

type cache struct {
	data map[string]string
	mu   sync.Mutex
}

var c = cache{
	data: make(map[string]string),
}

func GetCache() *cache {
	return &c
}

// TODO
// Write to cmd log file
func (c *cache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]

	if !ok {
		return "No value found", nil
	}

	return v, nil
}

// TODO
// Write to cmd log file
func (c *cache) ListKeys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.data))

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

// TODO
// Write to cmd log file
func (c *cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

// TODO
// Write to cmd log file
func (c *cache) Set(key string, value string) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	return true, nil
}

func InitPersistence() error {
	filePath := "./cmd_logs.txt"

	_, err := os.Lstat(filePath)
	if err != nil {
		fmt.Println("Failed to find existing log file. Creating new log file")

		err := os.WriteFile(filePath, []byte(""), 0600)
		if err != nil {
			fmt.Println("Failed to create log file: ", err)
			return err
		}

		fmt.Println("Log file created successfully")
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to read contents of log file: ", err)
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ln := scanner.Text()
		args := strings.Split(ln, " ")

		cmd := args[0]
		key := args[1]

		// TODO: Abstract commands so they can be reused + validation
		switch cmd {
		case "get":
			continue
		case "del":
			c.Remove(key)
			continue
		case "set":
			v := args[2]
			c.Set(key, v)
			continue
		default:
			fmt.Println("Invalid command found in log file: ", cmd, " Skipping operation")
		}
	}

	fmt.Println("Persistence layer initialized successfully")

	return nil
}
