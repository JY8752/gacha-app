package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	config struct {
		Mongo mongo
		exist bool //singltonにしたかったのでインスタンスがあるかどうかのフラグ
	}

	mongo struct {
		Uri string
	}
)

var c config

func GetConfig() config {
	if c.exist {
		fmt.Println("return config. already exist instance.")
		return c
	}

	m := os.Getenv("MODE")

	var f string
	switch m {
	case "production":
		f = "application.toml"
		if _, err := os.Stat(f); err != nil {
			f = "config/application.toml"
		}
	case "local":
		f = "application.local.toml"
		if _, err := os.Stat(f); err != nil {
			f = "config/application.local.toml"
		}
	case "test":
		f = "application.test.toml"
		if _, err := os.Stat(f); err != nil {
			f = "config/application.test.toml"
		}
	default:
		f = "application.local.toml"
		if _, err := os.Stat(f); err != nil {
			f = "config/application.local.toml"
		}
	}

	_, err := toml.DecodeFile(f, &c)
	if err != nil {
		log.Fatalf("fail decode toml file filePath: %s\n", f)
	}

	// mongoの秘匿情報セット
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_DOMAIN"))
	c.Mongo.Uri = uri

	c.exist = true
	return c
}
