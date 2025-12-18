package dbo

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ENGINE string
type DRIVER string

const (
	DRIVER_MYSQL  DRIVER = "mysql"
	DRIVER_PGSQL  DRIVER = "pgsql"
	DRIVER_SQLITE DRIVER = "sqlite"
)

const (
	ENGINE_INNODB ENGINE = "InnoDB"
)

type Options struct {
	Driver          DRIVER
	Host            string
	Port            string
	Username        string
	Password        string
	DBName          string
	Charset         string
	Collation       string
	DSN             string
	Engine          ENGINE
	Config          *gorm.Config
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // Maximum lifetime for a connection (in milliseconds)
	ConnMaxIdleTime int // Maximum idle time for a connection (in milliseconds)
}

func (x *Options) getConfig() *gorm.Config {
	if x.Config != nil {
		return x.Config
	}

	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
}

func (x *Options) getDialector() (conn gorm.Dialector, err error) {
	switch x.getDriver() {
	case DRIVER_MYSQL:
		conn = func() gorm.Dialector {
			if x.DSN != "" {
				return mysql.Open(x.DSN)
			}

			return mysql.Open(fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
				x.Username,
				x.Password,
				x.Host,
				x.Port,
				x.DBName,
				x.getCharset(),
				x.getCollation(),
			))
		}()

	case DRIVER_PGSQL:
		conn = func() gorm.Dialector {
			if x.DSN != "" {
				return postgres.Open(x.DSN)
			}

			return postgres.Open(fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s",
				x.Username,
				x.Password,
				x.Host,
				x.Port,
				x.DBName,
			))
		}()

	case DRIVER_SQLITE:
		conn = sqlite.Open(x.DSN)

	default:
		err = errors.New("DB Driver not available")
	}

	return
}

func (x *Options) getDriver() DRIVER {
	if x.Driver == "" {
		x.Driver = DRIVER_PGSQL
	}

	return x.Driver
}

func (x *Options) getEngine() ENGINE {
	if x.Engine == "" {
		x.Engine = ENGINE_INNODB
	}

	return x.Engine
}

func (x *Options) getCharset() string {
	if x.Charset == "" {
		x.Charset = "utf8mb4"
	}

	return x.Charset
}

func (x *Options) getCollation() string {
	if x.Collation == "" {
		x.Collation = "utf8mb4_general_ci"
	}

	return x.Collation
}
