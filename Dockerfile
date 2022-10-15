FROM golang:1.19.2-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN apk add --no-cache git && go build -tags=jsoniter -o superiad . && apk del git

FROM alpine
WORKDIR /app
COPY --from=builder /app/superiad .
CMD [ "./superiad" ]