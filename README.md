# grafana-speedtest

All in one speedtest monitor with dashboard using lightweight docker containers

### Run
```console
docker compose up
```

### Setup
```console
docker compose exec grafana grafana cli plugins install frser-sqlite-datasource
```

### Todo
1. do real speed test
2. nice dashboard
3. periodic test runs

https://community.grafana.com/t/plugins-with-grafana-grafana-docker/2153/2
