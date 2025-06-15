# Operações matemáticas utilizando GO, CI-CD com Dcoker
Exemplo de implementação das operações matemáticas utilizando Go Lang, com integração contínua utilizando Docker

## Processos de desenvolvimento

### Objetivo geral
- Utilização de actions do github para realizar CI/CD e Dockerfile para criação de imagem e container;

### Objetivos especificos
- Trabalhar com um branch develop protegido;
- Subir alterações em branchs específicos e fazer pull request e merge para o branch develop;
- Utilização de Docker para criação de imagem e container;
- Utilização de Github Actions para CI/CD;

### Passo 1 - Criar as operações
- Criação das operações de adição, subtração, multiplicação e divisão;
- Inicialização do modulo go: `go mod init calculator`;
- Baixar as dependências: `go mod tidy`;
- Executar as operações: `go run cmd/main.go`;

### Passo 2 - Criar teste de unidade das operações
- Criação dos testes de unidade para as operações de adição, subtração, multiplicação e divisão;
- Executar os testes: `go test ./...`;

### Passo 3 - Criar o arquivo .github/workflows/ci.yaml
- Criação do arquivo ".github/workflows/ci.yaml" para realizar a CI/CD;
- Neste passo, somente será executado o teste de unidade das operações no evento push do branch main;
```yaml
name: ci-cd-go-workflow
on: [push]
jobs:
  check-application:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: '1.19'
      - run: go test ./...
```