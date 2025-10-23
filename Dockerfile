FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./ ./cmd/williamsLab/
RUN chmod +x williamsLab

EXPOSE 80

CMD ["./williamsLab", "serve", "--dev", "--http=0.0.0.0:80"]
