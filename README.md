# Todo API Service

Учебный REST API на Go для управления пользователями и задачами.

Проект реализован с использованием роутера `go-chi` и содержит Swagger-документацию, сгенерированную с помощью `swaggo`.

API позволяет:

- создавать, получать, обновлять и удалять пользователей;
- создавать, получать, обновлять и удалять задачи;
- просматривать описание API через Swagger UI;
- тестировать REST endpoints через Swagger UI или `curl`.

---

## Используемые технологии

- Go
- go-chi
- Swaggo
- Swagger UI
- REST API
- JSON

---

## Требования

Для запуска проекта необходимо установить:

- Go версии 1.22 или выше;
- Git;
- `swag` CLI для генерации Swagger-документации.

Проверить установленную версию Go:

```bash
go version
```

Установить `swag` CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Если команда `swag` недоступна после установки, необходимо добавить Go binaries в `PATH`:

```bash
echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.bashrc
source ~/.bashrc
```

Проверить установку:

```bash
swag --version
```

---

## Установка проекта

Клонировать репозиторий:

```bash
git clone https://gitverse.ru/netology/ai-assist-it.git
```

Перейти в папку проекта:

```bash
cd ai-assist-it
```

Загрузить зависимости:

```bash
go mod download
```

---

## Основные зависимости

Основные зависимости проекта указаны в `go.mod`:

- `github.com/go-chi/chi/v5 v5.0.10` — HTTP-роутер;
- `github.com/swaggo/http-swagger/v2 v2.0.2` — Swagger UI для Go HTTP-сервера;
- `github.com/swaggo/swag v1.16.2` — генерация Swagger-документации.

Дополнительные зависимости устанавливаются автоматически через Go Modules.

---

## Генерация Swagger-документации

Swagger-аннотации находятся в обработчиках внутри:

```text
internal/handlers/
```

Документация моделей находится в:

```text
internal/models/
```

Для генерации Swagger выполнить:

```bash
swag init -g main.go -o docs
```

После успешной генерации появятся файлы:

```text
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```

---

## Запуск приложения

Запустить сервер:

```bash
go run main.go
```

После запуска сервер будет доступен по адресу:

```text
http://localhost:8080
```

Swagger UI:

```text
http://localhost:8080/swagger/
```

Swagger JSON:

```text
http://localhost:8080/docs/swagger.json
```

---

## API

В проекте реализованы endpoints для пользователей и задач.

### Users

| Метод | Endpoint | Описание |
|---|---|---|
| GET | `/api/v1/users/` | Получить список пользователей |
| GET | `/api/v1/users/{id}` | Получить пользователя по ID |
| POST | `/api/v1/users/` | Создать пользователя |
| PUT | `/api/v1/users/{id}` | Обновить пользователя |
| DELETE | `/api/v1/users/{id}` | Удалить пользователя |

### Todos

| Метод | Endpoint | Описание |
|---|---|---|
| GET | `/api/v1/todos/` | Получить список задач |
| GET | `/api/v1/todos/{id}` | Получить задачу по ID |
| POST | `/api/v1/todos/` | Создать задачу |
| PUT | `/api/v1/todos/{id}` | Обновить задачу |
| DELETE | `/api/v1/todos/{id}` | Удалить задачу |

---

## Примеры запросов

### Создание пользователя

Запрос:

```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User"}'
```

Пример ответа:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "test@example.com",
    "name": "Test User",
    "created_at": "2026-10-03T12:01:29Z"
  }
}
```

---

### Получение пользователя

Запрос:

```bash
curl http://localhost:8080/api/v1/users/1
```

Пример успешного ответа:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "test@example.com",
    "name": "Test User"
  }
}
```

---

### Создание задачи

Запрос:

```bash
curl -X POST http://localhost:8080/api/v1/todos/ \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"title":"Изучить Swagger","description":"Завершить домашнее задание","done":false}'
```

Пример ответа:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Изучить Swagger",
    "description": "Завершить домашнее задание",
    "done": false
  }
}
```

---

### Получение задачи

Запрос:

```bash
curl http://localhost:8080/api/v1/todos/1
```

Пример ответа:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Изучить Swagger",
    "description": "Завершить домашнее задание",
    "done": false
  }
}
```

---

### Обновление задачи

Запрос:

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Swagger готов","done":true}'
```

Пример ответа:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Swagger готов",
    "description": "Завершить домашнее задание",
    "done": true
  }
}
```

---

### Удаление задачи

Запрос:

```bash
curl -i -X DELETE http://localhost:8080/api/v1/todos/1
```

Успешный ответ:

```text
HTTP/1.1 204 No Content
```

---

## HTTP-коды

В API используются следующие основные HTTP-коды:

- `200 OK` — запрос успешно выполнен;
- `201 Created` — ресурс успешно создан;
- `204 No Content` — ресурс успешно удалён;
- `400 Bad Request` — переданы некорректные данные;
- `404 Not Found` — пользователь или задача не найдены.

---

## Структура проекта

```text
ai-assist-it/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── internal/
    ├── handlers/
    │   ├── todo.go
    │   └── user.go
    └── models/
        ├── todo.go
        └── user.go
```

### Назначение основных файлов

- `main.go` — точка входа приложения и настройка маршрутов;
- `internal/handlers/` — HTTP-обработчики и Swagger-аннотации;
- `internal/models/` — структуры данных и документация моделей;
- `docs/` — автоматически сгенерированная Swagger-документация;
- `go.mod` — описание Go-модуля и зависимостей;
- `go.sum` — контрольные суммы зависимостей;
- `README.md` — документация проекта.

---

## Чек-лист проекта

- [x] Реализованы REST endpoints для пользователей
- [x] Реализованы REST endpoints для задач
- [x] Добавлены Swagger-аннотации ко всем handlers
- [x] Добавлены docstrings к моделям
- [x] Сгенерирован `docs/swagger.json`
- [x] Сгенерирован `docs/swagger.yaml`
- [x] Swagger UI успешно запускается
- [x] В Swagger UI отображаются все endpoints
- [x] API протестирован через `curl`
- [x] Проверены POST-запросы
- [x] Проверены GET-запросы
- [x] Проверены PUT-запросы
- [x] Проверены DELETE-запросы
- [x] Проект успешно собирается через Go

---

## Swagger UI

После запуска приложения интерактивная документация доступна по адресу:

```text
http://localhost:8080/swagger/
```

Swagger позволяет:

- просматривать все endpoints;
- видеть HTTP-методы;
- просматривать параметры запросов;
- просматривать модели данных;
- видеть возможные HTTP-ответы;
- выполнять запросы прямо из браузера.

---

## Лицензия

Проект выполнен в учебных целях.

Для проекта может использоваться лицензия MIT.

---

## Контакты

GitHub: `RavenDark322`

Проект выполнен в рамках учебного задания по документированию REST API с использованием Swagger и Swaggo.
