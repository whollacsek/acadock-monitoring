package cpu

import (
  "Acadock/lxc"
  "fmt"
  "log"
  "os"
  "runtime"
  "strconv"
  "time"
)

const (
	LXC_CPUACCT_DIR        = "/sys/fs/cgroup/cpuacct/lxc"
	LXC_CPUACCT_USAGE_FILE = "cpuacct.usage"
  REFRESH_TIME           = 1
)

var (
  previousCpuUsages = make(map[string]int64)
  cpuUsages = make(map[string]int64)
)


func cpuacctUsage(container string) int64 {
  file := fmt.Sprintf("%s/%s/%s", LXC_CPUACCT_DIR, container, LXC_CPUACCT_USAGE_FILE)
  f, err := os.Open(file)
  if err != nil {
    log.Fatalln(err)
  }

  buffer := make([]byte, 16)
  n, err := f.Read(buffer)
  buffer = buffer[:n]

  bufferStr := string(buffer)
  bufferStr = bufferStr[:len(bufferStr)-1]

  res, err := strconv.ParseInt(bufferStr, 10, strconv.IntSize)
  if err != nil {
    log.Fatalln("Fail to parse : ", err)
  }
  return res
}

func Monitor() {
  tick := time.NewTicker(REFRESH_TIME * time.Second)
  for {
    <-tick.C
    for k, v := range cpuUsages {
       previousCpuUsages[k] = v
    }

    containers, err := lxc.ListContainer()
    if err != nil {
      log.Fatalln(err)
    }

    for _, container := range containers {
      cpuUsages[container] = cpuacctUsage(container)
    }
  }
}

func GetUsage(name string) int64 {
  return int64((float64((cpuUsages[name] - previousCpuUsages[name])) / float64(1e9) / float64(runtime.NumCPU())) * 100)
}