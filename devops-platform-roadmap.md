# DevOps Platform — Roadmap

> Большой учебный production-like проект для DevOps-портфолио.
>
> **Статусы:** `[ ]` не сделано · `[x]` сделано · `[~]` в процессе

---

## 0. Базовая структура проекта

- [x] Создать репозиторий `devops-platform`
- [x] Создать структуру проекта
- [x] Создать Go API
- [x] Подключить PostgreSQL
- [x] Создать Dockerfile
- [x] Создать docker-compose.yml
- [x] Добавить `.gitignore`
- [x] Добавить `.dockerignore`
- [x] Запустить приложение локально
- [x] Проверить `/health`
- [x] Проверить `/ready`
- [x] Проверить `GET /api/users`
- [x] Проверить `POST /api/users`
- [x] Сделать первый Git commit
- [x] Запушить проект на GitHub

---

# ЭТАП 1 — Backend и локальная инфраструктура

## 1.1 Go API

- [x] HTTP-сервер на Go
- [x] Endpoint `/`
- [x] Endpoint `/health`
- [x] Endpoint `/ready`
- [x] `GET /api/users`
- [x] `POST /api/users`
- [ ] Вынести конфигурацию приложения
- [ ] Добавить корректную обработку ошибок
- [ ] Добавить graceful shutdown
- [ ] Использовать `http.Server`
- [ ] Добавить request ID
- [ ] Добавить структурированные JSON-логи
- [ ] Добавить middleware для логирования запросов

## 1.2 PostgreSQL

- [x] PostgreSQL в Docker Compose
- [x] Healthcheck PostgreSQL
- [x] Подключение Go → PostgreSQL
- [x] Таблица `users`
- [ ] Перейти от `CREATE TABLE` при старте к миграциям
- [ ] Добавить индексы
- [ ] Добавить нормальную схему БД
- [ ] Продумать backup/restore
- [ ] Проверить восстановление БД из backup

## 1.3 Docker

- [x] Multi-stage Dockerfile
- [x] Минимальный runtime image
- [x] Non-root пользователь
- [x] Docker Compose
- [x] Persistent volume PostgreSQL
- [x] Healthcheck PostgreSQL
- [ ] Добавить healthcheck приложения
- [ ] Ограничить container capabilities
- [ ] Проверить образ через Dockle
- [ ] Проверить образ через Trivy

## 1.4 Локальная эксплуатация

- [x] Добавить Makefile
- [x] `make run`
- [x] `make test`
- [x] `make build`
- [x] `make lint`
- [x] `make docker-build`
- [x] `make docker-up`
- [x] `make docker-down`
- [ ] Обновить README
- [ ] Добавить архитектурную схему

---

# ЭТАП 2 — Testing + GitHub Actions CI

## 2.1 Go tests

- [x] Unit-тесты для handlers
- [ ] Unit-тесты для бизнес-логики
- [ ] Тесты ошибок
- [x] Тесты `/health`
- [x] Тесты `/ready`
- [ ] Тесты API users
- [ ] Проверить race conditions
- [ ] Добиться адекватного test coverage

## 2.2 Lint

- [ ] Подключить `golangci-lint`
- [ ] Настроить `.golangci.yml`
- [ ] Запускать lint локально
- [ ] Исправить замечания linter

## 2.3 GitHub Actions

Создать:

`.github/workflows/ci.yml`

- [ ] Запуск CI на `push`
- [ ] Запуск CI на `pull_request`
- [ ] Настроить Go environment
- [ ] `go mod download`
- [ ] `go vet`
- [ ] Unit tests
- [ ] Test coverage
- [ ] `golangci-lint`
- [ ] Проверка сборки бинарника

## 2.4 Security в CI

- [ ] Dependency vulnerability scan
- [ ] Secret scanning
- [ ] SAST
- [ ] Trivy для проекта
- [ ] Trivy для Docker image
- [ ] Dockle
- [ ] Проверить, что security job ломает pipeline при критических проблемах

## 2.5 Docker в CI

- [ ] Собирать Docker image в GitHub Actions
- [ ] Использовать BuildKit/buildx
- [ ] Добавить Docker layer cache
- [ ] Проверять Docker image после сборки
- [ ] Генерировать понятный tag для image

---

# ЭТАП 3 — Container Registry

## 3.1 GitHub Container Registry

- [ ] Настроить GHCR
- [ ] Создать image `ghcr.io/ferrtul/devops-platform`
- [ ] Настроить authentication
- [ ] Push image из GitHub Actions
- [ ] Использовать commit SHA как tag
- [ ] Добавить tag `latest` только для main
- [ ] Проверить pull image
- [ ] Добавить README с примером запуска image

## 3.2 CI/CD разделение

- [ ] Разделить CI и CD
- [ ] CI запускается на PR
- [ ] Build image после merge
- [ ] Push image после merge
- [ ] Не публиковать image при обычном PR
- [ ] Использовать GitHub Actions environments/secrets где необходимо

---

# ЭТАП 4 — Infrastructure as Code: Terraform

Цель: перестать создавать инфраструктуру вручную.

## 4.1 Terraform

- [ ] Создать `terraform/`
- [ ] Настроить provider
- [ ] Настроить variables
- [ ] Настроить outputs
- [ ] Настроить `.tfvars.example`
- [ ] Разделить dev/prod configuration
- [ ] Добавить Terraform formatting
- [ ] Добавить Terraform validation

## 4.2 Виртуальные машины

- [ ] Создать VM для Kubernetes
- [ ] Настроить сеть
- [ ] Настроить firewall/security rules
- [ ] Настроить SSH access
- [ ] Проверить Terraform plan
- [ ] Выполнить Terraform apply
- [ ] Проверить Terraform destroy

## 4.3 Terraform в GitHub Actions

- [ ] `terraform fmt`
- [ ] `terraform validate`
- [ ] `terraform plan`
- [ ] Отдельный workflow для infrastructure
- [ ] Безопасно хранить credentials
- [ ] Не хранить secrets в Git

---

# ЭТАП 5 — Ansible

Цель: автоматизировать настройку серверов.

## 5.1 Ansible structure

Создать:

```text
ansible/
├── inventory/
├── group_vars/
├── roles/
└── playbooks/
```

- [ ] Создать inventory
- [ ] Создать group_vars
- [ ] Создать роли
- [ ] Создать playbook

## 5.2 Настройка VM

- [ ] Обновление пакетов
- [ ] Создание пользователя
- [ ] Настройка SSH
- [ ] Настройка firewall
- [ ] Установка Docker
- [ ] Установка необходимых системных пакетов
- [ ] Настройка timezone
- [ ] Настройка sysctl
- [ ] Настройка логирования

## 5.3 Kubernetes preparation

- [ ] Подготовить control-plane node
- [ ] Подготовить worker node
- [ ] Установить container runtime
- [ ] Настроить необходимые kernel modules
- [ ] Настроить необходимые sysctl параметры
- [ ] Проверить idempotency Ansible

---

# ЭТАП 6 — Kubernetes

Цель: развернуть приложение не через Docker Compose, а в Kubernetes.

## 6.1 Cluster

- [ ] Создать Kubernetes cluster
- [ ] Control Plane
- [ ] Worker Node
- [ ] Проверить `kubectl get nodes`
- [ ] Настроить kubeconfig
- [ ] Проверить cluster health

## 6.2 Application

- [ ] Deployment для Go API
- [ ] Service
- [ ] ConfigMap
- [ ] Secret
- [ ] PostgreSQL Deployment/StatefulSet
- [ ] PostgreSQL Service
- [ ] PersistentVolume
- [ ] PersistentVolumeClaim

## 6.3 Production practices

- [ ] Liveness probe
- [ ] Readiness probe
- [ ] Startup probe
- [ ] Resource requests
- [ ] Resource limits
- [ ] Pod anti-affinity
- [ ] Rolling update
- [ ] PodDisruptionBudget
- [ ] RBAC
- [ ] ServiceAccount
- [ ] NetworkPolicy
- [ ] SecurityContext
- [ ] Non-root container
- [ ] Secrets management

## 6.4 Ingress

- [ ] Установить Ingress Controller
- [ ] Создать Ingress
- [ ] Настроить routing
- [ ] Настроить TLS
- [ ] Проверить внешний доступ к API

---

# ЭТАП 7 — Helm

Цель: упаковать Kubernetes application в Helm chart.

## 7.1 Chart

Создать:

```text
helm/devops-platform/
├── Chart.yaml
├── values.yaml
└── templates/
```

- [ ] Создать Helm chart
- [ ] Deployment template
- [ ] Service template
- [ ] ConfigMap template
- [ ] Secret template
- [ ] Ingress template
- [ ] PVC template
- [ ] ServiceAccount template
- [ ] RBAC templates

## 7.2 Values

- [ ] `values.yaml`
- [ ] `values-dev.yaml`
- [ ] `values-prod.yaml`
- [ ] Image repository
- [ ] Image tag
- [ ] Resource limits
- [ ] Replica count
- [ ] Ingress configuration

## 7.3 Helm CI

- [ ] `helm lint`
- [ ] `helm template`
- [ ] Проверять chart в GitHub Actions
- [ ] Проверять корректность Kubernetes manifests

---

# ЭТАП 8 — GitOps + ArgoCD

Цель: deployment должен происходить через GitOps.

## 8.1 GitOps repository structure

- [ ] Создать директорию `gitops/`
- [ ] Разделить application manifests и infrastructure
- [ ] Хранить desired state в Git
- [ ] Версионировать image tag

## 8.2 ArgoCD

- [ ] Установить ArgoCD
- [ ] Подключить Git repository
- [ ] Создать ArgoCD Application
- [ ] Настроить automatic sync
- [ ] Настроить self-healing
- [ ] Настроить pruning
- [ ] Проверить rollback

## 8.3 CI → GitOps

Pipeline:

```text
Git push
   ↓
GitHub Actions
   ↓
Tests
   ↓
Security
   ↓
Docker Build
   ↓
Push GHCR
   ↓
Update image tag in GitOps
   ↓
ArgoCD detects change
   ↓
Kubernetes deployment
```

- [ ] GitHub Actions обновляет image tag
- [ ] ArgoCD замечает изменение
- [ ] ArgoCD deploys application
- [ ] Проверить rollback
- [ ] Проверить self-healing

---

# ЭТАП 9 — Observability

Цель: полноценный monitoring/logging/alerting stack.

## 9.1 Prometheus

- [ ] Установить Prometheus
- [ ] Настроить scraping
- [ ] Метрики Go application
- [ ] HTTP request count
- [ ] HTTP request duration
- [ ] HTTP status codes
- [ ] Error rate
- [ ] In-flight requests
- [ ] Go runtime metrics
- [ ] Process CPU
- [ ] Process memory
- [ ] PostgreSQL metrics
- [ ] Kubernetes metrics
- [ ] Node metrics

## 9.2 Grafana

- [ ] Установить Grafana
- [ ] Подключить Prometheus
- [ ] Подключить Loki
- [ ] Создать dashboard application
- [ ] Создать dashboard Kubernetes
- [ ] Создать dashboard nodes
- [ ] Создать dashboard PostgreSQL
- [ ] Создать dashboard CI/CD

## 9.3 Loki + Promtail

- [ ] Установить Loki
- [ ] Настроить log collection
- [ ] Собрать application logs
- [ ] Собрать container logs
- [ ] Добавить labels
- [ ] Настроить поиск логов в Grafana

## 9.4 Alertmanager

- [ ] Установить Alertmanager
- [ ] Подключить к Prometheus
- [ ] Настроить alert rules
- [ ] High error rate
- [ ] High latency
- [ ] Service down
- [ ] High CPU
- [ ] High memory
- [ ] PostgreSQL unavailable
- [ ] Kubernetes pod crash
- [ ] Настроить Telegram notifications
- [ ] Проверить отправку alert

---

# ЭТАП 10 — Security

## 10.1 Container security

- [ ] Non-root containers
- [ ] Minimal base image
- [ ] Trivy
- [ ] Dockle
- [ ] Scan dependencies
- [ ] Scan filesystem
- [ ] Scan Docker image
- [ ] Fix critical vulnerabilities

## 10.2 Kubernetes security

- [ ] RBAC
- [ ] NetworkPolicy
- [ ] SecurityContext
- [ ] Read-only filesystem где возможно
- [ ] Drop Linux capabilities
- [ ] Resource limits
- [ ] Pod Security standards
- [ ] Secrets не хранятся в Git
- [ ] Проверить permissions

## 10.3 CI security

- [ ] Dependabot/Renovate
- [ ] Secret scanning
- [ ] SAST
- [ ] Dependency scanning
- [ ] Container scanning
- [ ] Least privilege GitHub Actions permissions

---

# ЭТАП 11 — Load testing и отказоустойчивость

## 11.1 Load testing

- [ ] Установить k6
- [ ] Написать load test
- [ ] Тест GET `/`
- [ ] Тест GET `/health`
- [ ] Тест GET `/api/users`
- [ ] Тест POST `/api/users`
- [ ] Проверить latency
- [ ] Проверить throughput
- [ ] Проверить error rate

## 11.2 Failure testing

- [ ] Убить application pod
- [ ] Проверить Kubernetes restart
- [ ] Убить PostgreSQL
- [ ] Проверить application behavior
- [ ] Проверить readiness
- [ ] Проверить liveness
- [ ] Проверить ArgoCD self-healing
- [ ] Проверить rolling update
- [ ] Проверить rollback

## 11.3 HPA

- [ ] Установить metrics-server
- [ ] Настроить HPA
- [ ] Увеличить нагрузку
- [ ] Проверить scale-out
- [ ] Уменьшить нагрузку
- [ ] Проверить scale-in

---

# ЭТАП 12 — Backup / Disaster Recovery

- [ ] Backup PostgreSQL
- [ ] Автоматический backup
- [ ] Хранение backup отдельно от БД
- [ ] Проверить restore
- [ ] Описать процедуру восстановления
- [ ] Проверить восстановление приложения
- [ ] Проверить восстановление Kubernetes resources
- [ ] Описать RPO/RTO

---

# ЭТАП 13 — Финальный CI/CD

Итоговый pipeline:

```text
Developer
    │
    ▼
GitHub
    │
    ▼
GitHub Actions
    │
    ├── Lint
    ├── Unit Tests
    ├── SAST
    ├── Dependency Scan
    ├── Docker Build
    └── Trivy
            │
            ▼
          GHCR
            │
            ▼
       Update GitOps
            │
            ▼
          ArgoCD
            │
            ▼
       Kubernetes
            │
       ┌────┴────┐
       ▼         ▼
    Go API   PostgreSQL
       │
       └──────┬─────────────┐
              ▼             ▼
          Prometheus       Loki
              │             │
              └──────┬──────┘
                     ▼
                  Grafana
                     │
                     ▼
                Alertmanager
                     │
                     ▼
                  Telegram
```

- [ ] PR → lint
- [ ] PR → tests
- [ ] PR → security
- [ ] Merge → build image
- [ ] Merge → push GHCR
- [ ] Update GitOps repository
- [ ] ArgoCD deploy
- [ ] Kubernetes health checks
- [ ] Monitoring
- [ ] Alerting
- [ ] Automatic recovery

---

# ЭТАП 14 — Documentation

Это важно для портфолио.

## README

- [ ] Описание проекта
- [ ] Цели проекта
- [ ] Architecture diagram
- [ ] Technology stack
- [ ] Local development
- [ ] Docker запуск
- [ ] Kubernetes запуск
- [ ] Helm
- [ ] ArgoCD
- [ ] CI/CD
- [ ] Monitoring
- [ ] Alerting
- [ ] Security
- [ ] Load testing
- [ ] Backup/restore
- [ ] Troubleshooting
- [ ] Screenshots Grafana
- [ ] Screenshots ArgoCD
- [ ] Screenshots GitHub Actions
- [ ] Screenshots Kubernetes
- [ ] Example API requests

## ADR

- [ ] Почему Go
- [ ] Почему PostgreSQL
- [ ] Почему Kubernetes
- [ ] Почему Helm
- [ ] Почему ArgoCD
- [ ] Почему Prometheus/Grafana/Loki
- [ ] Почему GitHub Actions
- [ ] Почему Terraform + Ansible

---

# ЭТАП 15 — Финальная полировка проекта

- [ ] Проверить все README
- [ ] Убрать секреты из Git
- [ ] Проверить `.gitignore`
- [ ] Проверить Dockerfile
- [ ] Проверить Terraform
- [ ] Проверить Ansible
- [ ] Проверить Helm
- [ ] Проверить Kubernetes manifests
- [ ] Проверить GitHub Actions
- [ ] Проверить alerts
- [ ] Проверить backup
- [ ] Проверить rollback
- [ ] Проверить disaster recovery
- [ ] Добавить architecture diagram
- [ ] Добавить screenshots
- [ ] Добавить badges в README
- [ ] Создать GitHub Release
- [ ] Сделать финальный demo

---

# Финальный чек-лист навыков

После завершения проекта в резюме можно будет обоснованно указать:

- [ ] Go
- [ ] Linux
- [ ] Bash
- [ ] Git
- [ ] GitHub Actions
- [ ] CI/CD
- [ ] Docker
- [ ] Docker Compose
- [ ] GHCR
- [ ] Kubernetes
- [ ] Helm
- [ ] ArgoCD
- [ ] GitOps
- [ ] Terraform
- [ ] Ansible
- [ ] PostgreSQL
- [ ] Prometheus
- [ ] Grafana
- [ ] Loki
- [ ] Alertmanager
- [ ] Trivy
- [ ] Dockle
- [ ] Nginx / Ingress
- [ ] RBAC
- [ ] NetworkPolicy
- [ ] HPA
- [ ] Load testing
- [ ] Backup / Restore
- [ ] Disaster Recovery

---

# Правило работы над проектом

Не делать всё сразу.

Работаем строго по этапам:

**1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10 → 11 → 12 → 13 → 14 → 15**

После каждого этапа:

1. Проверяем, что всё работает.
2. Делаем Git commit.
3. Пушим изменения.
4. Обновляем этот checklist.
5. Только после этого переходим дальше.

## Текущий статус

**Этап 1 — практически завершён.**

**Следующая задача: Этап 2.1 — Makefile + Go tests.**