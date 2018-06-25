package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	// external
	"github.com/cep21/xdgbasedir"
	"gopkg.in/yaml.v2"
)

var (
	ProgramName         string = "impi"
	configDirectoryPath string = "./conf"
)

// ServiceConfig contains configuration information for a service
type ServiceConfig struct {
	Token string
	User  string
}

// OutputConfig sontains configuration information for an output
type OutputConfig struct {
	SpinnerIndex    int    `yaml:"spinnerIndex"`
	SpinnerInterval int    `yaml:"spinnerInterval"`
	SpinnerColor    string `yaml:"spinnerColor"`
}

// Config contains configuration information
type Config struct {
	// Services     map[string]*ServiceConfig `yaml:"services"`
	GoRoot  string                   `yaml:"goRoot"`
	GoSrc   string                   `yaml:"goSrc"`
	GoPath  string                   `yaml:"goPath"`
	Outputs map[string]*OutputConfig `yaml:"outputs"`
}

/*
// GetService returns the configuration information for a service
func (config *Config) GetService(name string) *ServiceConfig {
	if config.Services == nil {
		config.Services = make(map[string]*ServiceConfig)
	}

	service := config.Services[name]
	if service == nil {
		service = &ServiceConfig{}
		config.Services[name] = service
	}
	return service
}
*/

// GetOutput returns the configuration information for an output
func (c *Config) GetOutput(name string) *OutputConfig {
	if c.Outputs == nil {
		c.Outputs = make(map[string]*OutputConfig)
	}

	output := c.Outputs[name]
	if output == nil {
		output = &OutputConfig{}
		c.Outputs[name] = output
	}
	return output
}

// ReadConfig reads the configuration information
func ReadConfig() (*Config, error) {
	file := configFilePath()

	var c Config
	if _, err := os.Stat(file); err == nil {
		// Read and unmarshal file only if it exists
		f, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// WriteConfig writes the configuration information
func (c *Config) WriteConfig() error {
	err := os.MkdirAll(configDirectoryPath, 0700)
	if err != nil {
		return err
	}
	fmt.Println("configFilePath=", configFilePath())
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath(), data, 0600)
}

func configFilePath() string {
	return path.Join(configDirectoryPath, fmt.Sprintf("%s.yaml", ProgramName))
}

func init() {
	baseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		log.Fatal("Can't find XDG BaseDirectory")
	} else {
		configDirectoryPath = path.Join(baseDir, ProgramName)
	}
}
