FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -o ./ ./cmd/williamsLab/
RUN chmod +x williamsLab

EXPOSE 80

# CMD ["./williamsLab", "migrate"]
CMD ["./williamsLab", "serve", "--http=0.0.0.0:80"]
