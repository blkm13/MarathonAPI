FROM golang:1.16

WORKDIR /usr/appl
COPY go.mod .

RUN go get -u github.com/jackc/tern
RUN go get github.com/gin-gonic/gin

COPY . .
RUN go build src/main.go

EXPOSE 8080


RUN ls -latr
RUN chmod 755 entrypoint.sh
# CMD ls -latr
ENTRYPOINT ["./entrypoint.sh"]
# CMD go run src/main.go