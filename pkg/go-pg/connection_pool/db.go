package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/kyl2016/Play-With-Golang/utility"
)

const (
	defaultDBConnMaxRetries         = 5
	defaultDBConnMaxLifetimeSeconds = 3600
	defaultDBMaxOpenConns           = 100
	defaultDBMinIdleConns           = 10
)

type DBService struct {
	*pg.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	r, err := q.FormattedQuery()
	fmt.Println(string(r))
	return err
}

type DatabaseConfig struct {
	URL string `yaml:"url"`

	ConnMaxLifetimeSeconds int `yaml:"conn_max_lifetime_sec"`
	MaxOpenConns           int `yaml:"max_open_conns"`
	MinIdleConns           int `yaml:"min_idle_conns"`
	MaxRetries             int `yaml:"max_retries"`
}

func InitDBService(config *DatabaseConfig) (DBService, error) {
	db, _, err := connectDatabase(config)
	db.AddQueryHook(dbLogger{})
	return DBService{db}, err
}

func IsDBNoRows(err error) bool {
	return errors.Is(err, pg.ErrNoRows)
}

func IsDBMultiRows(err error) bool {
	return errors.Is(err, pg.ErrMultiRows)
}

func connectDatabase(c *DatabaseConfig) (*pg.DB, string, error) {
	opt, err := pg.ParseURL(c.URL)
	if err != nil {
		return nil, "", err
	}
	opt.MaxConnAge = time.Second * time.Duration(c.ConnMaxLifetimeSeconds)
	opt.MinIdleConns = int(c.MinIdleConns)
	opt.PoolSize = int(c.MaxOpenConns)
	opt.MaxRetries = c.MaxRetries
	// opt.IdleCheckFrequency
	db := pg.Connect(opt)
	// var dbVersion string
	// _, err = db.QueryOne(pg.Scan(&dbVersion), "SELECT version()")
	// return db, dbVersion, err
	return db, "", err
}

func checkDBConfig(c *DatabaseConfig) error {
	if c.URL == "" {
		return errors.New("database address must be specified. (database.url)")
	}
	if c.ConnMaxLifetimeSeconds == 0 {
		c.ConnMaxLifetimeSeconds = defaultDBConnMaxLifetimeSeconds
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = defaultDBMaxOpenConns
	}
	if c.MinIdleConns == 0 {
		c.MinIdleConns = defaultDBMinIdleConns
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = defaultDBConnMaxRetries
	}
	return nil
}

// Wrap SELECT产生的错误对象，NoRows错误将被忽略
func WrapDBErrorOnLoad(err error) utility.Error {
	if err == nil || IsDBNoRows(err) {
		return nil
	}
	return utility.ErrDatabase.WrapWithAdditionalCallerSkip(err, 1)
}

// Wrap INSERT, UPDATE, DELETE产生的错误对象，按不同的错误使用相应的错误类型
func WrapDBErrorOnExec(err error) utility.Error {
	if err == nil {
		return nil
	}
	if IsDBNoRows(err) {
		return utility.ErrDatabaseNoChanges.WrapWithAdditionalCallerSkip(err, 1)
	}
	pgErr, ok := err.(pg.Error)
	if !ok {
		return utility.ErrDatabase.WrapWithAdditionalCallerSkip(err, 1)
	}
	if pgErr != nil && pgErr.IntegrityViolation() {
		return utility.ErrDatabaseConstraintViolation.WrapWithAdditionalCallerSkip(err, 1)
	}
	return utility.ErrDatabase.WrapWithAdditionalCallerSkip(err, 1)
}
