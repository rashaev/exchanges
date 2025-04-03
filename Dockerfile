FROM golang:1.24-bullseye AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


FROM base AS build

WORKDIR /app

COPY . .

RUN go build -o /out/exchanges cmd/main.go


FROM golang:1.24-bullseye AS release

WORKDIR /app

COPY --from=build /out/exchanges /app/

CMD ["/app/exchanges"]