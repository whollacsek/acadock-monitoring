package mem

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Scalingo/acadock-monitoring/config"
	"github.com/Scalingo/acadock-monitoring/docker"
)

const (
	LXC_MEM_USAGE_FILE = "memory.usage_in_bytes"
)

var (
	containerIds []string
)

type MemUsage struct {
	Id    string
	Usage int64
}

func Monitor() {
	containers := docker.RegisterToContainersStream()
	for c := range containers {
		fmt.Println("monitor mem", c)
		containerIds = append(containerIds, c)
	}
}

func GetUsage(id string) (int64, error) {
	id, err := docker.ExpandId(id)
	if err != nil {
		log.Println("Error when expanding id:", err)
		return 0, err
	}

	path := config.CgroupPath("memory", id) + "/" + LXC_MEM_USAGE_FILE
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error while opening:", err)
		return 0, err
	}
	defer f.Close()

	buffer := make([]byte, 16)
	n, err := f.Read(buffer)
	if err != nil {
		log.Println("Error while reading ", path, ":", err)
		return 0, err
	}

	buffer = buffer[:n-1]
	val, err := strconv.ParseInt(string(buffer), 10, 64)
	if err != nil {
		log.Println("Error while parsing ", string(buffer), " : ", err)
		return 0, err
	}

	return val, nil
}

func GetAllUsage() ([]MemUsage, error) {
	var usages []MemUsage
	var updatedContainerIds []string
	for _, id := range containerIds {
		mem, err := GetUsage(id)
		if err != nil {
			log.Println("Error while getting mem usage for ", id)
		} else {
			usage := MemUsage{id, mem}
			usages = append(usages, usage)
			updatedContainerIds = append(updatedContainerIds, id)
		}
	}
	containerIds = updatedContainerIds

	return usages, nil
}
