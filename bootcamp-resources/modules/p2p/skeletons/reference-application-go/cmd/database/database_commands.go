package database

const (
	DatabaseConnectionString = "postgresql://postgres:%s@%s:5432/postgres?sslmode=disable"
	CreateDatabaseQuery      = "SELECT 'CREATE DATABASE counter' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'counter')"
	CreateTableQuery         = "CREATE TABLE if not exists counter (name varchar(50) PRIMARY KEY, counter bigint NOT NULL);"
	SelectFromCounterTable   = "SELECT * FROM counter WHERE NAME='%s'"
)
