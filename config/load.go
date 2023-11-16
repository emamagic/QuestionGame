package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"
)

func Load(configPath string) Config {
	// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	var k = koanf.New(".")

	// Load default values using the confmap provider.
	// We provide a flat map with the "." delimiter.
	// A nested map can be loaded by setting the delimiter to an empty string "".
	k.Load(confmap.Provider(defaultConfig, "."), nil)

	// Load YAML config and merge into the previously loaded config (because we can).
	k.Load(file.Provider(configPath), yaml.Parser())

	k.Load(env.Provider("QuestionAPP_", ".", func(s string) string {
		str := strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "QuestionAPP_")), "_", ".", -1)

		// TODO - this is just a workaround
		// for multiword items such as "sign_key" that we should use like "QuestionAPP_AUTH_SIGN__KEY"
		// find a better solution if needed..
		return strings.Replace(str, "..", "_", -1)
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg
}