# GO Expert - Open Telemetry

## Run locally

In the project root execute:
```shell
docker-compose up -d

curl -X POST -H "Content-Type: application/json" -d '{"cep": "70083900"}' http://localhost:8080/weather
```

Go to http://localhost:9411/
Click on `Run Query`

## Test on the cloud

```shell
curl -X POST -H "Content-Type: application/json" -d '{"cep": "70083900"}' http://34.69.76.102:8080/weather
```

Go to http://34.69.76.102:9411/
Click on `Run Query`

## APIs

### POST /weather

200:
```json
Request Body:
{"cep":"12345678"}

Response Body:
{"city":"somewhere","temp_C":16,"temp_F":60.8,"temp_K":289.1}
```

422:
```
invalid zipcode
```

404:
```
can not find zipcode
```

## URLs

| Service    | URL                    | GCP Project               |
| ---------- | ---------------------- | ------------------------- |
| Jaeger     | http://localhost:16686 | http://34.69.76.102:16686 |
| Prometheus | http://localhost:9090  | -                         |
| Zipkin     | http://localhost:9411  | http://34.69.76.102:9411  |
| Webserver  | http://localhost:8080  | http://34.69.76.102:8080  |

 