package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"m/pkg/dockerclient"
	"m/pkg/kafkaclient"
	"m/pkg/sshhandler"
)

type Task struct {
	Id   int    `json:"task_id"`
	Data string `json:"data"`
}

func main() {
	log.Println("starting...")
	// Contexto para controle de operações
	ctx := context.Background()

	log.Println("NewDockerClient...")
	// Inicializa o cliente Docker
	dockerClient, err := dockerclient.NewDockerClient()
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	log.Println("NewKafkaConsumer...")
	// Inicializa o consumidor Kafka
	kafkaClient, err := kafkaclient.NewKafkaConsumer("localhost:29092", "simulation_tasks")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaClient.Close()

	scanner := bufio.NewScanner(os.Stdin)

	var task Task
	// Consome mensagens do Kafka
	for {
		msg, err := kafkaClient.ConsumeMessage(ctx)
		if err != nil {
			log.Printf("Error consuming message: %v", err)
			continue
		}

		log.Printf("Received task: %s", string(msg.Value))

		json.Unmarshal(msg.Value, &task)

		log.Printf("Check JSON %v", task)

		containerName := "sim_" + strconv.Itoa(task.Id)
		hostPort := strconv.Itoa(2222 + task.Id)

		log.Printf("Starting container %s in port %s\n", containerName, hostPort)

		// Cria e inicia o container
		containerID, err := dockerclient.CreateContainer(ctx, dockerClient, containerName, hostPort)
		if err != nil {
			log.Printf("Failed to create container: %v", err)
			continue
		}
		log.Printf("Container created with ID: %s", containerID)

		fmt.Print("Continue...")
		scanner.Scan()

		err = dockerclient.StartContainer(ctx, dockerClient, containerID)
		if err != nil {
			log.Printf("Failed to start container: %v", err)
			continue
		}
		log.Printf("Container started with ID: %s", containerID)

		running, err := dockerclient.IsContainerRunning(ctx, dockerClient, containerID)
		if err != nil {
			log.Printf("Container is not running: %v", err)
			continue
		}
		if !running {
			log.Printf("Container is not running!")
			continue
		}

		fmt.Print("Continue...")
		scanner.Scan()

		// Conecta ao container via SSH e executa um comando
		err = sshhandler.ConnectAndExecute(containerID, "echo 'SSH conectado com sucesso ao container!'")
		if err != nil {
			log.Printf("Failed to connect via SSH: %v", err)
			continue
		}

		// Simula o processamento da tarefa
		time.Sleep(10 * time.Second)

		// Remove o container após o processamento
		err = dockerclient.RemoveContainer(ctx, dockerClient, containerID)
		if err != nil {
			log.Printf("Failed to remove container: %v", err)
		}
		log.Printf("Container removed: %s", containerID)
	}
}
