# All-day

This APi, created to testing, study and try new tips in golang.

## Structure

- `cmd/`: Application entry points.
- `pkg/`: Library code for external use.
- `internal/`: Private application code.
- `go.mod`: Module dependencies.
- `README.md`: Project documentation.

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
