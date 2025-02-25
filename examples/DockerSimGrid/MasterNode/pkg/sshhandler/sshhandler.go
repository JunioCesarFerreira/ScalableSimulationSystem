package sshhandler

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// ConnectAndExecute conecta ao container via SSH e executa um comando
func ConnectAndExecute(containerName string, port string, command string) error {
	// Configuração do cliente SSH
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Conecta ao container via SSH
	// !# Prod
	client, err := ssh.Dial("tcp", containerName+":22", config)
	// !# Debug
	//client, err := ssh.Dial("tcp", "localhost:"+port, config)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer client.Close()

	// Cria uma sessão SSH
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Executa o comando remoto
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	fmt.Println("SSH command output:", b.String())
	return nil
}
