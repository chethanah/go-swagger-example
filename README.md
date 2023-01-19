# Swagger UI for example App
The example app here is simple addition and multiplication that takes two numbers as JSON input
We'll write the swagger definitions and serve the swagger ui

## How to build App
- Clone and build app
```
$ cd go-swagger-example 
$ go build 
$ ./go-swagger-example 
Starting calc API server on port :8080...
```
- Now we can GET/POST to this server on port 

## How to serve swagger UI
- Install `swag` - https://github.com/swaggo/swag/ 
  - Download binary from release https://github.com/swaggo/swag/releases and copy to PATH
- Installer `swagger` - https://github.com/go-swagger/go-swagger/
  - Download binary from release https://github.com/go-swagger/go-swagger/releases and copy to PATH
- See `main.go` for the swagger definitions above each handlers.
- Generate the swagger doc and serve 
```
$ cd go-swagger-example 
$ swag init 
2023/01/19 09:30:58 Generate swagger docs....
2023/01/19 09:30:58 Generate general API Info, search dir:./
2023/01/19 09:30:58 Generating main.InputNumbers
2023/01/19 09:30:58 Generating main.Sum
2023/01/19 09:30:58 Generating main.Prod
2023/01/19 09:30:58 create docs.go at  docs/docs.go
2023/01/19 09:30:58 create swagger.json at  docs/swagger.json
2023/01/19 09:30:58 create swagger.yaml at  docs/swagger.yaml
```
- This will create the `doc` dir with `swagger.json`
- Serve `swagger.json` using `swagger`
```
$ swagger serve --no-open --host <IP> --port 8022 docs/swagger.json 
2023/01/19 09:32:13 serving docs at http://IP:8022/docs

```