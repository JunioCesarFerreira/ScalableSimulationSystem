# Usar a imagem base do Ubuntu
FROM ubuntu:latest

# Instalar o OpenSSH Server
RUN apt-get update && \
    apt-get install -y openssh-server && \
    mkdir /var/run/sshd

# Definir a senha do usuário root (você pode mudar conforme necessário)
RUN echo 'root:password' | chpasswd

# Permitir conexões SSH com senha
RUN sed -i 's/^#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# Expor a porta 22
EXPOSE 22

# Iniciar o serviço SSH
CMD ["/usr/sbin/sshd", "-D"]
