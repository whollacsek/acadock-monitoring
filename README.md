Acadock Monitoring - Docker container monitoring
================================================

This webservice provides live data on Docker containers. It takes
data from the Linux kernel control groups and from the namespace of
the container and expose them through a HTTP API.

> The solution is still a work in progress.

Configuration
-------------

From environment

* `PORT`: port to bind (4244 by default)
* `DOCKER_URL`: docker endpoint (http://127.0.0.1:4243 by default)
* `REFRESH_TIME`: number of second between CPU/net refresh (1 by default)
* `CGROUP_SOURCE`: "docker" or "systemd" (docker by default)
  docker:  /sys/fs/cgroup/:cgroup/memory/docker
  systemd: /sys/fs/cgroup/:cgroup/memory/system.slice/docker-#{id}.slice

API
---

* Memory consumption (in bytes)

    `GET /containers/:id/mem`

    Return 200 OK
    Content-Type: text/plain

* CPU usage (percentage)

    Return 200 OK
    Content-Type: text/plain
    `GET /containers/:id/cpu`

* Network usage (bytes and percentage)

    Return 200 OK
    Content-Type: application/json
    `GET /containers/:id/net`

### Developers

> Léo Unbekandt `<leo@unbekandt.eu>`
