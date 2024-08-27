package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"sync"
	"tracking_system/internal/entities"
	"tracking_system/internal/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	log := logger.GetLogger()
	// Parse accountID for filtering events
	accountIDArg := flag.String("accountID", "", "The initial account ID to filter messages by")
	flag.Parse()

	// Set up channels for accountID argument updates and shutdown
	accountIDChan := make(chan *int64)

	// Start a goroutine to read accountID argument updates from standard input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			newAccountIDArg := scanner.Text()
			if newAccountIDArg != "" {
				newAccountID, _ := strconv.ParseInt(newAccountIDArg, 10, 64)
				if newAccountID == -1 {
					accountIDChan <- nil
				} else {
					accountIDChan <- &newAccountID
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.WithError(err).Error("Error reading input")
		}
	}()

	var currentAccountID *int64
	if accountIDArg != nil && len(*accountIDArg) > 0 {
		accountID, _ := strconv.ParseInt(*accountIDArg, 10, 64)
		currentAccountID = &accountID
	}

	config := kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "messages",
		"auto.offset.reset": "smallest",
	}

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		log.WithError(err).Error("Error creating new consumer")
	}

	var mu sync.Mutex
	consumer.Subscribe("account-message-events", nil)
	for {
		select {
		case newAccountID := <-accountIDChan:
			mu.Lock()
			currentAccountID = newAccountID
			if currentAccountID != nil {
				log.Infof("Account ID updated to: %d", *currentAccountID)
			} else {
				log.Info("Account ID filter removed")
			}
			mu.Unlock()

		default:
			event := consumer.Poll(100)
			switch e := event.(type) {
			case *kafka.Message:
				event := &entities.Event{}
				err := json.Unmarshal(e.Value, event)
				if err != nil {
					log.WithError(err).Error("Error occured while unmarshaling event data")
				}

				// Display only events if there is no accountID filter set or event's accountID matched accountID filter
				if currentAccountID == nil || event.AccountID == *currentAccountID {
					log.Infof("AccountID: %d | Message: %s | Timestamp: %v", event.AccountID, event.Data, event.Timestamp)
				}

			case kafka.Error:
				log.WithError(err).Error("Kafka error")
			}
		}
	}
}
