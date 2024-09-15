# backend

### PURE net/http YEAH

### Toos used 
1. `go-swag` for generating swagger file 
2. `go-migrate` for generating migration file to database

### Running the application 
```
swag init && go run . 
```

### Create migration 

```
migrate create -ext sql -dir ${MIGRATION_PATH} -seq [migration_name]
```

### Up & Down the Migration 

```
migrate -database ${POSTGRES_URL} -path ${MIGRATION_PATH} up

migrate -database ${POSTGRES_URL} -path ${MIGRATION_PATH} down

```

### Useful Link 

Excalidraw Skech : [here](https://excalidraw.com/#room=e2412169cd113b60fdfb,UOX13qftyo7Oj6IEZ4X-qA)