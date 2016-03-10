# convolutional-neural-network

**TODO**

- Reads events from Kafka

## Usage

### Input kafka configuration

- INPUT_KAFKA, "localhost", The DNS name for input Kafka broker service
- INPUT_KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "etl-create-agreement", The topic to consume
- PARTITION, "all", The partitions to consume, can be 'all' or comma-separated numbers
- OFFSET, "newest", The offset to start with. Can be 'oldest' or 'newest'
- BUFFER_SIZE, "256", The buffer size of the input message channel.

### Common configuration

- VERBOSE, "false", 'true' if you want to turn on sarama logging

## Run Example:

INPUT_KAFKA=broker.kafka.skydns.local ./main
