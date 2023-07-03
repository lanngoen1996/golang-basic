# golang-basic

golang basic learning

### Folder Structure

```
api
|
├── .env
├── Dockerfile.local
├── docker-compose.yml
├── controllers
    └── users.go
├── models
    └── users.go
├── database
    └── migration.go
├── start
    └── routes.go
    └── routes
        └── users.go
├── validators
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
