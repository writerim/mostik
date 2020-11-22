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

	strings_conf_node := map[string]string{
		`{"database":{"driver":"mysql","host":"127.0.0.1","port":3306,"name_db":"sd","password":"we","login":"df"}}`: ``,
	}

	format_error := func(err, err_valid string) string {
		return fmt.Sprintf("%s != %s", err, err_valid)
	}

	for item, err_res := range strings_conf_node {
		_, err := NewConfig([]byte(item)) // TODO
		if err != nil && err.Error() != err_res {
			t.Fatal(format_error(err.Error(), err_res))
		} else if err == nil && err_res != "" {
			t.Fatal(format_error(``, err_res))
		}
		// other good
	}

}
