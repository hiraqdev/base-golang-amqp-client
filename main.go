package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

var (
	amqpURL string
	isDebug bool
	logger  *Logger
)

const (
	amqpURLDefault string = "amqp://guest:guest@localhost:5672"
	amqpURLUsage   string = "Set your amqp url address"
	isDebugUsage   string = "Flag current process for debugging"
	isDebugDefault bool   = false
)

func init() {
	flag.StringVar(&amqpURL, "amqpUrl", amqpURLDefault, amqpURLUsage)
	flag.BoolVar(&isDebug, "debug", isDebugDefault, isDebugUsage)
	flag.Parse()

	logger = logBuilder(isDebug)
}

func main() {
	logger.Info(fmt.Sprintf("RabbitMQURL: %s", amqpURL))
	logger.Info(fmt.Sprintf("Is debug? %t", isDebug))

	// handle connection and start to watch connection
	// if connection broken, then it should try to reconnecting..
	sessConn := make(chan *amqp.Connection)
	handleConnection(amqpURL, sessConn)
	go func() {
		for {
			select {
			case conn := <-sessConn:
				logger.Info("Start to listening queues...")

				logger.Info(conn) // you can remove this log
				// setup your channel / routing / queue here
			}
		}
	}()

	forever := make(chan bool)
	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sigchan {
			logger.Info(fmt.Sprintf("Interrupted process: %s", sig))
			forever <- false
		}
	}()

	<-forever
	logger.Info("Stopping process...")
}
