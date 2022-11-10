# syntax=docker/dockerfile:1

FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN ls

RUN go build ./cmd/main/main.go

EXPOSE 8000

CMD [ "./main" ]