package system

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func RunScript(file string) (string, error) {
	cmd, err := exec.Command("/bin/bash", file).Output()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "system",
			"file":     "bash.go",
			"function": "RunScript",
			"message":  err,
		}).Errorf("error while execute script")
		return "", err
	}

	output := string(cmd)

	return output, nil
}
