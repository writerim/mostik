package main

import (
	"fmt"
	"testing"
)

/*
	{
		database : {
			driver : mysql | sqlite3 | etc
			host : 127.0.0.1
			port : 3306
			login : root
			password : 123
			name_db : test
			path : ./database.db
		},
		node : {
			port : 8080
			secret : secret_key
		}
		rest : {
			port : 3030
			secret : secret_key
		}
	}
*/
func TestConfigValidate(t *testing.T) {

	valid_configs_databases := []Config{
		Config{
			Database: Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     3306,
				Login:    "root",
				Password: "123",
				NameDb:   "test",
			},
		},
		Config{
			Database: Database{
				Driver:   "postgres",
				Host:     "127.0.0.1",
				Port:     3306,
				Login:    "root",
				Password: "123",
				NameDb:   "test",
			},
		},
		Config{
			Database: Database{
				Driver:   "mssql",
				Host:     "127.0.0.1",
				Port:     3306,
				Login:    "root",
				Password: "123",
				NameDb:   "test",
			},
		},
		Config{
			Database: Database{
				Driver: "sqlite3",
				Path:   "./home",
			},
		},
	}

	for _, item := range valid_configs_databases {
		if item.validate() != nil {
			t.Fatal(fmt.Sprintf("%+v invalid database config", item))
		}
	}

}
