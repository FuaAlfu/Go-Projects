---
stack: Go, sql, postger
lang: data migiration
---

# Database migrations in Go
to be..



## Creating Migrations
```
migrate create -ext sql -dir db/migrations create_posts_table
```

## pq package to access postgres database
```
go get github.com/lib/pq
```

## useful
```
go mod tidy
```