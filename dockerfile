FROM golang:latest

LABEL "com.mockapi.vendor"="Khakpouri"
LABEL version="1.0"
WORKDIR /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main .

EXPOSE 8080

CMD [ "/app/main" ]