# TRACKING SYSTEM
The tracking system project comprises three key components:
- server
- kafka
- client
The server sends events to Kafka, and the client retrieves these events, displaying them in the console.

## PROJECT SETUP

DOCKER COMPOSE:
 - postgres
 - kafka
 - zookeeper
 - redis
 - client
 - server
```sh
docker-compose up
```

RUN MIGRATIONS
```sh
make migrateup
```

CLIENT
```sh
make client
```