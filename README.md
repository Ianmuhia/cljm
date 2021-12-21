
# CLJ Maranatha church monolith.
Uses:
- Go version go1.17.5
- The [Gin framework](https://github.com/gin-gonic/gin).
- The [pgx](https://github.com/jackc/pgx/v4) PostgreSQL Driver and Toolkit.
- The [jwt-go](https://github.com/dgrijalva/jwt-go) for JSON-web-tokens.
- The [SCS](github.com/alexedwards/scs/v2) HTTP Session Management for Go.



Handling Issues:
- Docker postgres doesn't start RUN: `sudo ss -lptn 'sport = :5432' & sudo kill PID`.
- `docker exec -it [container-id] bash` to shell into specific container.
- `migrate create -ext sql -dir db/migrate -seq init_schema` 
- `migrate -path db/migration -database "db-url" -verbose up` 
- Work around on dirt database delete schema before migrations.