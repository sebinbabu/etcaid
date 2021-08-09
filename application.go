// Package etcaid implements a framework for configuration backup & restore of apps.
package etcaid

// ApplicationConfig is the raw configuration describing an application.
type ApplicationConfig struct {
	Name           string   // Name of the application
	HomeFiles      []string // HomeFiles is the list of application files that exist relative to the user home directory
	XDGConfigFiles []string // XDG Config is the list of application files that exist relative to the user XDG Config directory
}

// Application represents an application instance.
type Application struct {
	config        ApplicationConfig // holds the application configuration
	targetPath    string            // target backup directory path
	homePath      string            // user home directory path
	xdgConfigPath string            // xdg config directory path
	logger        logger            // holds the logger instance
}

// NewApp constructs an instance of Application using ApplicationConfig, paths & a concrete instance of logger.
func (a *Application) NewApplication(config ApplicationConfig, homePath string, xdgConfigPath string, logger logger) *Application {
	return &Application{
		config:        config,
		homePath:      homePath,
		xdgConfigPath: xdgConfigPath,
		logger:        logger,
	}
}
