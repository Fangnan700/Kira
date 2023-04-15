package model

type Config struct {
	Domain    string   `yaml:"DOMAIN"`
	Host      string   `yaml:"HOST"`
	Port      int      `yaml:"PORT"`
	Character string   `yaml:"CHARACTER"`
	ApiKeys   []string `yaml:"APIKEY"`
}
