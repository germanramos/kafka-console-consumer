# kafka-console-consumer

kafka-console-consumer implemented in golang and using [sarama](https://github.com/Shopify/sarama) driver.

## Features

- Works with Apache Kafka 0.8, 0.9 and 0.10
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

## Set your GOPATH environment
For example:
```
export GOPATH=`pwd`/gopath
```

## Update dependencies (you probably do not need this)
```
rm -rf gopath
go get -v ./...
cd gopath
find . -type d -name ".git" | xargs rm -rf
```

## Local Build
```
# This will generate kafka-console-consumer binary
go build
```
or
```
# This will generate "service" binary
./build.sh
```

## Build in Docker (outputs Linux binany)
```
# This will generate "service" *Linux* binary
./build-in-docker.sh
```

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

## Download

https://github.com/germanramos/kafka-console-consumer/releases/download/v0.3.0/kafka-console-consumer

## Run Example

```
KAFKA_SERVICE=192.168.1.45 TOPIC=foo ./kafka-console-consumer
```

## Related work

https://github.com/germanramos/kafka-console-producer

## License

MIT - German Ramos
