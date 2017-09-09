package config

import (
	"io/ioutil"
	"encoding/json"
)

type config struct {
	DBName           string `json:"db_name"`
	DBSource         string `json:"db_source"`
}

var c config
func init() {
	c.DBName = "mysql"
	c.DBSource = "root:xys147852@/labgo?charset=utf8&parseTime=True&loc=Local"
}

func LoadConfig(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var ctmp config
	err = json.Unmarshal(b, &ctmp)
	if err != nil {
		return err
	}

	c = ctmp
	return nil
}

func GetDBName() string {
	return c.DBName
}

func GetDBSource() string {
	return c.DBSource
}



