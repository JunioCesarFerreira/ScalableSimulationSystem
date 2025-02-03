package sshhandler

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// ConnectAndExecute conecta ao container via SSH e executa um comando
func ConnectAndExecute(containerID string, command string) error {
	// Configuração do cliente SSH
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("root"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Conecta ao container via SSH
	client, err := ssh.Dial("tcp", "localhost:2223", config)
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
