package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	
}

func GetConfigurationPath (isFirstTime *bool) (string, error) {
	var globalConfigPath, err = os.UserConfigDir()
	var appConfigPath = filepath.Join(globalConfigPath, "yeel")
	
        if err != nil {
		fmt.Println("Unable to detect device config path!")
		fmt.Println("Defaulting to current...")
		appConfigPath = "."
	}

	var _, statErr = os.Stat(appConfigPath)
	if os.IsNotExist(statErr) {
		fmt.Printf("App config does not exist, creating at '%v'\n", appConfigPath)
		*isFirstTime = true
		if err := os.MkdirAll(appConfigPath, os.ModePerm); err != nil {
			fmt.Println("ERROR Unable to create config path!")
			return appConfigPath, err
		}
	}

	return appConfigPath, nil
}
