package docker

import (
	"io/ioutil"
	"strings"

	"github.com/Scalingo/acadock-monitoring/config"
	"github.com/Scalingo/acadock-monitoring/debug"
)

func Pid(id string) (string, error) {
	path := config.CgroupPath("memory", id)
	content, err := ioutil.ReadFile(path + "/tasks")
	if err != nil {
		return "", err
	}
	debug.Printf("Content of tasks file for %v: %v", id, string(content))
	return strings.Split(string(content), "\n")[0], nil
}
