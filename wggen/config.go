package wggen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var cfgFileName string

func GetAllWGConfigFiles(dir string) ([]string, error) {
	if dir == "" {
		return nil, fmt.Errorf("dir cannot be empty")
	}
	files := make([]string, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileext := filepath.Ext(path)
		if !info.IsDir() && (fileext == ".yml" || fileext == ".yaml") {
			file := filepath.Base(path)
			files = append(files, file)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetWGConfig(dir, endpoint string) (*WGSrv, error) {

	cfgFiles, err := GetAllWGConfigFiles(dir)
	if err != nil {
		return nil, err
	}
	if len(cfgFiles) == 0 {
		return nil, fmt.Errorf("no config files found in %s", dir)
	}

	for _, cfgFile := range cfgFiles {
		if strings.Contains(cfgFile, endpoint) {
			cfgFileName = cfgFile
			break
		}
	}

	if cfgFileName == "" {
		return nil, fmt.Errorf("no config file found in %s with endpoint %s", dir, endpoint)
	}

	path, err := filepath.Abs(dir)
	cfgfile := filepath.Join(path, cfgFileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cfgfile)
	if err != nil {
		return nil, err
	}

	var config WGSrv
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (wg *WGSrv) SaveWGConfig(dir string) error {
	// Create the directory if it doesn't exist
	err := createDirectory(dir)
	if err != nil {
		return err
	}

	// Generate the filename based on the endpoint
	fullpath, _ := filepath.Abs(dir)
	filepath := filepath.Join(fullpath, wg.Endpoint)

	// Check if the file already exists
	if _, err := os.Stat(filepath); err == nil {
		return fmt.Errorf("%s does already exist", cfgFileName)
	}

	// Convert server struct to YAML
	data, err := yaml.Marshal(wg)
	if err != nil {
		return err
	}

	// Write the YAML data to the file
	err = os.WriteFile(filepath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (wg *WGSrv) UpdateWGConfig(dir string) error {

	// Generate the filename based on the endpoint

	fullpath, _ := filepath.Abs(dir)
	filepath := filepath.Join(fullpath, cfgFileName)

	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filepath)
		}
		return err
	}

	// Convert server struct to YAML
	data, err := yaml.Marshal(wg)
	if err != nil {
		return err
	}

	// Write the YAML data to the file
	err = os.WriteFile(filepath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func createDirectory(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}
