# Система рекомендаций фильмов

Веб-приложение для управления фильмами и пользователями с простым REST API.

## Функциональность

- Регистрация и авторизация пользователей
- Добавление фильмов в базу данных
- Поиск фильмов по жанрам
- Веб-интерфейс для взаимодействия с системой

## Технологии

- **Backend**: Go 1.24
- **База данных**: PostgreSQL
- **Frontend**: HTML, CSS, JavaScript
- **Контейнеризация**: Docker (для PostgreSQL)

## Требования

- Go 1.24 или выше
- Docker и Docker Compose
- Git

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/highdolen/recommendation_system.git
cd recommendation_system
```

### 2. Настройка базы данных PostgreSQL

#### Запуск PostgreSQL в Docker:

```bash
# Создание и запуск контейнера PostgreSQL
docker run --name postgres-recdb \
  -e POSTGRES_DB=recdb \
  -e POSTGRES_USER=fedya \
  -e POSTGRES_PASSWORD=secret \
  -p 5433:5432 \
  -d postgres:15
```

#### Создание схемы базы данных:

```bash
# Подключение к контейнеру и создание таблиц
docker exec -i postgres-recdb psql -U fedya -d recdb < shema.sql
```

Или через psql:

```bash
# Подключение к базе данных
docker exec -it postgres-recdb psql -U fedya -d recdb

# Внутри psql выполнить содержимое shema.sql:
\i /path/to/shema.sql
```

### 3. Установка зависимостей Go

```bash
go mod download
```

### 4. Настройка подключения к базе данных

Проверьте настройки подключения в файле `db/connect.go`. По умолчанию используются:

- Host: `localhost`
- Port: `5433`
- User: `fedya`
- Password: `secret`
- Database: `recdb`

Если нужно изменить параметры подключения, отредактируйте строку подключения в `db/connect.go`.

### 5. Запуск приложения

```bash
go run main.go
```

Приложение будет доступно по адресу: http://localhost:8080

## API Endpoints

### Пользователи

- `POST /models/users` - Добавление нового пользователя
  
  ```json
  {
    "nickname": "username",
    "name": "User Name",
    "email": "user@example.com"
  }
  ```

- `POST /api/login` - Авторизация пользователя

### Фильмы

- `POST /models/films` - Добавление нового фильма
  
  ```json
  {
    "title": "Название фильма",
    "genre": "Жанр",
    "year": 2023
  }
  ```

- `GET /models/films/bygenre?genre=action` - Получение фильмов по жанру

## Структура проекта

```
recommendation_system/
├── db/
│   └── connect.go          # Подключение к базе данных
├── internal/
│   ├── api/
│   │   ├── api.go         # REST API обработчики
│   │   └── login.go       # Авторизация
│   └── models/
│       ├── film.go        # Модель фильма
│       └── user.go        # Модель пользователя
├── web/
│   ├── auth.html          # Страница авторизации
│   ├── auth.css
│   ├── auth.js
│   ├── films.html         # Страница фильмов
│   ├── films.js
│   └── film.css
├── main.go                # Точка входа
├── shema.sql             # Схема базы данных
├── go.mod
└── README.md
```

## Использование

1. Откройте браузер и перейдите на http://localhost:8080
2. Пройдите авторизацию на главной странице
3. После авторизации вы попадете на страницу управления фильмами
4. Добавляйте фильмы и фильтруйте их по жанрам

## Остановка и очистка

```bash
# Остановка приложения: Ctrl+C

# Остановка и удаление контейнера PostgreSQL
docker stop postgres-recdb
docker rm postgres-recdb
```

## Восстановление данных из бэкапа

Если у вас есть файл `backup.dump`, вы можете восстановить данные:

```bash
docker exec -i postgres-recdb pg_restore -U fedya -d recdb < backup.dump
```

## Разработка

Для разработки убедитесь, что:

1. PostgreSQL запущен и доступен
2. Схема базы данных создана
3. Go модули установлены

```bash
# Запуск в режиме разработки с автоперезагрузкой
go run main.go
```

## Проблемы и решения

**Ошибка подключения к базе данных:**
- Убедитесь, что PostgreSQL контейнер запущен
- Проверьте порт 5433 (не занят ли другим процессом)
- Проверьте правильность credentials в `db/connect.go`

**Порт 8080 занят:**
- Измените порт в `main.go` на свободный
- Или остановите процесс, использующий порт 8080

## Лицензия

МIT License

# 🎬 Recommendation System

Приложение для рекомендаций фильмов. Пользователи могут добавлять фильмы по жанрам, а затем получать список добавленных фильмов - независимо от автора. Используется PostgreSQL для хранения и обработки данных. Развёрнуто на Render. Адаптировано под мобильные устройства.

## 🚀 Стек технологий

- Go (Golang)
- PostgreSQL
- Render
