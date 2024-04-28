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

# Create a user with a known UID/GID within range 10000-20000.
# This is required by Choreo to run the container as a non-root user.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10014 \
    "choreo"
# Use the above created unprivileged user
USER 10014

WORKDIR /app/src

EXPOSE 9020

CMD ["./ruleengine"]

