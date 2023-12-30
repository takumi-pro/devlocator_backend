# ============ develop ===============
FROM golang:1.21.4-bullseye as dev

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install git -y

COPY ./go.mod ./go.sum /go/src/app/
RUN go mod download

CMD ["air", "-c", ".air.toml"]

# ============ build ===============
FROM golang:1.21.4-bullseye as build

WORKDIR /go/src/app

RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install git -y
COPY go.mod go.sum ./
RUN go mod download
COPY . /go/src/app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# ============ production ===============
FROM gcr.io/distroless/base-debian11 AS production

WORKDIR /go/src/app

EXPOSE 8000

COPY --from=build /main /main

ENTRYPOINT ["/main"]