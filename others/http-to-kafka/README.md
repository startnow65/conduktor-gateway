## Http-to-kafka-publisher


## Prerequisites to run the project

1. Run the docker compose file that is present in the root of the project
```bash
docker-compose up -d --force-recreate
```
**Note**: When you spin down the docker-compose setup, the volumes created by schema-registry, broker and zookeeper are not removed by default. Hence in order to bring down the setup please run the following command
```bash
docker-compose down --remove-orphans --volumes
```

2. Set the following environment variables to run the webserver
```bash
export KAFKA_SERVER=localhost
export KAFKA_PORT=9092
export KAFKA_TOPIC=cloudcomms
export HTTP_PORT=7070
```

3. Make sure the docker containers are up and running. Please open the Kafka control center in the browser using **http://localhost:9021**
![control-center](https://github.com/kanapuli/github.com/kanapuli/http-to-kafka/assets/17735366/f71bbbfa-4ec9-4079-a9be-b61739a0f812)


5. Create a topic. For testing the topic name is **cloudcomms**

![topic-creation](https://github.com/kanapuli/github.com/kanapuli/http-to-kafka/assets/17735366/53d23168-f173-4f5a-b360-36928e913c27)

### Run the webserver

You can run the webserver using the following command from the project's root directory
```bash
go run *.go
```


### Curl command for the API server

Run the following curl command to publish a message to the kafka server
```bash
curl -XPOST  -H "Content-Type: application/json" --data '{"name": "athavan", "height": 7.3, "favouriteFood": "apple", "email": "akan@google.com"}' http://localhost:7070/publish
```
![messages](https://github.com/kanapuli/github.com/kanapuli/http-to-kafka/assets/17735366/7b32657b-f1d6-44a1-a3a4-ebb561b317ad)


## API response

- Returns *HTTP500* when there is an error while publishing to kafka
- Returns *HTTP200* when the message publishing to kafka is successful
