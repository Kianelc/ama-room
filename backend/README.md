# AMA Public Room Project - Backend

## Descrição do Projeto

Este diretório contém a implementação do backend da aplicação web para salas públicas de AMA (Ask Me Anything). O backend é responsável por gerenciar as salas, lidar com a comunicação em tempo real entre os participantes e gerenciar as votações para priorizar as perguntas mais relevantes.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Comunicação em Tempo Real:** WebSocket
- **Gerenciamento de Concorrência:** Recursos nativos do Go para código concorrente e paralelo

## Estrutura de Pastas

- **`/cmd`**: Contém os comandos principais para iniciar a aplicação.
- **`/internal`**: Código interno do backend, incluindo lógica de negócios e serviços.

## Como Executar

1. **Instale as dependências:**

   ```sh
   go mod tidy
   ```

2. **Execute a API:**

   ```sh
   go run cmd/wsrs/main.go
   ```

3. **Configuração do ambiente:**
   - Variáveis de ambiente necessárias estão definidas no arquivo `.env.example`.
   - Renomeie para `.env` e ajuste conforme necessário.
