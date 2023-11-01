package config

type Service struct {
	Addr    string `yaml:"Addr"`
	Name    string `yaml:"Name"`
	Version string `yaml:"Version"`
}

type Mail struct {
	Name  string `yaml:"Name"`
	Email string `yaml:"Email"`
}

type Smtp struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type JWT struct {
	PrivateKeyPath string `yaml:"PrivateKeyPath"`
	PublicKeyPath  string `yaml:"PublicKeyPath"`
}

type System struct {
	Theme        string `yaml:"Theme"`
	Title        string `yaml:"Title"`
	Keywords     string `yaml:"Keywords"`
	Description  string `yaml:"Description"`
	RecordNumber string `yaml:"RecordNumber"`
}

type Event struct {
	DelCache string `yaml:"DelCache"`
}
