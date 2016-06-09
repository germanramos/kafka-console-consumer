# kafka-console-consumer

Golang kafka-console-consumer using sarama driver.

## Features

- Statically compiled. No dependencies. It run on every linux distribution.
- Just one binary file of ~ 6 Mb
- It prints consumed messages in stdout
- It prints state messages in stderr
- Very easy to configure trough environment variables
- Auto discover kafka peers from DNS name
- Consumer group support
- Customizable initial offset for topic consuming
- Waits for kafka to be ready
- Waits for topic to be created
- Auto reconnect

## Usage

```
./kafka-console-consumer
```

## Configuration

environment variable, default value, description

- KAFKA_SERVICE, "kafka", The DNS name or IP for input Kafka broker service
- KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "mytopic", The topic to consume
- KAFKA_GROUP, "kafka-console-consumer", The kafka consumer group ID
- OFFSET, "newest", The offset to start with in new topic. Can be `oldest`, `newest`
- BUFFER_SIZE, "256", The buffer size of stdout message channel
- VERBOSE, "false, Set to `true` if you want verbose output

## Run Example

```
KAFKA_SERVICE=192.168.1.45 TOPIC=foo ./kafka-console-consumer
```

## License

MIT - German Ramos
