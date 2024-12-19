# calc_service

Этот проект представляет собой веб-сервис, который принимает арифметическое выражение через HTTP POST-запрос и возвращает результат его вычисления.

## Инструкция по запуску

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
   
2. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/sticker2/calc_service
   cd calc_service

   ```

3. Запустите сервер:

   ```bash
   go run ./cmd/main.go
   ```

   После этого сервер будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

## Примеры использования

Для взаимодействия с сервером отправляйте POST-запросы. Для этого вы можете использовать программу Postman или команду `curl`.

### Пример успешного запроса:

Отправьте POST-запрос с телом:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

Ответ:

```json
{
  "result": "6"
}
```

### Ошибка 422 (невалидное выражение):

Если выражение содержит недопустимые символы, сервер вернет ошибку 422:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+a"
}'
```

Ответ:

```json
{
  "error": "Expression is not valid"
}
```

### Ошибка 500 (внутренняя ошибка сервера):

В случае возникновения внутренней ошибки, например, деления на ноль, сервер вернет ошибку 500:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2/0"
}'
```

Ответ:

```json
{
  "error": "Internal server error"
}
```
