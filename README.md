# ruleengine prototype

This is a go-micro prototype generated using WeDAA, you can find documentation and help at [WeDAA Docs](https://www.wedaa.tech/docs/introduction/what-is-wedaa/)

## Prerequisites

- go version >= 1.20

## Project Structure

```
├── docker/ (contains docker compose files for external components based on architecture design)
├──src
    ├── config/ (configuration properties and configuration properties loader)
    ├── controllers/ (api controllers)
    ├── go.mod
    └── main.go
├── Dockerfile (for packaging the application as docker image)
├── README.md (Project documentation)
├── comm.yo-rc.json (generator configuration file for communications)
```

## Get Started

Install required dependencies: `go mod tidy`

Run the prototype locally: `go run .`

Open [http://localhost:9020/management/health/readiness](http://localhost:9020/management/health/readiness) to view it in your browser.

The page will reload when you make changes.

## Containerization

Build the docker image: `docker build -t ruleengine:latest .`

Start the container: `docker run -d -p 9020:9020 ruleengine:latest`
