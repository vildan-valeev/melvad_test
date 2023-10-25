
# Тестовое задание  Golang melvad.com
Необходимо создать консольное приложение-сервис, которое принимает HTTP POST запросы:

---

* по пути POST `/redis/incr` с json вида
```json
{
  "key": "age",
  "value": 19
}
```
подключается к Redis DB (хост и порт указываются при запуске в параметрах `-host` и `-port`),
инкрементирует значение по ключу, указанному в "key" на значение из "value", и возвращает его в
виде:
```json
{
  "value": 20
}
```
---
* по пути POST `/sign/hmacsha512` с json вида
```json
{
  "text": "test",
  "key": "test123"
}
```
и возвращает HMAC-SHA512 подпись значения из "text" по ключу "key" в виде hex строки

---
* по пути POST `/postgres/users` с json вида
```json
{
  "name": "Alex",
  "age": 21
}
```
создает в базе данных PostgreSQL таблицу users, если она не существует, добавляет в нее строку
("Alex", 21) и возвращает идентификатор добавленного пользователя в виде
```json
{
  "id": 1
}
```

---

Дополнительные условия:
* можно использовать любые библиотеки для работы с http, Redis DB и PostgreSQL;
* приложение должно быть протестировано unit-тестами (любой тестовый фреймворк);
* результат нужно разместить на github;
* наибольшее внимание следует уделить качеству коду.
