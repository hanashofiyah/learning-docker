FROM golang:1.22.3

WORKDIR /myapp

COPY go.mod go.sum ./

COPY . .

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /project-my-compose

EXPOSE 77

CMD ["/project-my-compose"]