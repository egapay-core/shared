package broker

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"strings"
)

type brokerConfig struct {
	ClientPros string `opfield:"client_properties"`
}

func readFromOptsString(opts string) kafka.ConfigMap {
	// reads the client configuration from client.properties
	// and returns it as a key-value map
	m := make(map[string]kafka.ConfigValue)
	
	// read the client properties
	scanner := bufio.NewScanner(strings.NewReader(opts))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}
	
	return m
}
