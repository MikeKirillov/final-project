# Файлы для итогового задания
В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

# Call examples
/api/nextdate GET call:
```
$ curl http://localhost:7540/api/nextdate?now=20250126&date=20250126&repeat=y
```

/api/task POST call:
```
$ curl -X POST http://localhost:7540/api/task \
-H "Content-Type: application/json" \
-d '{"date":"20240201","title":"some title","comment":"some comment","repeat":"y"}'
```

/api/task?id=... GET call:
```
$ curl http://localhost:7540/api/task?id=25
```

/api/task PUT call:
```
$ curl -X PUT http://localhost:7540/api/task \
-H "Content-Type: application/json" \
-d '{"id":"204","date":"20240201","title":"some title","comment":"some comment","repeat":"y"}'
```