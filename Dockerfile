# Stage build
FROM golang:1.22-alpine AS build

WORKDIR /build

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 go test -tags=unit  ./...
RUN go build -o ./main

# Stage Deploy
FROM alpine:edge
WORKDIR /app
COPY --from=build /build/main ./main

CMD ["/app/main"]