package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	PG_HOST          = "PG_HOST"
	PG_USER          = "PG_USER"
	PG_DB_NAME       = "PG_DB_NAME"
	PG_PASSWORD      = "PG_PASSWORD"
	PG_PORT          = "PG_PORT"
	PG_POOL_SIZE     = "PG_POOL_SIZE"
	PG_POOL_MAX_LIFE = "PG_POOL_MAX_LIFE"
)

type PostgresConfig struct {
	Host        string
	Port        uint32
	User        string
	Password    string
	DBName      string
	PoolSize    uint8
	PoolMaxLife time.Duration
}

func NewPostgresConn(cfg *PostgresConfig) *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(int(cfg.PoolSize))
	db.SetMaxOpenConns(int(cfg.PoolSize))
	db.SetConnMaxLifetime(cfg.PoolMaxLife)

	return db
}

func NewPostgresConfig() *PostgresConfig {

	host := os.Getenv(PG_HOST)
	if host == "" {
		panic(fmt.Sprintf("%s env not found", PG_HOST))
	}

	port := os.Getenv(PG_PORT)
	if port == "" {
		panic(fmt.Sprintf("%s env not found", PG_PORT))
	}

	portVal, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("%s %s", PG_PORT, err.Error()))
	}

	user := os.Getenv(PG_USER)
	if user == "" {
		panic(fmt.Sprintf("%s env not found", PG_USER))
	}

	password := os.Getenv(PG_PASSWORD)
	if password == "" {
		panic(fmt.Sprintf("%s env not found", PG_PASSWORD))
	}

	dbname := os.Getenv(PG_DB_NAME)
	if dbname == "" {
		panic(fmt.Sprintf("%s env not found", PG_DB_NAME))
	}

	poolSize := os.Getenv(PG_POOL_SIZE)
	if poolSize == "" {
		poolSize = "10"
	}

	poolSizeVal, err := strconv.Atoi(poolSize)
	if err != nil {
		panic(fmt.Sprintf("%s %s", PG_POOL_SIZE, err.Error()))
	}

	poolMaxLife := os.Getenv(PG_POOL_MAX_LIFE)
	if poolMaxLife == "" {
		poolMaxLife = "60s"
	}

	poolMaxLifeVal, err := time.ParseDuration(poolMaxLife)
	if err != nil {
		panic(fmt.Sprintf("%s %s", PG_POOL_MAX_LIFE, err.Error()))
	}

	return &PostgresConfig{
		Host:        host,
		User:        user,
		Port:        uint32(portVal),
		Password:    password,
		DBName:      dbname,
		PoolSize:    uint8(poolSizeVal),
		PoolMaxLife: poolMaxLifeVal,
	}
}
