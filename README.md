# All-day

This APi, created to testing, study and try new tips in golang.

## Structure

- `cmd/`: Application entry points.
- `pkg/`: Library code for external use.
- `internal/`: Private application code.
- `go.mod`: Module dependencies.
- `README.md`: Project documentation.

```
all-day/
│
├── cmd/
│     └── main.go
├── configs/
│     └── dbConfig.go
├── internal/
│     ├── delivery/
│     │     ├── handlers/
│     │     │     └── bookHandler/
│     │     │          └── bookHandler.go
│     │     ├── data/
│     │     │     ├── request/
│     │     │     │     └── bookReq/
│     │     │     │          └── bookRequest.go
│     │     │     └── response/
│     │     │          ├── bookRes/
│     │     │          │     └── bookResponse.go
│     │     │          └── response.go
│     │     └── router/
│     │           ├── bookRouter/
│     │           │          └── bookRouter.go
│     │           └── router.go
│     ├── domain/
│     │     ├── models/
│     │     │     └── books.go
│     │     ├── repositories/
│     │     │     ├── bookRepo/
│     │     │     │      └── bookRepo.go
│     │     └── services/
│     │           ├── bookService/
│     │           │      └── bookService.go
│     └── infrastructure/
│           └── database/
│                  ├── database.go
│                  └── migrations.go
├── pkg/
│     ├── utils/
│     │     └── base.go
│     └── helpers/
│           └── errorPanic.go
├── .env.example
├── .gitignore
├── go.mod
└── go.sum
```

## Getting Started

1. Install dependencies:

    ```sh
    go mod tidy
    ```

2. Run the application:

    ```sh
    go run main.go
    ```

3. Build the application:

    ```sh
    go build -o bin/myapp main.go
    ```

## Configuration

The configuration file (`config.yaml`) should be placed in the root directory:

```yaml
server_address: ":8080"
