# DockerSimGrid

🌍 *[Português](README_pt.md)*

DockerSimGrid is a prototype system for managing distributed simulations using **Docker, Kafka, and SSH**, enabling the orchestration of simulations across multiple nodes.

**Attention!** This version is a prototype and is still incomplete.

## 📌 Architecture
The system consists of the following components:

1. **MongoDB (Optional)**: Can be used for persistent storage of tasks and results.
2. **Kafka**: Message middleware for asynchronous communication between services.
3. **MasterNode (Go)**: Responsible for managing simulations, interacting with Docker, and executing processes via SSH.
4. **DataCollector**: Collects and stores simulation data for later analysis.
5. **UbuntuDocker**: Base image for running simulations inside containers.

📜 **Workflow:**
1. The **MasterNode** receives simulation requests via Kafka.
2. It instantiates and manages Docker containers for each simulation.
3. If necessary, it can use **SSH** to execute commands remotely.
4. The results are collected by the **DataCollector** and stored.

## 🚀 Setup and Execution
### **1. Clone the Repository**
```sh
git clone https://github.com/JunioCesarFerreira/ScalableSimulationSystem
cd examples/DockerSimGrid
```

### **2. Build and Start Containers**
```sh
docker-compose up --build -d
```

### **3. View Service Logs**
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

## 📂 Project Structure
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

## 🔍 Querying Data
If MongoDB is configured for storage:
```sh
docker exec -it mongodb mongosh
```
**Query tasks and results:**
```sh
use simulation_db
db.simulations_tasks.find().pretty()
db.simulations_results.find().pretty()
```

## 🛠 Initial Test

### 1. Build the UbuntuDocker image

In the `UbuntuDocker` directory, run:

```bash
docker build -t ubuntu-docker .
```

### 2. Build with Docker Compose

In the `SimGrid` directory, run:

```bash
docker compose build
```

### 3. Run Containers with Docker Compose

In the `SimGrid` directory, run:

```bash
docker compose up -d
```

## 📜 License
This project is licensed under the [MIT License](../../LICENSE).

