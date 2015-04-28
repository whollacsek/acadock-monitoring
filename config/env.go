package config

import (
	"os"
	"strconv"
)

var ENV = map[string]string{
	"DOCKER_URL":       "http://localhost:4243",
	"PORT":             "4244",
	"REFRESH_TIME":     "5",
	"CGROUP_SOURCE":    "docker",
	"BASE_CGROUP_PATH": "/sys/fs/cgroup",
}

var RefreshTime int

func init() {
	for k, v := range ENV {
		if os.Getenv(k) != "" {
			ENV[k] = os.Getenv(k)
		} else {
			os.Setenv(k, v)
		}
	}

	var err error
	RefreshTime, err = strconv.Atoi(ENV["REFRESH_TIME"])
	if err != nil {
		panic(err)
	}
}

func CgroupPath(cgroup string, id string) string {
	if ENV["CGROUP_SOURCE"] == "docker" {
		return ENV["BASE_CGROUP_PATH"] + "/" + cgroup + "/docker/" + id
	} else if ENV["CGROUP_SOURCE"] == "systemd" {
		return ENV["BASE_CGROUP_PATH"] + "/" + cgroup + "/system.slice/docker-" + id + ".scope"
	} else {
		panic("unknown cgroup source" + ENV["CGROUP_SOURCE"])
	}
}
