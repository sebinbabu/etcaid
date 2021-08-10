// Package etcaid implements a framework for configuration backup & restore of apps.
package etcaid

import (
	"path/filepath"
)

const (
	xdgConfigTargetDirName string = "xdg_config"
	homeTargetDirName      string = "home"
)

// ApplicationConfig is the raw configuration describing an application.
type ApplicationConfig struct {
	Name           string   // Name of the application
	HomePaths      []string // HomePaths is the list of source application file paths that exist relative to the user home directory
	XDGConfigPaths []string // XDGConfigPaths is the list of source application file paths that exist relative to the user XDG Config directory
}

// Application represents an application instance.
type Application struct {
	config        ApplicationConfig // holds the application configuration
	homePath      string            // user home directory path
	xdgConfigPath string            // xdg config directory path
	logger        logger            // holds the logger instance

	targetPath          string // target backup directory path
	targetHomePath      string // target backup directory path for application home config
	targetXDGConfigPath string // target backup directory path for application XDG config
}

// NewApp constructs an instance of Application using ApplicationConfig, paths & a concrete instance of logger.
func NewApplication(
	config ApplicationConfig,
	homePath string,
	xdgConfigPath string,
	targetPath string,
	logger logger,
) *Application {
	return &Application{
		config:        config,
		homePath:      homePath,
		xdgConfigPath: xdgConfigPath,
		logger:        logger,

		targetPath:          targetPath,
		targetHomePath:      filepath.Join(targetPath, homeTargetDirName),
		targetXDGConfigPath: filepath.Join(targetPath, xdgConfigTargetDirName),
	}
}

// Backup backups the application configuration relative to the target path.
// It assumes that your backup directory has been backed up externally (for example, git)
// and may overwrite files in case of conflict.
func (a *Application) Backup() {
	for _, p := range a.config.HomePaths {
		src := filepath.Join(a.homePath, p)
		dest := filepath.Join(a.targetHomePath, p)

		err := prepareAndCopy(src, dest, a.logger)
		if err != nil {
			a.logger.Error(err)
		}
	}

	for _, p := range a.config.XDGConfigPaths {
		src := filepath.Join(a.xdgConfigPath, p)
		dest := filepath.Join(a.targetXDGConfigPath, p)

		err := prepareAndCopy(src, dest, a.logger)
		if err != nil {
			a.logger.Error(err)
		}
	}
}

// Restore restores the application configuration relative to the target path.
// It assumes that your backup directory has been backed up externally (for example, git)
// and may overwrite files in case of conflict.
func (a *Application) Restore() {
	for _, p := range a.config.HomePaths {
		src := filepath.Join(a.targetHomePath, p)
		dest := filepath.Join(a.homePath, p)

		err := prepareAndCopy(src, dest, a.logger)
		if err != nil {
			a.logger.Error(err)
		}
	}

	for _, p := range a.config.XDGConfigPaths {
		src := filepath.Join(a.targetXDGConfigPath, p)
		dest := filepath.Join(a.xdgConfigPath, p)

		err := prepareAndCopy(src, dest, a.logger)
		if err != nil {
			a.logger.Error(err)
		}
	}
}

// ApplicationError records an error and the operation and that caused it.
type ApplicationError struct {
	Op      string
	Message string
	Path    string
	Err     error
}

func (e *ApplicationError) Error() string {
	s := e.Op + ": " + e.Message + ", "
	if e.Path != "" {
		s = s + e.Path + ", "
	}

	return s + e.Err.Error()
}

func (e *ApplicationError) Unwrap() error { return e.Err }
