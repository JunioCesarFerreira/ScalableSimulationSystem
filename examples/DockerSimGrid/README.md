# DockerSimGrid

DockerSimGrid é um protótipo de sistema para gerenciamento de simulações distribuídas utilizando **Docker, Kafka e SSH**, permitindo a orquestração de simulações em múltiplos nós.

**Atenção!** Esta versão é um protótipo e ainda está incompleta.

## 📌 Arquitetura
O sistema é composto pelos seguintes componentes:

1. **MongoDB (Opcional)**: Pode ser usado para armazenamento persistente de tarefas e resultados.
2. **Kafka**: Middleware de mensagens para comunicação assíncrona entre serviços.
3. **MasterNode (Go)**: Responsável por gerenciar as simulações, interagir com Docker e executar processos via SSH.
4. **DataCollector**: Coleta e armazena os dados das simulações para posterior análise.
5. **UbuntuDocker**: Imagem base para execução das simulações dentro de contêineres.

📜 **Fluxo de trabalho:**
1. O **MasterNode** recebe solicitações de simulação via Kafka.
2. Ele instancia e gerencia contêineres Docker para cada simulação.
3. Se necessário, pode usar **SSH** para executar comandos remotamente.
4. Os resultados são coletados pelo **DataCollector** e armazenados.

## 🚀 Configuração e Execução
### **1. Clonar o Repositório**
```sh
 git clone https://github.com/JunioCesarFerreira/ScalableSimulationSystem
 cd examples/DockerSimGrid
```

### **2. Construir e Iniciar os Containers**
```sh
 docker-compose up --build -d
```

### **3. Ver Logs dos Serviços**
- **MasterNode:**
```sh
 docker-compose logs -f master-node
```
- **Kafka:**
```sh
 docker-compose logs -f kafka
```
- **DataCollector:**
```sh
 docker-compose logs -f data-collector
```

## 📂 Estrutura do Projeto
```
DockerSimGrid/
│── docker-compose.yaml
│── MasterNode/
│   │── Dockerfile
│   │── go.mod
│   │── go.sum
│   │── cmd/
│   │   │── main.go
│   │── pkg/
│   │   │── dockerclient/
│   │   │── kafkaclient/
│   │   │── sshhandler/
│── DataCollector/
│   │── Dockerfile
│   │── main.py
│── UbuntuDocker/
│   │── Dockerfile
```

## 🔍 Consultando os Dados
Se o MongoDB estiver configurado para armazenamento:
```sh
docker exec -it mongodb mongosh
```
**Consultar tarefas e resultados:**
```sh
use simulation_db
db.simulations_tasks.find().pretty()
db.simulations_results.find().pretty()
```

## 🛠 Teste Inicial

### 1. Construir a imagem UbuntuDocker

No diretório `UbuntuDocker`, execute:

```bash
docker build -t ubuntu-docker .
```

### 2. Construir com Docker Compose

No diretório `SimGrid`, execute:

```bash
docker compose build 
```

### 3. Executar os contêineres com Docker Compose

No diretório `SimGrid`, execute:

```bash
docker compose up -d
```

## 📜 Licença
Este projeto está licenciado sob a [Licença MIT](../../LICENSE).

