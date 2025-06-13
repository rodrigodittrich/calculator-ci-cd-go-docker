# calculator-ci-cd-go-docker
Exemplo de implementação das operações matemáticas utilizando Go Lang, com integração contínua utilizando Docker

### Passo 1 - Criar as operações
- Criação das operações de adição, subtração, multiplicação e divisão;
- Inicialização do modulo go: `go mod init calculator`;
- Baixar as dependências: `go mod tidy`;
- Executar as operações: `go run cmd/main.go`;

Agora vou pro escritório e em outro momento forá de hora darei seguimento no passo 2.

### Passo 2 - Criar teste de unidade das operações
- Criação dos testes de unidade para as operações de adição, subtração, multiplicação e divisão;
- Executar os testes: `go test ./...`;