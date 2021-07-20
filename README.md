# udp-server-client
Simple UDP server and client written in golang for testing connections.

## How to Use

```
docker build -t udp .
docker run -it -e SERVER_PORT=7777 --name server udp
docker run -it --entrypoint /go/bin/udp-client -e CLIENT_CONNECTION_STRING="server:7777" --link server:server udp
```