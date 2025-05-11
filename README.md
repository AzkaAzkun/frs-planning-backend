# FRS Planning Backend

## Description

FRS Planning Backend is a Go-based server designed to handle the backend logic of your project. This server includes features like database migrations and data seeding for smooth setup and operation.

## How to Run

Run the application
To start the application normally, use:

```bash
go run main.go
```

Run with Migrations
To apply database migrations, run:

```bash
go run main.go --migrate
```

Run with Seeder
To seed the database with initial data, use:

```bash
go run main.go --seeder
```
