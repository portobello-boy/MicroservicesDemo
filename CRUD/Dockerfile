FROM golang:1.16

WORKDIR /home/danielmillson/Documents/Projects/Git Projects/MicroservicesDemo/CRUD/CRUD

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["CRUD"]

EXPOSE 3000