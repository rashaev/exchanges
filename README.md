# exchanges

Service to exchange banknotes and coins.

Request example
```bash
curl -X POST --header "Content-Type: application/json" --data '{"amount":500,"banknotes":[200,100]}' http://<IP>:<PORT>/exchange
```


Answer example
```bash
{
  "exchanges": [
    [
      200,
      200,
      100
    ],
    [
      200,
      100,
      100,
      100
    ],
    [
      100,
      100,
      100,
      100,
      100
    ]
  ]
}
```



## Deployment

To run this project

```bash
  go run cmd/main.go
```


## Environment Variables

To run this project, you will need to set the following environment variables

`EXCHANGES_SERVER_ADDR` address and port to listen on. Default :8000

`EXCHANGES_LOGGER_LEVEL` log level. One of Debug, Info, Warn, Error. Default Info.

