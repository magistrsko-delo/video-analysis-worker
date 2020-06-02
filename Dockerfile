FROM golang:alpine

RUN apk upgrade -U \
 && apk add ca-certificates ffmpeg libva-intel-driver \
 && rm -rf /var/cache/*

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .
COPY .live.env .
RUN cp .live.env .env && mkdir assets && mkdir assets/chunks
COPY magisterij-6d3594ec69ea.json .
RUN export GOOGLE_APPLICATION_CREDENTIALS="magisterij-6d3594ec69ea.json"

CMD ["/dist/main"]