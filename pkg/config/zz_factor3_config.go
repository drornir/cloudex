// GENERATED FILE DO NOT EDIT
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var _ = strconv.ParseInt // just in case strconv is not used

func (self *Config) Factor3Load(argv []string) error {
	conf := zz_factor3_Config{}
	conf.Filename = "config.yaml"

	type jsonStruct struct {
		Port string `json:"port"`
	}

	loadConfigFile := func(filename string) error {
		type jsonStruct struct {
			Port string `json:"port"`
		}
		var jsoner json.Unmarshaler
		if x, ok := interface{}(self).(json.Unmarshaler); ok {
			jsoner = x
		} else {
			jsoner = &zz_factor3_JSONer[jsonStruct]{t: (*jsonStruct)(self)}
		}
		file, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}

		fileExt := filename[strings.LastIndex(filename, ".")+1:]
		switch fileExt {
		case "yaml", "yml":
			intoMap := make(map[string]interface{})
			err = yaml.Unmarshal(file, intoMap)
			if err != nil {
				break
			}
			intoJSON, e := json.Marshal(intoMap)
			if e != nil {
				err = e
				break
			}
			err = json.Unmarshal(intoJSON, jsoner)
		case "json":
			err = json.Unmarshal(file, jsoner)
		default:
			return fmt.Errorf("unsupported file type %q", fileExt)
		}
		if err != nil {
			return fmt.Errorf("unmarshaling: %w", err)
		}

		return nil
	}

	loadEnv := func(prefix string) error {
		var s string
		s = os.Getenv(prefix + "PORT")
		if s != "" {
			self.Port = s
		}
		return nil
	}

	parseFlags := func(argv []string) error {
		if len(argv) == 0 {
			return nil
		}
		fset := pflag.NewFlagSet("Config", pflag.ContinueOnError)
		fset.StringVarP(&self.Port, "port", "", self.Port, "")

		if err := fset.Parse(argv); err != nil {
			return fmt.Errorf("parsing flags: %w", err)
		}
		return nil
	}

	if err := loadConfigFile(conf.Filename); err != nil {
		return fmt.Errorf("loading config from file %q: %w", conf.Filename, err)
	}
	if err := loadEnv(conf.EnvPrefix); err != nil {
		return fmt.Errorf("loading config from env: %w", err)
	}
	if err := parseFlags(argv); err != nil {
		return fmt.Errorf("loading config from pflags: %w", err)
	}

	return nil
}
