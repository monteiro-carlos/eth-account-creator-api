# Ethereum Account Creator API

It's an API that helps you create an EOA Ethereum account by generating private and public keys and also returns your EOA address from a public key.

## Running Local

Clone the project

```bash
  git clone https://github.com/monteiro-carlos/eth-account-creator-api.git
```

Go to the project directory

```bash
  cd eth-account-creator-api
```

Install dependencies

```bash
  go mod download
```

```bash
  make swagger 
```

Start the server

```bash
  go run main.go
```

## Running Docker

Start the API

```bash
  make docker-up
```

## API Docs

- After running project, the swagger documentation is available at: [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)

## Metrics

- You can access metrics at [http://localhost:5000/metrics](http://localhost:5000/metrics)

## Authors

- [@monteiro-carlos](https://www.github.com/monteiro-carlos)

## References

- [Golang Standard Project Layout](https://github.com/golang-standards/project-layout)
- Onion Architecture
- Microservices Architecture
