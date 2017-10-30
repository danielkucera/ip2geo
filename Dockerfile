FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN apt update && apt-get install -y \
	libgeoip-dev \
&& rm -rf /var/lib/apt/lists/*

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 8000

CMD ["go-wrapper", "run"] # ["app"]
