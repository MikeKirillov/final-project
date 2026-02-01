# Файлы для итогового задания
В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

add task POST call example:
```
curl -X POST http://localhost:7540/api/task \
-H "Content-Type: application/json" \
-d '{"date":"20240201","title":"some title","comment":"some comment","repeat":"y"}'
```