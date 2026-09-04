# devops-platform

                         Git
                          │
                          ▼
                  ┌───────────────┐
                  │    GitLab     │
                  └───────┬───────┘
                          │
                          ▼
                  ┌───────────────┐
                  │   GitLab CI   │
                  └───────┬───────┘
                          │
             ┌────────────┼────────────┐
             ▼            ▼            ▼
           Lint         Tests      Security
                                      │
                                      ▼
                              Docker image
                                      │
                                      ▼
                              Container Registry
                                      │
                                      ▼
                                Helm chart
                                      │
                                      ▼
                              GitOps repository
                                      │
                                      ▼
                                  ArgoCD
                                      │
                                      ▼
                           ┌──────────────────┐
                           │    Kubernetes    │
                           │                  │
                           │ ┌──────────────┐ │
                           │ │   Ingress    │ │
                           │ └──────┬───────┘ │
                           │        ▼         │
                           │ ┌──────────────┐ │
                           │ │  Application │ │
                           │ └──────┬───────┘ │d
                           │        │         │
                           │        ▼         │
                           │   PostgreSQL     │
                           └──────────────────┘
                                      │
                   ┌──────────────────┼──────────────────┐
                   ▼                  ▼                  ▼
              Prometheus           Loki             Exporters
                   │                  │
                   ▼                  ▼
                Grafana          Grafana
                   │
                   ▼
              Alertmanager
                   │
                   ▼
                Telegram


        Infrastructure
              │
              ▼
          Terraform
              │
              ▼
        Virtual Machines
              │
              ▼
           Ansible
              │
              ▼
     Docker / Kubernetes



     PHASE 1
Go application
       ↓
Docker
       ↓
PostgreSQL
       ↓
Git

↓

PHASE 2
GitLab CI
       ↓
Tests
       ↓
Security
       ↓
Docker Registry

↓

PHASE 3
Terraform
       ↓
VM
       ↓
Ansible
       ↓
Kubernetes

↓

PHASE 4
Helm
       ↓
ArgoCD
       ↓
GitOps

↓

PHASE 5
Prometheus
Grafana
Loki
Alertmanager

↓

PHASE 6
Security
Load testing
Failure testing
Documentation