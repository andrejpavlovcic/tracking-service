# TRACKING SYSTEM
The tracking system project comprises three key components:
- server
- kafka
- client
The server sends events to Kafka, and the client retrieves these events, displaying them in the console.

### ARCHITECTURE
![Posnetek zaslona 2024-08-26 140648](https://github.com/user-attachments/assets/7a9b99c1-9734-4e48-8664-63d523fb467b)

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
