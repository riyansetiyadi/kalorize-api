package config

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Host                   string `mapstructure:"host"`
	Port                   string `mapstructure:"port"`
	DBName                 string `mapstructure:"dbname"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	InstanceConnectionName string `mapstructure:"instance_connection_name"`
	UsePrivate             bool   `mapstructure:"use_private"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

func InitDB() (*sql.DB, error) {
	viper.SetConfigFile("prod.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling config: %w", err)
	}

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	var opts []cloudsqlconn.DialOption
	if config.Database.UsePrivate {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}

	mysql.RegisterDialContext("cloudsqlconn", func(ctx context.Context, addr string) (net.Conn, error) {
		return d.Dial(ctx, config.Database.InstanceConnectionName, opts...)
	})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:%s)/%s?parseTime=true",
		config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DBName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return db, nil
}
