package main

import (
	"encoding/json"
	"errors"
)

const (
	DATABASE_UNDEFINED_DRIVER string = "undefined driver"
	DATABASE_INVALID_HOST     string = "invalid host"
	DATABASE_INVALID_PORT     string = "invalid port"
	DATABASE_INVALID_LOGIN    string = "invalid login"
	DATABASE_INVALID_NAME     string = "invalid database name"
	DATABASE_INVALID_PATH     string = "invalid path"

	EMPTY_SECRET   string = "empty secret"
	OUT_RANGE_PORT string = "out of range ports"

	DATABASE_MYSQL    DatabaseDriver = "mysql"
	DATABASE_POSTGRES DatabaseDriver = "postgres"
	DATABASE_MSSQL    DatabaseDriver = "mssql"
	DATABASE_SQLITE3  DatabaseDriver = "sqlite3"
)

type (
	DatabaseDriver string
	ConfigPort     int
	TCPService     struct {
		Secret string     `json:"secret"`
		Port   ConfigPort `json:"port"`
	}
	Database struct {
		Driver   DatabaseDriver `json:"driver"`
		Port     ConfigPort     `json:"port"`
		Host     string         `json:"host"`
		Login    string         `json:"login"`
		Password string         `json:"password"`
		NameDb   string         `json:"name_db"`
		Path     string         `json:"path"`
	}

	Node struct {
		TCPService
	}
	Rest struct {
		TCPService
	}

	Config struct {
		Database `json:"database"`
		Node     `json:"node"`
		Rest     `json:"rest"`
	}
)

func NewConfig(data []byte) (Config, error) {
	c := Config{}
	err := json.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	return c, c.validate()
}

/*
	Validate config file
*/
func (c Config) validate() error {
	if err := c.Database.validate(); err != nil {
		return err
	}
	if err := c.Node.validate(); err != nil {
		return err
	}
	if err := c.Rest.validate(); err != nil {
		return err
	}
	return nil
}

/*
	Validate database object
*/
func (d Database) validate() error {

	err := d.Driver.validate()
	if err != nil {
		return err
	}
	err = d.Port.validate()
	if err != nil {
		return err
	}

	switch d.Driver {
	case DATABASE_MYSQL, DATABASE_POSTGRES, DATABASE_MSSQL:
		if d.Host == "" {
			return errors.New(DATABASE_INVALID_HOST)
		}
		if d.Port == 0 {
			return errors.New(DATABASE_INVALID_PORT)
		}
		if d.Login == "" {
			return errors.New(DATABASE_INVALID_LOGIN)
		}
		if d.NameDb == "" {
			return errors.New(DATABASE_INVALID_NAME)
		}
	case DATABASE_SQLITE3:
		if d.Path == "" {
			return errors.New(DATABASE_INVALID_PATH)
		}
	default:
		return errors.New(DATABASE_UNDEFINED_DRIVER)
	}
	return nil
}

func (t TCPService) validate() error {
	if t.Port != 0 && t.Secret == "" {
		return errors.New(EMPTY_SECRET)
	}
	return nil
}

func (d DatabaseDriver) validate() error {
	if d != DATABASE_MYSQL &&
		d != DATABASE_POSTGRES &&
		d != DATABASE_MSSQL &&
		d != DATABASE_SQLITE3 {
		return errors.New(DATABASE_UNDEFINED_DRIVER)
	}
	return nil
}

func (p ConfigPort) validate() error {
	if p <= 0 || p > 65535 {
		return errors.New(OUT_RANGE_PORT)
	}
	return nil
}
