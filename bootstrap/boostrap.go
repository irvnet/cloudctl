package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"
)

// check for local bootstrap script and create if not exist
func EnsureBootstrapScriptExists(machineName string) error {

	scriptName := fmt.Sprintf("bootstrap.%s.sh", machineName)
	scriptPath := filepath.Join(".", scriptName)

	//check if file exists
	if _, err := os.Stat(scriptPath); err == nil {
		fmt.Printf("[bootstrap] Script already exists:%s\n", scriptPath)
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking bootstrap script", err)
	}

	// create bootstrap script since not exist
	content := fmt.Sprintf(`#!/bin/bash
	set -e
	echo "Running bootstrap script for %s"
	
	# Update and upgrade packages
	sudo apt-get update && sudo apt-get upgrade -y
	
	`, machineName)

	err := os.WriteFile(scriptPath, []byte(content), 0755)
	if err != nil {
		return fmt.Errorf("error creating bootstrap script: %w", err)
	}

	fmt.Printf("[bootstrap] Created new boostrap script at %s\n", scriptPath)

	return nil

}
