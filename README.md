# promotion-app
Simple CRUD app
## Requirements
- go 1.19 +
- Docker/Docker-compose
- postgresql 12+
## Launch
- You can run the service manually:
```
go run main.go
```
or
```
go build -o ./prom && ./prom
```
In this case you have to set `PROFILE` env variable:
```
export PROFILE=local
```
or you can use docker-compose:
```
 cd docker && docker-compose up 
```
## API
### GET /promotions - get all promotions
```
curl --location --request GET 'http://localhost:8080/promotions' \
--header 'Content-Type: application/json'
```
### GET /promotions/{id} - get a promotion by id
```
curl --location --request GET 'http://localhost:8080/promotions/1' \
--header 'Content-Type: application/json'
```
### POST /promotions - add new promotion
```
curl --location --request POST 'http://localhost:8080/promotions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "originalId":     "ec15cb91-74ab-4e06-988b-18de33629746",
	"price":          222.1,
	"expirationDate": "2022-07-06 14:32:33.600000 +00:00"
}'
```
### POST /promotions/file - upload promotions using CSV file
```
curl --location --request POST 'http://localhost:8080/promotions/file' \
--form 'file=@"/Downloads/promotions.csv"'
```
### DELETE /promotions - delete all promotions
```
curl --location --request DELETE 'http://localhost:8080/promotions' \
--header 'Content-Type: application/json' 
```
