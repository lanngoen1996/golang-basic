# golang-basic

golang basic learning

### Folder Structure

```
api
|
├── .env
├── Dockerfile.local
├── cors

    └── app.go

    └── database.go

    └── env.go

    └── migration.go
├── service

    └── users
        └── route.go
        └── controller.go
        └── model.go
        .
        .
        .

├── docker-compose.yml
.
.
.
.

└── README.md
```

### How to start project

-- copy /api/.env.example => /api/.env

rootPath run command

```
docker compose up --build
```

API Listening on 0.0.0.0:3000 or http://localhost:3000
