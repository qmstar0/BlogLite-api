package config

type Database struct {
	//Host    string `yaml:"Host"`
	//Port    int    `yaml:"Port"`
	Addr    string `yaml:"Addr"`
	User    string `yaml:"User"`
	Name    string `yaml:"Name"`
	Charset string `yaml:"Charset"`
}

type Redis struct {
	//Host string `yaml:"Host"`
	//Port string `yaml:"Port"`
	Addr string `yaml:"Addr"`
}
