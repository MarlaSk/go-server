package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port string `json:"port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (c Config) GetDatabaseDSN() string {
	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8",
		c.Database.Host,
		c.Database.Port,
		c.Database.Username,
		c.Database.Password,
		c.Database.Name,
	)
	return s
}

func(c Config) GetServerPort() string {
	return ":" + c.Port
}

func Load(path string) (Config, error) {
 file,err :=	os.Open(path)

 if err != nil {
	return Config{}, err
 }

 defer file.Close()

 var cfg Config
 
 json.NewDecoder(file).Decode(&cfg)

 return cfg,nil
}
