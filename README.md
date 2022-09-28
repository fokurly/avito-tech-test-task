# avito-tech-test-task

### Сервис для взаимодействия с балансами пользователей

## Инструкция по запуску
Проект запускается с помощью docker-compose. Для этого перейдите
в корневую папку проекта и в консоли запустите команду:
```
    docker-compose up
```
Api апускается на 4000 порту по умолчанию
## Взаимодействие
После того, как проект запустится вам доступны следующие запросы

##POST
```http request
 http://localhost:4000/increase_client_balance
```
##`/increase_client_balance`
`/increase_client_balance` принимает в теле запроса json структуру вида:
```json
{
    "client_id": 10,
    "money":1000
}
```
`client_id` - id по которому будет создан баланс с изначальной суммой 0 и на неё же будет 
зачислена сумму денег в виде размера `money`
После того, как мы пополнили баланс какого-то из пользователей, мы можем сделать запрос на
снятие денег, показ баланса или перевод другому пользователю.

```http request
 http://localhost:4000/decrease_client_balance
```
##`/decrease_client_balance`
`/decrease_client_balance` - принимает json структуру вида

```json
{
    "client_id": 10,
    "money":1000
}
```
`client_id` - пользователь с которого снимут деньги <br>
`money` - сумма денег
##`/transfer_money`

```http request
 http://localhost:4000/transfer_money
```
`/transfer_money` - принимает в теле запроса json структуру вида
```json
{
    "sender_id": 10,
    "recipient_id":11, 
    "transfer_amount":500
}
```
`sender_id` - id отправителя <br>
`recipient_id` - id получателя <br>
`transfer_amount` - сумма перевода <br>
При переводе на несуществующий id пользователя, счёт клиента будет создаваться,
так как по условия мы создаём счёт при пополнении баланса, а перевод = пополнение
##GET
##`/get_client_balance`
```http request
 http://localhost:4000/get_client_balance
```
Принимает обязательный `query` параметр `client_id` по которому возвращает
сумму денег в рублях на счёту. Может принимать необязательный параметр `currency`, чтобы конвертировать и показать клиенту
его деньги в другой валюте. <br>
###Обычный запрос и ответ:
####Запрос
```http request
http://localhost:4000/get_client_balance?client_id=10
```

####Ответ
```json
{
    "client_balance (RUB)": 2000
}
```

###Запрос с изменением валюты и ответ:
####Запрос
```http request
http://localhost:4000/get_client_balance?client_id=10&currency=USD
```

####Ответ
```json
{
    "client_balance (USD)": 34.29
}
```


```http request
 http://localhost:4000/get_all_clients
```

##`/get_all_clients` 
возвращает всех пользователей, которые есть в сервисе<br>
Пример ответа:
```json
[
  {
    "client_id": 10,
    "money": 2000
  },
  {
    "client_id": 11,
    "money": 1000
  }
]
```