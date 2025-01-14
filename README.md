## Go-Auth

This is a Go authentication project. It uses various packages and modules to provide a complete authentication solution.

### Project Structure

The project has the following directory and file structure:

```
.gitignore
cmd/
    main.go
config/
    config.go
controllers/
    auth_controller.go
docker-compose.yml
Dockerfile
go.mod
go.sum
middleware/
    jwt_middleware.go
models/
    user.go
repository/
    user_repository.go
services/
    auth_service.go
utils/
    port.go
    validations.go
```

### Dependencies

This project uses several dependencies, including:

- `github.com/golang-jwt/jwt/v4` for JWT token generation and validation.
- `github.com/labstack/echo/v4` for HTTP request handling and routing.
- `gorm.io/gorm` for data persistence and database operations.
- `github.com/mattn/go-sqlite3` for SQLite database support.

You can see all the dependencies in the `go.sum` file.

### How to Run the Project

To run this project, follow these steps:

1. Make sure you have Go and Docker installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.

#### Using Docker

Build and run the project using Docker Compose:
```bash
docker-compose up --build
```

#### Without Docker

Install the dependencies:
```bash
go get -u ./...
```

Update the dependencies:
```bash
go mod tidy
```

Run the project:
```bash
go run cmd/main.go
```

### How to Run Tests

To run the tests for this project, follow these steps:

1. Make sure you have Go installed on your machine.
2. Navigate to the project directory.
3. Run the tests using the following command:
```bash
go test ./...
```

### How to Contribute

If you want to contribute to this project, please fork the repository, make your changes, and then submit a pull request.

### License

This project is licensed under the terms of the MIT license.
