## Сервис обработки сообщений

сборка : `make build`  

запуск : `./main`

### API:

`api/v1/message/:id` GET получение сообщения по id  

`api/v1/message/` POST отправка сообщения

### Команды для запуска kafka
* `./zookeeper-server-start.sh ../config/zookeeper.properties`
* `./kafka-server-start.sh ../config/server.properties`
* для подписки на топик `./kafka-console-consumer.sh  --topic my-topic --bootstrap-server localhost:9092 --consumer-property auto.offset.reset-earliest`