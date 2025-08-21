# Real-Time Log Monitoring and Alerting Platform

## Phase 1: Linux & Scripting (Base Environment)
- **Goal: Prepare a reproducible Linux environment for local dev.**
- Install Minikube.
- Write Bash scripts for:
    - Setting up environment variables.
    - Deploying Kafka locally.
    - Automating log generation for testing.

## Phase 2: Microservices (Golang + Python)
- **Service 1 (Golang):**
- Simple HTTP API returning JSON.
-Includes **/health** and **/metrics** endpoints (expose Prometheus metrics using **prometheus/client_golang**).
-Generates logs (info, warn, error) with structured JSON format.
- **Service 2 (Python):**
-  Kafka producer that reads from a local log file and pushes to Kafka.
-  Logs in a consistent format (**timestamp**, **service**, **level**, **message**).

## Phase 3: Kafka Integration
- Deploy a **Kafka cluster** in Kubernetes using Helm (**bitnami/kafka** or **confluentinc/cp-helm-charts**).
- Configure both services to produce logs to a **logs-topic**.
- Add a Python Kafka consumer that:
    - Reads from **logs-topic**.
    - Filters only **ERROR** logs.
    - Pushes them to Loki via HTTP API.

## Phase 4: Kubernetes Deployment
- Containerize both microservices.
- Write **Kubernetes manifests** (**Deployment**, **Service**, **ConfigMap**).
- Deploy to Minikube using **kubectl**.
- Add **Helm charts** so the entire stack can be deployed with:
```bash
helm install realtime-logging ./charts/realtime-logging
```

## Phase 5: Observability (Loki + Prometheus)
- **Prometheus:**
    - Install with Helm (**prometheus-community/kube-prometheus-stack**).
    - Scrape metrics from Golang service and K8s cluster.
    - Create recording rules for CPU/memory alerts.
- **Loki:**
    - Install with Helm (**grafana/loki-stack**).
    - Configure Python Kafka consumer to push logs to Loki.
    - Create Grafana dashboards to query logs.

## Phase 6: Alerting
- Set up **Prometheus Alertmanager** to send alerts on:
    - High error rate in logs.
    - High CPU usage in services.
- Send alerts to Slack or email.

## Phase 7: Automation & Testing
- Add Bash scripts for:
    - Full environment setup.
    - Log load-testing.
- Add Linux troubleshooting tasks:
    - Simulate node failures in Minikube.
    - Kill pods and check auto-recovery.