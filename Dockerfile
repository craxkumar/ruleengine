FROM golang:1.20-alpine AS build-env


WORKDIR /app

COPY . .

WORKDIR /app/src


RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ruleengine .

FROM alpine:latest  

WORKDIR /app/

COPY ./ ./

COPY --from=build-env /app/src/ruleengine /app/src/ruleengine 

WORKDIR /app/src

EXPOSE 9020

CMD ["./ruleengine"]

