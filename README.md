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