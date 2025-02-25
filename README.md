# ScalableSimulationSystem

🌍 *[Português](README_pt.md)*

ScalableSimulationSystem is a set of tools for studying and implementing scalable distributed simulations using technologies such as Docker, Kafka, MongoDB, SSH, and Golang. This repository provides different examples of architectures for efficiently executing distributed simulations.

## 📌 Overview
The repository is structured as follows:

```
ScalableSimulationSystem/
│── examples/
│   │── DockerSimGrid/
│   │── GoMongoQueue/
│── docs/
│── LICENSE
│── README.md
```

Each project within `examples/` implements a specific approach to executing scalable simulations:

### 🛠 DockerSimGrid
- Uses **Docker** for distributed simulation execution.
- Asynchronous communication with **Kafka**.
- Remote management via **SSH**.

### 📌 Conceptual architecture diagram:

![pic1](./docs/DockerSimGrid.png)

📂 **Location:** `examples/DockerSimGrid/`

📜 [Read more about DockerSimGrid](examples/DockerSimGrid/README.md)

---

### 🛠 GoMongoQueue
- Based on **Go** and **MongoDB Change Streams** for simulation queue control.
- Distributed task system without the need for Kafka.
- Simple and efficient for executing database-controlled simulations.

### 📌 Conceptual architecture diagram:

![pic2](./docs/GoMongoQueue.png)

📂 **Location:** `examples/GoMongoQueue/`

📜 [Read more about GoMongoQueue](examples/GoMongoQueue/README.md)

---

## 🚀 How to Use

1. **Clone the repository:**
```sh
git clone https://github.com/JunioCesarFerreira/ScalableSimulationSystem
cd ScalableSimulationSystem
```

2. **Access one of the available examples:**
```sh
cd examples/DockerSimGrid
# or
cd examples/GoMongoQueue
```

3. **Follow the specific instructions in each project for configuration and execution.**

## 📜 License
This project is licensed under the [MIT License](LICENSE).



