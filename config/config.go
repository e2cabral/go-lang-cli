package config

import (
	"flag"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

// Config - struct to contain the parameters
type Config struct {
	Package  string
	Path     string
	Filename string
}

// NewConfig - returns a new config struct
func NewConfig() *Config {
	return &Config{}
}

// Setup - sets all flag variables and their shorthand
func (c *Config) Setup() {
	flag.StringVar(&c.Filename, "filename", "", "Is the name of the file to be generated.")
	flag.StringVar(&c.Filename, "f", "", "Is the name of the file to be generated.")

	flag.StringVar(&c.Path, "path", "", "Determines the path where the archive will be generated.")

	flag.StringVar(&c.Package, "package", "main", "The name of the package.")
	flag.StringVar(&c.Package, "p", "main", "The name of the package.")
}

// SetPackage - sets the package name in the archive
func (c *Config) SetPackage(file *os.File) (*os.File, error) {
	t := `package %v

// %v - Your description comment
func %v () {
	
}
`
	text := []byte(fmt.Sprintf(t, c.Package, strcase.ToCamel(c.Filename), strcase.ToCamel(c.Filename)))
	if _, err := file.Write(text); err != nil {
		return nil, err
	}

	return file, nil
}

// ExecuteCommand - creates the file with the specified parameters
func (c *Config) ExecuteCommand() error {
	if err := os.Mkdir(c.Package, 0777); err != nil {
		return err
	}

	if c.Package == "main" {
		_, err := os.Create(c.Filename + ".go")
		if err != nil {
			return err
		}
		return nil
	}

	f, err := os.Create(c.Path + "/" + c.Filename + ".go")
	if err != nil {
		return err
	}

	if _, err := c.SetPackage(f); err != nil {
		return err
	}

	return nil
}
