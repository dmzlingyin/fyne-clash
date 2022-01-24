package api

import (
	"os"
	"os/exec"
)

func AutoStart(enable bool) error {
	if enable {
		if _, err := os.Stat("/home/lingyin/.xprofile"); os.IsNotExist(err) {
			return err
		}

		cmd := exec.Command("bash", "-c", `echo "clashG &" >> /home/lingyin/.xprofile`)
		file, err := os.OpenFile("/home/lingyin/.xprofile", os.O_WRONLY, 0777)
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
