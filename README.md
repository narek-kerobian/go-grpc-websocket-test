
## Instructions

### Run the application
- Run `docker-compose up` to build and run the container image

The application will migrate the in-memory database and generate mock data on
first run.   
Id's for generated users and wallets range `1..5`.

### Send queries

#### Http
Http requests can be sent from any client using provided curl requests in
[Example cURL requests](#Example-cURL-requests) section as an example.

#### Websocket and gRPC
Websocket and gRPC requests need to be ran from Postman or Insomnia using
provided collections (See [Supplementary information](#Supplementary-information) section)

## Tech stack
Applicaton language: go   
Http frawework: [gin gonic](https://gin-gonic.com/)   
Websocket implementation library:
[gorilla](https://github.com/gorilla/websocket)   
gRPC implementation library: native [grpc](https://pkg.go.dev/google.golang.org/grpc)   
In-memory database: [go-memdb](https://github.com/hashicorp/go-memdb)   
Reverse prox: [Caddy](https://caddyserver.com)   
Process control system: [supervisord](http://supervisord.org/)   
Containerization technology: docker/docker-compose   

## Supplementary information
Collections for Insomnia (postman alternative) can be found in `http` directory.
As I have used it as a client, it is the recommended way of running http,
websocket, and rpc requests.

### Example cURL requests

#### Get the balance for user with the id 1
```bash
curl --request GET \
  --url http://localhost:8080/api/wallet/balance/1
```

#### Augment the balance of wallet with id 1
```bash
curl --request POST \
  --url http://localhost:8080/api/wallet/deposit \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 1,
	"amount": 100
}'
```

#### Deduct the balance of the wallet with id 1
```bash
curl --request POST \
  --url http://localhost:8080/api/wallet/withdraw \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 1,
	"amount": 10
}'
```

