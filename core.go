package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	// "interfaces/restful"
	"interfaces/rpcnode"
	"io/ioutil"
)

func main() {

	// example usage
	// ./core --config path_to_config.json
	path_to_config := flag.String("path_config", "", "set path to config file")
	flag.Parse()

	// parse config
	file, err := ioutil.ReadFile(*path_to_config)
	if err != nil {
		logrus.Warn(err.Error())
	}

	_, err_config := NewConfig([]byte(file))
	if err_config != nil {
		logrus.Warn(err.Error())
	}

	rpcnode.Init(9999)
	// restful.New(9999)
}
