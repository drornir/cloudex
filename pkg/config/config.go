package config

// go install github.com/drornir/factor3@latest
//go:generate factor3 generate

//factor3:generate --filename config.yaml
type Config struct {
	//factor3:pflag port
	Port string
	//factor3:pflag log-level
	LogLevel string
	//factor3:pflag db
	SQLiteURL string
}
