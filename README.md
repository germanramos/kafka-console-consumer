# kafka-console-consumer

Golang kafka-consumer using sarama.

It prints consumed messages in stdout

## Usage

```
./service
```

## Configuration

Use environment variables

- KAFKA_SERVICE, "kafka", The DNS name for input Kafka broker service
- KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "create-agreement", The topic to consume
- KAFKA_GROUP, "kafka-console-consumer", The kafka consumer group ID
- OFFSET, "newest", The offset to start with in new topic. Can be `oldest`, `newest`
- VERBOSE, "false, Set to `true` if you want verbose output
