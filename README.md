# backend

### Как поставить:
```bash
pip install go (см. ОС)
pip install mysql-server (см. ОС)
make deps
```

### Куда клонить:
```bash
юзер/go/src
```

### Как запустить:
```bash
make start -> запуск
make daemon_start -> запуск в демоне
```

### Вкратце что такое:
```bash
Список всех постов, есть пагинация
GET localhost:8080/api/posts?from=20&count=20

Добавление поста
POST localhost:8080/api/posts { "url": "http://lorempixel.com/400/200/sports/"}
```
