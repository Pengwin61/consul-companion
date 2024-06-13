# consul-companion

# About
Компаньон взаимодействует с API Hashicorp Consul и позволяет регистрировать и дерегистрировать сервисы.




# Run

```
go run main.go --host="node-name"
```

#

```
export CONSUL_HTTP_ADDR="127.0.0.1:8500"
export CONSUL_HTTP_TOKEN="TOKEN"

export CONSUL_HTTP_SCHEME="https" (default:http)

```