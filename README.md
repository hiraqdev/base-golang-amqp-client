# base-golang-amqp-client
A skeleton to create rabbitmq listener using golang.  This skeleton
used only to setup basic amqp client listener, and also provide a reconnection
process.

Notes:

```
This skeleton should be always simple and the purpose only to setup new project
to listen amqp (rabbitmq), so I will not provide any magic functionalities or will
develop this skeleton to a framework in the future.
```

---

## Dependencies

RabbitMQ Client: [streadway/amqp](github.com/streadway/amqp)

---

## Command line usages

```
go run *.go -amqpUrl="your-rabbitmq-url"
```

By default, `amqpUrl` will be: `amqp://guest:guest@localhost:5672`

---

This skeleton also provide some debugging process :

```
go run *.go -debug 
```

## Reconnection

This client system will listen on `NotifyClose` based on description [here](https://godoc.org/github.com/streadway/amqp#Connection.NotifyClose).  The implementation
is simple, and i just follow these description :

```
Auto reconnect and re-synchronization of client and server topologies.

    Reconnection would require understanding the error paths when the topology cannot be declared on reconnect. This would require a new set of types and code paths that are best suited at the call-site of this package. AMQP has a dynamic topology that needs all peers to agree. If this doesn't happen, the behavior is undefined. Instead of producing a possible interface with undefined behavior, this package is designed to be simple for the caller to implement the necessary connection-time topology declaration so that reconnection is trivial and encapsulated in the caller's application code.
```
[Source](https://github.com/streadway/amqp#non-goals)

Based on these explanation, reconnection process should be happened on caller's application code.