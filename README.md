# Atividade Ponderada 4

Este projeto consiste em uma aplicação que envolve múltiplos serviços, incluindo backend em Python e Golang, coleta e visualização de logs com Elasticsearch, Kibana e Filebeat, e um gateway para gerenciamento de tráfego.

## Estrutura do Projeto

A estrutura do projeto é composta pelos seguintes serviços:

1. [backend-python](./backend-python/): Aplicação Python que fornece um CRUD para produtos e usuários.
2. [backend-golang](./backend-golang/): Aplicação Golang que fornece uma API com CRUD e documentação Swagger.
3. [filebeat](./filebeat/): Serviço que coleta e envia logs para o Elasticsearch.
4. [elasticsearch](./elasticsearch/): Serviço de busca e análise de logs.
5. [kibana](./kibana/): Interface gráfica para visualizar os logs coletados pelo Elasticsearch.
6. [gateway](./gateway/): Serviço de gateway que gerencia o tráfego entre os serviços backend.

> [!NOTE]
> Todos os serviços (gateway, backend-python e backend-golang) estão configurados para enviarem logs para os arquivos que podem ser encontrados na pasta [logs](./logs/).

### [Backend-python](./backend-python/)

Esse é um serviço em Python que fornece um CRUD para produtos e usuários com um sqlite. Ele é executado na porta 8000. Para acessar a documentação da API, acesse a rota `/fastapi/docs`.

Rotas disponíveis:

### [Backend-golang](./backend-golang/)

Esse é um serviço em Golang que fornece um CRUD usuários com um sqlite. Ele é executado na porta 8080. Para acessar a documentação da API, acesse a rota `gin-gonic/swagger/index.html`.

Rotas disponíveis:

### [Gateway](./gateway/)

Foi introduzido a esse sistema um gateway com NGINX que gerencia o tráfego entre os serviços backend.

### [Filebeat](./filebeat/)

Coleta logs dos serviços backend-python e backend-golang e envia para o Elasticsearch para análise e visualização no Kibana.

### [Elasticsearch](./elasticsearch/)

Armazena e indexa os logs coletados pelo Filebeat. Fornece funcionalidades de busca e análise de dados.

### [Kibana](./kibana/)

Interface gráfica para visualização e análise dos logs armazenados no Elasticsearch.

## Funcionamento

## Como executar

Para executar o projeto, siga os passos abaixo:

1. Clone o repositório:

```bash
git clone https://github.com/Lemos1347/inteli-modulo-10-ponderada-4.git
```

2. Acesse a pasta do projeto:

```bash
cd inteli-modulo-10-ponderada-4
```

3. Execute o comando abaixo para subir os serviços:

```bash
docker-compose up
```

Pronto! Agora você pode acessar os serviços nos seguintes endereços:

- Backend Python: http://localhost:3000/fastapi
- Backend Golang: http://localhost:3000/gin-gonic
- Kibana: http://localhost:5601

> [!NOTE]
> Para acessar a documentação da API do backend-python, acesse a rota `http://localhost:3000/fastapi/docs`.
> Para acessar a documentação da API do backend-golang, acesse a rota `http://localhost:3000/gin-gonic/swagger/index.html`.

> [!NOTE]
> Nesse projeto tem também um [Justfile](./Justfile), ele é uma alternativa ao Makefile, caso você queira saber mais sobre isso: [https://github.com/casey/just](https://github.com/casey/just). Se você já conher e já estiver instalado, para subir os serviços basta executar o comando `just app` 😁.
