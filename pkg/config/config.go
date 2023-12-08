package config

//go:generate factor3 generate

//factor3:generate --filename config.yaml
type Config struct {
	//factor3:pflag port
	Port string
}
