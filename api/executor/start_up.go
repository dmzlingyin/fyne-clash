package executor

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

const profile = "/home/lingyin/.xprofile"

func IsStartUp() bool {
	if _, err := os.Stat(profile); os.IsNotExist(err) {
		return false
	}

	file, err := os.Open(profile)
	if err != nil {
		return false
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return false
	}

	return strings.Contains(string(content), "clashG &")
}

func AutoStart(enable bool) error {
	if enable {
		if _, err := os.Stat(profile); os.IsNotExist(err) {
			return err
		}

		cmd := exec.Command("bash", "-c", `echo "clashG &" >> /home/lingyin/.xprofile`)
		file, err := os.OpenFile(profile, os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		defer file.Close()

		if err = cmd.Start(); err != nil {
			return err
		}
		cmd.Wait()
	} else {
		cmd := exec.Command("bash", "-c", `sed -i 's/clashG &//' /home/lingyin/.xprofile`)
		if err := cmd.Start(); err != nil {
			return err
		}
		cmd.Wait()
	}
	return nil
}
