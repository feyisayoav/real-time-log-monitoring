## services/
**Golang service**
- HTTP endpoints (/, /health, /metrics).
- Prometheus client library for metrics.
- Logs to stdout in structured JSON.

**Python producer**
- Reads application.log and sends each line as a Kafka message.

**Python consumer**
- Reads from Kafka.
- Filters ERROR level logs.
- Sends them to Loki via REST API.

**k8s-manifests/** & **helm/**
- First deploy raw YAMLs to debug.
- Then package into Helm charts so you can:
```bash
helm install realtime-logging ./helm/charts
```
**scripts/**
- Automate everything:
- **setup_env.sh** → Install Minikube, Helm, kubectl.
- **deploy_all.sh** → Install Kafka, services, Loki, Prometheus, Grafana.
- **generate_logs.sh** → Simulate traffic.
- **port_forward.sh** → View Grafana/Prometheus locally.
- **cleanup.sh** → Tear down environment.

**dashboards/**
- Prebuilt Grafana JSON exports so you can instantly see:
- Logs from specific services.
- Error rates over time.
- CPU/memory metrics.

