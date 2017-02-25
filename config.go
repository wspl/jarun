package main

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

const ConfigPath = "./jarun.toml"
var config *Config

type Config struct {
	toml *toml.TomlTree
}

func (c *Config) Set(key string, value string) {
	if c.toml == nil {
		c.toml, _ = toml.Load("")
	}
	c.toml.Set(key, value)
	ioutil.WriteFile(ConfigPath, []byte(c.toml.String()), 0644)
}

func (c *Config) String(key string) string {
	if c.toml == nil {
		t, err := toml.LoadFile(ConfigPath)
		if err != nil {
			c.toml, _ = toml.Load("")
		} else {
			c.toml = t
		}
	}
	if !c.toml.Has(key) {
		return ""
	}
	return c.toml.Get(key).(string)
}

func init() {
	config = new(Config)
}