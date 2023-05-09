FROM golang:alpine


WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o online-course

EXPOSE 3001

ENTRYPOINT ["/app/online-course"]