FROM golang as builder

# Add Maintainer Info
LABEL maintainer="diksha_product_service"

ENV GO111MODULE=on

RUN mkdir /app

WORKDIR /go/src/diksha_product_service
ADD . /go/src/diksha_product_service

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o engine main.go

## Distribution
FROM alpine:latest

ARG APP_ENV
ENV env_state=$APP_ENV
RUN apk update && apk upgrade && apk add --no-cache tzdata &&\
    mkdir /app

EXPOSE 3000
COPY --from=builder /go/src/diksha_product_service/engine /app
COPY --from=builder /go/src/diksha_product_service/Environment /go/src/diksha_product_service/Environment
COPY --from=builder /go/src/diksha_product_service/Services/Migration /go/src/diksha_product_service/Services/Migration
COPY --from=builder /go/src/diksha_product_service/log /go/src/diksha_product_service/log

CMD /app/engine -env=$env_state