🎯 Objetivo do projeto

Criar uma API REST com GoLang, simulando um sistema de catálogo de produtos, com foco em:

Estrutura idiomática e modular

Boas práticas (handlers, serviços, repositórios)

Testes básicos

Observabilidade inicial (logs, métricas)

Pronto para deploy simples (Docker)

📁 Estrutura proposta do projeto

go-api-produtos/
│
├── cmd/                 # Entrada principal da aplicação
│   └── main.go
│
├── internal/            # Domínio e regras do negócio
│   ├── produto/         # Pacote de produto
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│
├── pkg/                 # Códigos utilitários (logger, observabilidade, etc)
│   ├── logger/
│   └── metrics/
│
├── config/              # Configurações
│   └── config.go
│
├── Dockerfile
├── go.mod
└── README.md

🔧 Tecnologias usadas
GoLang 1.21+

Gin ou Fiber (router HTTP)

Zap (logs estruturados)

Prometheus client (métricas básicas)

Testify (testes unitários)

Docker (deploy local)