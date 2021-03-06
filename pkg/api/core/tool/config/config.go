package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller Controller `json:"controller"`
	Node       Node       `json:"node"`
	DB         DB         `json:"db"`
	ImaCon     []ImaCon   `json:"imacon"`
	Mail       Mail       `json:"mail"`
	Radius     Radius     `json:"radius"`
	Slack      []Slack    `json:"slack"`
}

type Controller struct {
	TemplateConfPath string `json:"template_path"`
	User             User   `json:"user"`
	Admin            Admin  `json:"admin"`
	Auth             Auth   `json:"auth"`
	PublicKeyPath    string `json:"public_key_path"`
}

type Node struct {
	User string `json:"user"`
}

type ImaCon struct {
	IP   string `json:"ip"`
	Port uint   `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type User struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	URL      string `json:"url"`
	NoVNCURL string `json:"novnc_url"`
}

type Admin struct {
	IP        string    `json:"ip"`
	Port      int       `json:"port"`
	URL       string    `json:"url"`
	NoVNCURL  string    `json:"novnc_url"`
	AdminAuth AdminAuth `json:"auth"`
}

type AdminAuth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Auth struct {
	Token1 string `json:"token1"`
	Token2 string `json:"token2"`
	Token3 string `json:"token3"`
}

type DB struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
}

type Mail struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	CC       string `json:"cc"`
	Contract string `json:"contract"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}

type Radius struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Slack struct {
	WebHookUrl string `json:"url"`
	Channel    string `json:"channel"`
	Name       string `json:"name"`
}

var Conf Config

func GetConfig(inputConfPath string) error {
	configPath := "./data.json"
	if inputConfPath != "" {
		configPath = inputConfPath
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var data Config
	json.Unmarshal(file, &data)
	Conf = data
	return nil
}
