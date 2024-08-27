# TRACKING SYSTEM

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