package config

import (
	"os"
	"strconv"
)

var ENV = map[string]string{
	"DOCKER_URL":    "http://localhost:4243",
	"PORT":          "4244",
	"REFRESH_TIME":  "5",
	"CGROUP_SOURCE": "docker",
	"CGROUP_DIR":    "/sys/fs/cgroup",
	"PROC_DIR":      "/proc",
	"RUNNER_DIR":    "/usr/bin",
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
		return ENV["CGROUP_DIR"] + "/" + cgroup + "/docker/" + id
	} else if ENV["CGROUP_SOURCE"] == "systemd" {
		return ENV["CGROUP_DIR"] + "/" + cgroup + "/system.slice/docker-" + id + ".scope"
	} else {
		panic("unknown cgroup source" + ENV["CGROUP_SOURCE"])
	}
}
