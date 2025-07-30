
FROM golang:1.24-alpine AS builder


RUN apk update && apk add --no-cache git


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init


RUN CGO_ENABLED=0 GOOS=linux go build -o /itinerary-api ./main.go



FROM alpine:latest


RUN apk add --no-cache chromium udev ttf-freefont


WORKDIR /app


COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/docs ./docs


COPY --from=builder /itinerary-api .


EXPOSE 8080


CMD ["/app/itinerary-api", "--no-sandbox"]
