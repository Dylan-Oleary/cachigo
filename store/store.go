package store

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

const PersistencePer = 0600

// TODO:
// Abstract Commands For Reusability

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

func (c *cache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]

	if !ok {
		return "No value found", nil
	}

	return v, nil
}

func (c *cache) ListKeys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.data))

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

func (c *cache) Remove(key string, logOp bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
	if logOp {
		logOperation(logOpts{key: key, cmd: "del"})
	}
}

func (c *cache) Set(key string, value string, logOp bool) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	if logOp {
		logOperation(logOpts{key: key, value: value, cmd: "set"})
	}

	return true, nil
}

func InitPersistence() error {
	filePath := "./cmd_logs.txt"

	_, err := os.Lstat(filePath)
	if err != nil {
		fmt.Println("Failed to find existing log file. Creating new log file")

		err := os.WriteFile(filePath, []byte(""), PersistencePer)
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
			c.Remove(key, false)
			continue
		case "set":
			v := args[2]
			c.Set(key, v, false)
			continue
		default:
			fmt.Println("Invalid command found in log file: ", cmd, " Skipping operation")
		}
	}

	fmt.Println("Persistence layer initialized successfully")

	return nil
}

type logOpts struct {
	cmd   string
	key   string
	value string
}

func logOperation(o logOpts) {
	filePath := "./cmd_logs.txt"

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, PersistencePer)
	if err != nil {
		fmt.Println("Failed to open log file: ", o.cmd, o.key, "Error:", err)
		return
	}

	var log []string

	log = append(log, o.cmd)
	log = append(log, o.key)

	if o.value != "" {
		log = append(log, o.value)
	}

	_, err = f.Write([]byte(strings.Join(log, " ") + "\n"))
	if err != nil {
		fmt.Println("Failed to write to log file: ", o.cmd, o.key, "Error:", err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println("Failed to close log file: ", o.cmd, o.key, "Error:", err)
		return
	}
}
