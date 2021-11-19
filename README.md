# twirpdemo

This repo was created for a local meetup to show some of the features of the [Twirp RPC framework](https://twitchtv.github.io/twirp/docs/intro.html).

## Usage

Generate proto code:

```
protoc  --twirp_out=. --go_out=. proto/rpc.proto
```

Run the calculator service:

```
$ go run cmd/calculatorsvc/main.go
2021/11/19 13:13:43 server started at port 8080
```

Make a request using curl: 

```
$ curl \
  --request "POST" \
  --header 'Content-Type:application/json' \
  --data '{"a": 10, "b": 5}' \
  http://localhost:8080/twirp/Calculator/Add

{"result":15}⏎
````


Make a request using generated Go client: 

```
$ go run cmd/calculatorctl/main.go
Result = 15
````



