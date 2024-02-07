
# Wb l0

Тестовое задание


## Deployment


```bash
git clone https://github.com/NeSoLeK/wb_l0.git
```

```bash
cd wb_l0
```

```bash
docker-compose up

```



## Test data insertion

```bash
go run nats_check.go
```


## Tech Stack

* [Gin](https://github.com/gin-gonic/gin)
* [Stan.go](https://github.com/nats-io/stan.go)
* [PostgreSql](https://github.com/lib/pq)


## Tasks

В БД: 
* Развернуть локально postgresql
* Создать свою бд
* Настроить своего пользователя. 
* Создать таблицы для хранения полученных данных

В сервисе:
* Подключение и подписка на канал в nats-streaming
* Полученные данные писать в Postgres
* Так же полученные данные сохранить in memory в сервисе (Кеш)
* В случае падения сервиса восстанавливать Кеш из Postgres
* Поднять http сервер и выдавать данные по id из кеша
* Сделать простейший интерфейс отображения полученных данных, для их запроса по id
