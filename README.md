# Monitoring Tools

### Prerequisite
docker, docker-compose, dep (Golang package manager) must be installed

### How To Run 
1. Run Prometheus
```
docker run --name prometheus -p 9090:9090 -v /tmp/prometheus.yml:$HOME/go/sr
c/github.com/hengkyawijaya/prometheus/prometheus.yml prom/prometheus     
```

2. Run Datadog Agent
```
DOCKER_CONTENT_TRUST=1 docker run --net=host -d --name dd-
agent -p 8125:8125 -v /var/run/docker.sock:/var/run/docker.sock:ro -v /proc/:/host/proc/:ro -v /s
ys/fs/cgroup/:/host/sys/fs/cgroup:ro -e DD_API_KEY=<YOUR_API_KEY> 
datadog/agent:latest
```

2. Run Grafana
```
docker run \                        
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_SERVER_ROOT_URL=http://grafana.server.name" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  grafana/grafana
```
3. Run demo app
```
cd demo-app
dep ensure -v --vendor-only
go run main.go
```
4. Run simpe app 
```
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```
5. Run simple app swarm mode
```
docker-compose -f docker-compose.yml build
docker stack deploy --compose-file=docker-compose.yml monitoring-tools
```
6. Run load testing app
```
cd load-test
dep ensure -v --vendor-only
go run main.go
```
