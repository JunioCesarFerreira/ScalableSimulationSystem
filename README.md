# ScalableSimulationSystem

 A scalable and distributed simulation system using Kafka, Docker, Python, and Golang to orchestrate tasks, process simulations, and collect results efficiently.

## Initial Test

### 1. Build Ubuntu Docker

In directory `UbuntuDocker` run:

```bash
docker build -t ubuntu-docker .
```

### 2. Build with Docker Compose

In directory `SimGrid` run:

```bash
docker compose build 
```

### 3. Run Containers with Docker Compose

In directory `SimGrid` run:

```bash
docker compose up -d
```