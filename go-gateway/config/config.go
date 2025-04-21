package config

type Config struct {
	App []App `yaml:"app"`
}

type App struct {
	App struct {
		Port    string `yaml:"port"`
		Version string `yaml:"version"`
		Name    string `yaml:"name"`
	} `yaml:"app"`

	Http     HttpCfg   `yaml:"http"`
	Producer *Producer `yaml:"kafka"`
}
