# syntax=docker/dockerfile:1
# setup build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# prepare base runtime and lib dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy and build sources
COPY *.go ./
RUN go build -o /pwgen

# setup run stage
FROM golang:1.19-alpine AS runner

WORKDIR /

# copy binary over, listen to port and run the app
COPY --from=builder /pwgen /pwgen
EXPOSE 8080
CMD [ "/pwgen" ]
