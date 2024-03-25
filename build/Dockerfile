FROM golang:1.19

ENV APP_DIR = /app 

RUN mkdir $APP_DIR 
WORKDIR $APP_DIR 

ENV PORT=8080
EXPOSE $PORT 

COPY . .

RUN go mod tidy 
RUN go mod vendor

RUN go build -o pismo_service cmd/pismo/main.go

CMD ["./pismo_service"]