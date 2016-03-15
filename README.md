# kafka-datasource

Golang kafka-consumer using sarama.

It prints consumed messages in stdout

## Usage

```
./main
```

## Configuration

Use environment variables

- INPUT_KAFKA, "localhost", The DNS name for input Kafka broker service
- INPUT_KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "create-agreement", The topic to consume
- PARTITION, "all", The partitions to consume, can be 'all' or comma-separated numbers
- OFFSET, "newest", The offset to start with. Can be `oldest`, `newest`
- BUFFER_SIZE, "256", The buffer size of the message channel. Recommended `256`
