# syntax=docker/dockerfile:1



##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /cart-service


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /cart-service /cart-service

EXPOSE 50051

USER nonroot:nonroot

CMD [ "/cart-service" ]