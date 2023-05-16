// Package theme provides functions for changing the system's GTK theme.
package theme

import (
	"autogtk/log"
	"os/exec"
)

var logger = log.New()

// Set changes the system's GTK theme to the provided theme name.
// It uses the "gsettings" command-line tool to change the theme.
func Set(name string) error {
	cmd := exec.Command(
		"gsettings",
		"set",
		"org.gnome.desktop.interface",
		"gtk-theme",
		name,
	)
	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error running command: %v, error: %v", cmd, err)
		return err
	}
	logger.Infof("%s theme set.", name)
	return nil
}
