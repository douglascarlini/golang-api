
# Golang API
RESTful API in Golang with JWT Authentication and Middlewares for Logging and App Access Control.
> :warning: This is an studies purpose project (in progress).

## Requirements

- [Golang](https://go.dev/doc/) 1.20
- [Docker](https://www.docker.com/)

### Configuration
Copy `.env-example` to `.env` and change as needed.

### Run Local
- Enter to `api` folder
- Download deps: `go get .`
- Run application: `go run .`

### Deploy
Run `docker compose up --build -d` on terminal.

### Endpoints
> Swagger UI available on `http://{HOST}:{DOC_PORT}`.
