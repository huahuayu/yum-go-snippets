package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

func init() {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	configPath := pwd + "/config/config.toml"

	var conf Config
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("title: %s,owner name:%s, database server:%s, clients hosts: %v\n ", conf.Title, conf.Owner.Name, conf.DB.Server, conf.Clients.Hosts[1])
	}
}
func main() {
}
