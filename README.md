# ☁️ Cloud Financial Analytics

A Go-based cloud financial analytics and infrastructure monitoring platform designed to track cloud costs, budgets, resources, alerts, system health, and performance metrics.

The project combines a modular Go backend with Prometheus for metrics collection and Grafana for visualization.

---

## 🚀 Features

- 💰 Cloud cost tracking and analytics
- 💵 Budget monitoring
- 🚨 Cost and system alerts
- 🖥️ Cloud resource monitoring
- ❤️ System health monitoring
- 📊 Application and system metrics
- 📡 Prometheus metrics integration
- 📈 Grafana dashboards
- 🐳 Docker support
- ⚙️ Docker Compose support
- 🧩 Modular Go backend architecture

---

## 🏗️ Architecture

```text
                    ┌─────────────────────┐
                    │       Client        │
                    │ Browser / Postman   │
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │      Go REST API    │
                    │  Routes / Handlers  │
                    └──────────┬──────────┘
                               │
              ┌────────────────┼────────────────┐
              │                │                │
              ▼                ▼                ▼
        ┌──────────┐     ┌──────────┐     ┌──────────┐
        │   Cost   │     │  Budget  │     │ Resource │
        │ Service  │     │ Service  │     │ Service  │
        └──────────┘     └──────────┘     └──────────┘
              │                │                │
              └────────────────┼────────────────┘
                               │
                               ▼
                       ┌──────────────┐
                       │  Monitoring  │
                       │   Metrics    │
                       └──────┬───────┘
                              │
                              ▼
                       ┌──────────────┐
                       │  Prometheus  │
                       └──────┬───────┘
                              │
                              ▼
                       ┌──────────────┐
                       │   Grafana    │
                       │  Dashboards  │
                       └──────────────┘
```

---

## 📁 Project Structure

```text
Cloud-Financial-Analytics/
│
├── cmd/
│   └── main.go
│
├── configs/
│
├── docs/
│
├── grafana/
│   ├── dashboards/
│   │
│   └── provisioning/
│       ├── dashboards/
│       │   └── dashboard.yml
│       │
│       └── datasources/
│           └── datasource.yml
│
├── internal/
│   │
│   ├── api/
│   │   ├── alerts.go
│   │   ├── budget.go
│   │   ├── cost.go
│   │   ├── health.go
│   │   ├── resource.go
│   │   ├── routes.go
│   │   └── system-metrics.go
│   │
│   ├── config/
│   │
│   ├── models/
│   │   ├── alerts.go
│   │   ├── budget.go
│   │   ├── cost.go
│   │   ├── health.go
│   │   ├── metrics.go
│   │   └── resource.go
│   │
│   ├── monitoring/
│   │   ├── collector.go
│   │   └── metrics.go
│   │
│   ├── services/
│   │   ├── alerts_service.go
│   │   ├── budget_service.go
│   │   ├── cost_service.go
│   │   ├── health_service.go
│   │   ├── metrics_service.go
│   │   └── resource_service.go
│   │
│   └── system/
│
├── prometheus/
│   └── prometheus.yml
│
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## 🧩 Core Modules

### 💰 Cost Management

Handles cloud cost-related data and analytics.

```text
internal/api/cost.go
internal/services/cost_service.go
internal/models/cost.go
```

The cost module provides the foundation for monitoring cloud spending and analyzing cost information.

### 💵 Budget Management

Handles budget-related operations and monitoring.

```text
internal/api/budget.go
internal/services/budget_service.go
internal/models/budget.go
```

It can be used to compare spending against configured budgets and identify potential overspending.

### 🚨 Alert Management

Handles cost and infrastructure alerts.

```text
internal/api/alerts.go
internal/services/alerts_service.go
internal/models/alerts.go
```

Possible alert conditions include:

- Budget threshold exceeded
- High resource usage
- Abnormal metrics
- Cost anomalies

### 🖥️ Resource Management

Handles information related to monitored cloud resources.

```text
internal/api/resource.go
internal/services/resource_service.go
internal/models/resource.go
```

### ❤️ Health Monitoring

Provides application and infrastructure health monitoring.

```text
internal/api/health.go
internal/services/health_service.go
internal/models/health.go
```

Example health endpoint:

```text
GET /api/health
```

Example response:

```json
{
  "status": "ok"
}
```

### 📊 Metrics & Monitoring

The monitoring subsystem collects application and system metrics.

```text
internal/monitoring/
├── collector.go
└── metrics.go
```

System metrics are also handled through:

```text
internal/api/system-metrics.go
```

---

## 📡 Prometheus

Prometheus is used to collect and store monitoring metrics.

Configuration:

```text
prometheus/prometheus.yml
```

Monitoring flow:

```text
Go Application
      │
      │ Metrics
      ▼
Prometheus
      │
      │ PromQL
      ▼
Grafana
      │
      ▼
Monitoring Dashboard
```

---

## 📊 Grafana

Grafana is used to visualize application and infrastructure metrics.

Configuration:

```text
grafana/
├── dashboards/
└── provisioning/
    ├── dashboards/
    │   └── dashboard.yml
    └── datasources/
        └── datasource.yml
```

Dashboards and data sources can be provisioned automatically when the monitoring stack starts.

---

## 🐳 Docker

The project includes a Dockerfile for containerized deployment.

Build the image:

```bash
docker build -t cloud-financial-analytics .
```

Run the container:

```bash
docker run -p 8080:8080 cloud-financial-analytics
```

---

## 🐳 Docker Compose

Start the complete environment:

```bash
docker compose up --build
```

Run in detached mode:

```bash
docker compose up -d --build
```

Stop the environment:

```bash
docker compose down
```

---

## 🛠️ Local Development

### 1. Clone the repository

```bash
git clone https://github.com/krishnagupta-007/Cloud-Financial-Analytics.git
```

### 2. Enter the project directory

```bash
cd Cloud-Financial-Analytics
```

### 3. Download Go dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run ./cmd
```

The application will start on the configured port.

---

## 🔍 API Testing

The APIs can be tested using:

- Postman
- cURL
- Browser
- REST API clients

### Health Check

```bash
curl http://localhost:8080/api/health
```

Example response:

```json
{
  "status": "ok"
}
```

---

## 📌 API Modules

| Module | Purpose |
|---|---|
| Health | Application health |
| Cost | Cloud cost information |
| Budget | Budget monitoring |
| Alerts | Cost and system alerts |
| Resources | Resource information |
| Metrics | Application and system metrics |

> Exact API routes are defined in `internal/api/routes.go`.

---

## 🧰 Tech Stack

### Backend

- Go
- REST API
- JSON

### Monitoring & Observability

- Prometheus
- Grafana

### DevOps

- Docker
- Docker Compose

### Development Tools

- Git
- GitHub
- VS Code
- Postman

---

## 🏛️ Backend Architecture

The backend follows a modular layered architecture:

```text
API Layer
    ↓
Service Layer
    ↓
Model Layer
    ↓
Monitoring / Infrastructure
```

### API Layer

Handles HTTP requests and responses.

```text
internal/api/
```

### Service Layer

Contains application and business logic.

```text
internal/services/
```

### Model Layer

Defines application data structures.

```text
internal/models/
```

### Monitoring Layer

Handles metric collection and monitoring.

```text
internal/monitoring/
```

---

## 📈 Monitoring Flow

```text
                    ┌───────────────┐
                    │ Go Application│
                    └───────┬───────┘
                            │
                     Exposes Metrics
                            │
                            ▼
                    ┌───────────────┐
                    │  Prometheus   │
                    └───────┬───────┘
                            │
                       PromQL Queries
                            │
                            ▼
                    ┌───────────────┐
                    │    Grafana    │
                    └───────┬───────┘
                            │
                            ▼
                    Monitoring Dashboard
```

---

## 🎯 Use Cases

Cloud Financial Analytics can be used as a foundation for:

- Cloud cost monitoring
- Budget tracking
- Infrastructure monitoring
- Resource utilization analysis
- Operational dashboards
- Cost threshold alerts
- Cloud FinOps applications
- DevOps monitoring systems

---

## 🔮 Future Improvements

- ☁️ AWS cost integration
- ☁️ Azure cost integration
- ☁️ Google Cloud cost integration
- 🤖 AI-based cost anomaly detection
- 📉 Automated cost optimization recommendations
- 🔔 Email and Slack notifications
- 🔐 Authentication and authorization
- 🗄️ PostgreSQL integration
- 📊 Advanced Grafana dashboards
- ☸️ Kubernetes deployment
- 🌍 Multi-cloud cost comparison
- 📦 Terraform infrastructure deployment
- 🔄 GitHub Actions CI/CD

---

## 🔐 Security

Never commit sensitive credentials to GitHub.

Examples:

```text
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
DATABASE_PASSWORD
API_KEYS
TOKENS
```

Use environment variables or a secure secrets-management solution instead.

---

## 📦 Requirements

Before running the project locally, install:

- Go 1.22+
- Docker
- Docker Compose
- Git

For monitoring:

- Prometheus
- Grafana

Docker Compose can be used to simplify the setup.

---

## 🤝 Contributing

Contributions are welcome.

Create a feature branch:

```bash
git checkout -b feature/new-feature
```

Make your changes and commit:

```bash
git add .
git commit -m "Add new feature"
```

Push your branch:

```bash
git push origin feature/new-feature
```

Then create a Pull Request.

---

## 📄 License

This project is intended for educational, development, and portfolio purposes.

---

## 👨‍💻 Author

**Krishna Gupta**

B.Sc. Computer Science Student | Aspiring DevOps / Cloud Engineer

### Focus Areas

```text
Go
Cloud Computing
DevOps
Docker
Kubernetes
AWS
Prometheus
Grafana
Backend Development
Cloud FinOps
```

---

## ⭐ Project

If you find this project useful, consider giving the repository a star.

```text
Cloud Financial Analytics
Go + Prometheus + Grafana + Docker
```