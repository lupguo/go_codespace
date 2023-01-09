```bash
// export db dsn
export USER_DB_DNS="dbuser:dbpass@tcp(dbhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

// automigrate
go run entiy_test.go
```