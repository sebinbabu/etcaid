package etcaid

import (
	"os"
	"path/filepath"
)

const (
	configFilename string = ".etcaidrc.toml" // Name of the file inside user home that stores etcaid configuration
	backupDir      string = "backup"         // Name of the dir that holds backup files
	applicationDir string = "applications"   // Name of the dir that holds application config files
)

// Controller represents the etcaid controller.
// External user interfaces such as CLI will invoke the controller for working with etcaid.
type Controller struct {
	mainDir        string // Holds path to etcaid directory in user home, used for syncing backups
	localDir       string // Holds path to local etcaid directory in user home, for local backups
	applicationDir string // Holds path to application config directory in mainDir
	mainBackupDir  string // Holds path to application backup directory in mainDir
	localBackupDir string // Holds path to local backup directory in localDir

	homePath      string                  // user home directory path
	xdgConfigPath string                  // xdg config directory path
	logger        logger                  // holds the logger instance
	applications  map[string]*Application // maps the unique application name with an instance of Application
}

// NewController constructs an instance of Controller using etcaid config path,
// user home path, xdg directory path ahd an instance of logger.
func NewController(homePath string, xdgConfigPath string, logger logger) *Controller {
	config := defaultConfig

	mainDir := filepath.Join(homePath, config.Dir)
	mainBackupDir := filepath.Join(mainDir, backupDir)
	applicationDir := filepath.Join(mainDir, applicationDir)

	localDir := filepath.Join(homePath, config.LocalDir)
	localBackupDir := filepath.Join(localDir, backupDir)

	return &Controller{
		mainDir:        mainDir,
		localDir:       localDir,
		applicationDir: applicationDir,
		mainBackupDir:  mainBackupDir,
		localBackupDir: localBackupDir,

		homePath:      homePath,
		xdgConfigPath: xdgConfigPath,
		logger:        logger,
	}
}

// Init initializes etcaid directories & configuration.
// It returns an error if it fails.
func (c *Controller) Init() error {
	dirs := []string{c.applicationDir, c.mainBackupDir, c.localBackupDir}

	for _, d := range dirs {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return &ApplicationError{
				Op:      "init",
				Message: "failed to create directory",
				Path:    d,
				Err:     err,
			}
		}
	}

	return nil
}

// LoadApplications loads the application configurations from the application config files
// and makes the applications ready for further use.
func (c *Controller) LoadApplications() error {
	c.applications = make(map[string]*Application)
	configDir := c.applicationDir

	files, err := os.ReadDir(configDir)
	if err != nil {
		if files == nil {
			return &ApplicationError{
				Op:      "LoadApplications",
				Message: "failed to read directory",
				Path:    configDir,
				Err:     err,
			}
		}

		c.logger.Error(&ApplicationError{
			Op:      "LoadApplications",
			Message: "failed to read some entries in directory",
			Path:    configDir,
			Err:     err,
		})
	}

	for _, f := range files {
		name, ext := parseFilename(f.Name())
		if !f.Type().IsRegular() || ext != ".toml" || !isValidAppName(name) {
			continue
		}

		confPath := filepath.Join(configDir, f.Name())
		file, err := os.Open(confPath)
		defer file.Close()
		if err != nil {
			c.logger.Error(&ApplicationError{
				Op:      "LoadApplications",
				Message: "failed to open configuration for reading",
				Path:    confPath,
				Err:     err,
			})

			continue
		}

		config, err := ParseApplicationConfig(file)
		if err != nil {
			c.logger.Error(&ApplicationError{
				Op:      "LoadApplications",
				Message: "failed to parse application configuration",
				Path:    confPath,
				Err:     err,
			})

			continue
		}

		application := NewApplication(
			config,
			c.homePath,
			c.xdgConfigPath,
			filepath.Join(c.mainBackupDir, name),
			c.logger,
		)
		c.applications[name] = application
	}

	return nil
}

// BackupAll backs up all applications from their locations
// as defined by the application configuration.
// Backed up applications are available in the etcaid dir.
func (c *Controller) BackupAll() {
	for _, a := range c.applications {
		a.Backup()
	}
}

// RestoreAll restores all applications from the etcaid dir
// to their locations as defined by the application configuration.
func (c *Controller) RestoreAll() {
	for _, a := range c.applications {
		a.Restore()
	}
}

// Create creates a new application configuration file to be used for backups.
// It returns the path to this file, or an error if it fails.
func (c *Controller) Create(name string) (string, error) {
	if !isValidAppName(name) {
		return "", &ApplicationError{
			Op:      "Create",
			Message: "failed to create application, names can only contain letters and numbers",
		}
	}

	confPath := filepath.Join(c.applicationDir, name+".toml")

	if _, err := os.Stat(confPath); err == nil || !os.IsNotExist(err) {
		return "", &ApplicationError{
			Op:      "Create",
			Message: "another application with the same name already exists",
			Err:     err,
		}
	}

	file, err := os.Create(confPath)
	if err != nil {
		return "", &ApplicationError{
			Op:      "Create",
			Message: "failed to create application configuration",
			Err:     err,
		}
	}
	defer file.Close()

	sampleConf := ApplicationConfig{
		Title:          name,
		HomePaths:      []string{},
		XDGConfigPaths: []string{},
	}

	err = WriteApplicationConfig(sampleConf, file)
	if err != nil {
		return "", &ApplicationError{
			Op:      "Create",
			Message: "failed to write configuration",
			Err:     err,
		}
	}

	return confPath, nil
}

// ApplicationConfigPath returns the path to application configuration file
// for an application with the given name. It returns the an error if it fails.
func (c *Controller) ApplicationConfigPath(name string) (string, error) {
	if !isValidAppName(name) {
		return "", &ApplicationError{
			Op:      "ApplicationConfigPath",
			Message: "failed to get config, names can only contain letters and numbers",
		}
	}

	confPath := filepath.Join(c.applicationDir, name+".toml")
	_, err := os.Stat(confPath)

	if err != nil {
		if os.IsNotExist(err) {
			return "", &ApplicationError{
				Op:      "ApplicationConfigPath",
				Message: "application " + name + " doesn't exist",
				Path:    confPath,
			}
		}

		return "", &ApplicationError{
			Op:      "ApplicationConfigPath",
			Message: "failed to find application config",
			Path:    confPath,
			Err:     err,
		}
	}

	return confPath, nil
}
