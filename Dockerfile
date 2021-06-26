#build stage
FROM golang:alpine AS builder
WORKDIR /app

RUN chmod +x /app
COPY . ./
RUN go mod download

#final stage
FROM scratch
COPY --from=builder ./ /app
WORKDIR /app

LABEL Name=widgethubbackend Version=0.0.1

# to add a service, add a port
EXPOSE 8080

# change this to test individually
CMD ["go", "run", "main.go", "all"]