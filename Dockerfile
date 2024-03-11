FROM golang:1.21

WORKDIR /app

COPY go.mod/ ./
RUN go mod download

COPY . ./

RUN go build -o /pricefetcher

EXPOSE 3000

CMD [ "/pricefetcher" ]