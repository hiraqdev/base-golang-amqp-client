package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func handleConnection(url string, sessConn chan *amqp.Connection) {
	logger.Info(fmt.Sprintf("Try to connect to: %s", url))
	conn, err := amqp.Dial(url)

	if err != nil {
		logger.Error(err)
		time.Sleep(5 * time.Second)
		handleConnection(url, sessConn)
	} else {
		go func() {
			logger.Info(fmt.Sprintf("Connected to: %s", url))
			sessConn <- conn
		}()

		sessErrConn := make(chan *amqp.Error)
		notifyOnClose := conn.NotifyClose(sessErrConn)
		go func() {
			for errClose := range notifyOnClose {
				logger.Error(fmt.Sprintf("Connection error: %s", errClose))
				ticker := time.NewTicker(5 * time.Second)

				select {
				case <-ticker.C:
					logger.Info("Reconnecting...")
					handleConnection(url, sessConn)
				}
			}
		}()
	}

}
