# 1. TRACKING SERVICE

## GOLANG
Golang’s combination of efficient concurrency, low memory overhead, fast execution, and 
scalability makes it a good choice for developing high-throughput systems. It is also the 
language I have worked with the most and I am most confident in.

## POSTGRESQL
I chose PostgreSQL for storing accounts and events data because it can efficiently handle large 
volumes of data and is easily scalable. Additionally, PostgreSQL is the database I am most 
familiar with.
REDIS
To reduce server-to-database calls, I chose to use Redis. When checking an account's status 
(active, not active) on every request, the account data is cached in Redis, reducing database 
load and saving time.

# 2. PUB/SUB SYSTEM

## KAFKA
Kafka serves as the central messaging system, acting as a distributed event streaming platform. 
I chose Kafka because it is designed to handle large volumes of data and can easily scale 
horizontally by adding more brokers. This makes it a perfect choice for the pub/sub system 
where the system needs to propagate events at high rates. Kafka is also known for its durability, 
as it stores data on disk and can replicate it across multiple brokers which ensures no data lose 
in the event of failure. This aligns well with the fault-tolerant requirement of the CLI client.

# 3. CLIENT
The CLI Client, built in Go, is designed to pool Kafka events and print them. It has a functionality 
which enables you to filter and print events propagated by a specific account, which can be 
changed at runtime

# PROJECT SETUP

```sh
docker-compose up
```

```sh
make client
```