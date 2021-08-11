package etcaid

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Dir      string `toml:"etcaid_dir"`       // Name of the etcaid directory in user home, used for syncing backups
	LocalDir string `toml:"etcaid_local_dir"` // Name of the local etcaid directory in user home, for local backups
}

// ApplicationConfig is the configuration describing an application.
type ApplicationConfig struct {
	Title          string   `toml:"title"`            // Name of the application
	HomePaths      []string `toml:"home_paths"`       // HomePaths is the list of source application file paths that exist relative to the user home directory
	XDGConfigPaths []string `toml:"xdg_config_paths"` // XDGConfigPaths is the list of source application file paths that exist relative to the user XDG Config directory
}

// defaultConfig is the default configuration for etcaid.
var defaultConfig Config = Config{
	Dir:      "etcaid",
	LocalDir: ".etcaid_local",
}

// ParseApplicationConfig accepts a reader and parses it into ApplicationConfig.
// It returns an error if it fails.
func ParseApplicationConfig(r io.Reader) (ApplicationConfig, error) {
	config := ApplicationConfig{}
	decoder := toml.NewDecoder(r)

	_, err := decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// WriteApplicationConfig accepts a writer and writes the ApplicationConfig on it.
// It returns an error if it fails.
func WriteApplicationConfig(config ApplicationConfig, w io.Writer) error {
	encoder := toml.NewEncoder(w)

	err := encoder.Encode(&config)
	if err != nil {
		return err
	}

	return nil
}
