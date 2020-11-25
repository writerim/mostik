package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"interfaces/restful"
	"interfaces/rpcnode"
	"io/ioutil"
	"repositories"
)

var sessions *cache.Cache

func init() {
	sessions = cache.New(cache.NoExpiration, cache.NoExpiration)
}

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
		Debug    bool           `json:"debug"`
	}

	Session struct {
		Secr
	}

	Node struct {
		TCPService
	}
	Rest struct {
		Port        ConfigPort `json:"port"`
		SessionLive int        `json:"session_live"` // Время жизни сессии
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

func (d Database) stringConnect() string {
	switch d.Driver {
	case DATABASE_MYSQL:
		return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC`,
			d.Login,
			d.Password,
			d.Host,
			d.Port,
			d.NameDb)
	case DATABASE_MSSQL:
		return fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`,
			d.Login,
			d.Password,
			d.Host,
			d.Port,
			d.NameDb)
	case DATABASE_SQLITE3:
		return d.Path
	case DATABASE_POSTGRES:
		return fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s`,
			d.Host,
			d.Port,
			d.Login,
			d.NameDb,
			d.Password)
	}
	return ""
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

func main() {

	// example usage
	// ./core --config path_to_config.json
	path_to_config := flag.String("path_config", "", "set path to config file")
	flag.Parse()

	// parse config
	file, err := ioutil.ReadFile(*path_to_config)
	if err != nil {
		logrus.Fatalf("Error: Read config file: %s", err.Error())
	}

	logrus.Info("Read config")
	config, err_config := NewConfig([]byte(file))
	if err_config != nil {
		logrus.Fatal(err_config.Error())
	}
	logrus.Info("Read config successfull")

	db, err := gorm.Open(config.Database.Driver, config.Database.stringConnect())
	if err != nil {
		logrus.Fatal(err.Error())
	}
	db.LogMode(config.Database.Debug)
	repositories.Init(db)

	logrus.Info("Start rpc service")
	go rpcnode.Init(int(config.Node.Port))
	logrus.Info("Start rpc service successfull")

	logrus.Info("Start rest service")

	restful.AddSession = add_to_session
	// restful.GetSession = get_session
	// restful.DelSession = del_session

	go restful.Init(int(config.Rest.Port))
	logrus.Info("Start rest service successfull")

	select {}
}

/*
	Добавление в сессию пользователя
*/
func add_to_session(user_id int) (string, error) {

	if user_id == 0 {
		return "", errors.New("user id == 0. add to session not access")
	}

	out_str_g, _ := uuid.NewUUID()
	out_str := out_str_g.String()

	live := cache.NoExpiration

	for ident, _ := range sessions.Items() {
		id_user, _ := sessions.Get(ident)
		if id_user.(int) == user_id {
			sessions.Set(out_str, user_id, live)
			return out_str, nil
		}
	}

	if config.SessionLive != 0 {
		live = time.Duration(config.SessionLive) * time.Second
	}

	sessions.Add(out_str, user_id, live)

	return out_str, nil
}
