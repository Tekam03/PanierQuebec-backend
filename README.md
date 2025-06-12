# Panier Quebec

A backend service built with Go, `sqlc` and `connectrpc`

## 🧱 Architecture Overview

```
cmd/
└── server/       # App entrypoint (main.go)
internal/
├── db/           # sqlc-generated database access
│   ├── queries/  # SQL files for sqlc
│   └── db.go     # sqlc-generated code
├── repo/         # Repository interfaces & sqlc-based implementations
├── service/      # Business logic
├── handler/      # gRPC/ConnectRPC handlers
```

## ⚙️ Tech Stack

| Layer     | Tool/Library                            | Purpose                                  |
|-----------|------------------------------------------|------------------------------------------|
| Database  | PostgreSQL                               | Persistent store                         |
| Queries   | [sqlc](https://sqlc.dev)                 | Type-safe SQL access                     |
| API       | [ConnectRPC](https://connect.build)      | Unified gRPC + HTTP/1.1 + HTTP/2 support |




##  Set up 

```bash
cp .env.example .env
# Update DATABASE_URL in the .env file
```

### Install dependencies

```bash
go mod tidy
```

### Generate SQL code

```bash
sqlc generate
```

### Run the server

```bash
go run ./cmd/server
```

Server will start on:

- gRPC: `localhost:8080`
- HTTP (Connect): `localhost:8081`




## API Overview

Using [ConnectRPC](https://connect.build):

- Protobuf definitions are located in `proto/`.
- ConnectRPC supports HTTP/1.1, HTTP/2, gRPC, and gRPC-Web.



## Testing

To run all tests:

```bash
go test ./...
```

You can mock repository interfaces in unit tests to isolate service logic.

## DB Migrations
Use [migrate](https://github.com/golang-migrate/migrate) for database migrations.


To create a new migration, run:

```bash
migrate create -ext sql -dir migrations -seq <migration_name>
```
This will create a new migration file in the `migrations` directory.

To apply migrations, ensure you have the `DATABASE_URL` set in your `.env` file, then run:

```bash
migrate -path migrations -database "$DATABASE_URL" up
```
