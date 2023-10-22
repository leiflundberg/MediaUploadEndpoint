FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify
COPY *.go ./
RUN go build -o /api
EXPOSE 8080
CMD [ "/api" ]
