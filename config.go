package main

import (
	"errors"
)

type DatabaseError string
type DatabaseEnum string

func (d_err DatabaseError) toError() error {
	return errors.New(string(d_err))
}
func (d_err DatabaseError) String() string {
	return string(d_err)
}

func (d_err DatabaseEnum) toError() error {
	return errors.New(string(d_err))
}
func (d_err DatabaseEnum) String() string {
	return string(d_err)
}

const (
	DATABASE_UNDEFINED_DRIVER DatabaseError = "undefined driver"
	DATABASE_INVALID_HOST     DatabaseError = "invalid host"
	DATABASE_INVALID_PORT     DatabaseError = "invalid port"
	DATABASE_INVALID_LOGIN    DatabaseError = "invalid login"
	DATABASE_INVALID_NAME     DatabaseError = "invalid database name"
	DATABASE_INVALID_PATH     DatabaseError = "invalid path"

	DATABASE_MYSQL    DatabaseEnum = "mysql"
	DATABASE_POSTGRES DatabaseEnum = "postgres"
	DATABASE_MSSQL    DatabaseEnum = "mssql"
	DATABASE_SQLITE3  DatabaseEnum = "sqlite3"
)

type (
	Database struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Login    string `json:"login"`
		Password string `json:"password"`
		NameDb   string `json:"name_db"`
		Path     string `json:"path"`
	}
	Node struct {
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	}
	Rest struct {
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	}

	Config struct {
		Database `json:"database"`
		Node     `json:"node"`
		Rest     `json:"rest"`
	}
)

/*
	Validate config file
*/
func (c Config) validate() error {
	if err := c.Database.validate(); err != nil {
		return err
	}

	// TODO other validate

	return nil
}

/*
	Validate database object
*/
func (d Database) validate() error {
	switch d.Driver {
	case DATABASE_MYSQL.String(), DATABASE_POSTGRES.String(), DATABASE_MSSQL.String():
		if d.Host == "" {
			return DATABASE_INVALID_HOST.toError()
		}
		if d.Port == 0 {
			return DATABASE_INVALID_PORT.toError()
		}
		if d.Login == "" {
			return DATABASE_INVALID_LOGIN.toError()
		}
		if d.NameDb == "" {
			return DATABASE_INVALID_NAME.toError()
		}
	case DATABASE_SQLITE3.String():
		if d.Path == "" {
			return DATABASE_INVALID_PATH.toError()
		}
	default:
		return DATABASE_UNDEFINED_DRIVER.toError()
	}
	return nil
}
