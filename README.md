# consul-companion

## About
Компаньон взаимодействует с API Hashicorp Consul и позволяет регистрировать и дерегистрировать сервисы.


# Run

```
go run main.go api deregister --target "node-name"

go run main.go sd watch --conf-dir "./temp/consul.d" --search "./temp/projects"
```

## ENV for API
```
export CONSUL_HTTP_ADDR="127.0.0.1:8500"
export CONSUL_HTTP_TOKEN="TOKEN"
export CONSUL_HTTP_SCHEME="https" (default:http)
```

## ENV for SD (service discovery)

Консул проверяет папки с проектами находя там KEY-VALUE по следующему паттерну:

Name pattern `EXT_SERVICE_PORT` in `.env`
```
EXT_NGINX_PORT=80
```

Консул создает и удаляет файлы с сервисами по следующему паттерну:
Name pattern `SERVICENAME-PROJECTNAME-ENVPROJECT.hcl`

Если вы хотите, чтобы консул не трогал ваши сервисы находящиеся в том же каталоге, что и динамически сгенерированные им сервисы, то необходимо именовать сервисы разделяя слова  `_` (аднерсором) и консул компаньон будет их игронировать.