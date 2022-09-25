# wallet-api

This project is an Event Sourcing use case sample which was built on top of [Event Horizon](https://github.com/looplab/eventhorizon) toolkit. 

## Stack

- [Golang](https://go.dev/)
- [MongoDB](https://www.mongodb.com/)
- [Pub/Sub](https://cloud.google.com/pubsub)

## Execution

There are two different simple ways to run and stop the application:

- By Makefile: `make start` | `make stop`
- By Docker:   `docker compose up -d --build` | `docker compose down -v`

### Migrations

To apply database migrations, run `make db-migrate` informing the database URI. Example:

```sh
make MONGO_URL="mongodb://mongo1:30001,mongo2:30002,mongo3:30003/balance?replicaSet=my-replica-set" migrate
```

## API

| Method |          Resource           | 
|:------:|:---------------------------:|
|  POST  |          /wallets           | 
| PATCH  | /wallets/{wallet-id}/credit | 
| PATCH  | /wallets/{wallet-id}/debit  | 

Import the [postman collection](/docs/wallet-api.json) for more details.

