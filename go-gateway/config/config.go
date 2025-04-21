package config

type Config struct {
	App []App `yaml:"app"`
}

type App struct {
	App struct {
		Port    string `yaml:"port"`
		Version string `yaml:"version"`
	} `yaml:"app"`

	Http HttpCfg `yaml:"http"`
}
