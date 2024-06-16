# Task API versão GRPC

  

## Descrição

  

Essa aplicação se trata de um backend para registro de atividades em um banco de dados MySQL.

  

## Complição do arquivo .proto e execução das chamadas

  

Para compilar os arquivos devemos executar o seguinte comando na pasta contendo os arquivos .proto

  

```
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative crud.proto
```
Para a realização das chamadas gRPC, o arquivo .proto está disponível na pasta raiz do projeto


## Observações

Desejo pontuar algumas coisas com relação à esse projeto

 - A cobertura dos testes unitários está longe do ideal, meu objetivo nesse projeto era apenas testar o funcionamento do protocolo, migrando uma aplicação que antes era http para grpc
  
 - Na pasta docker-compose se encontra o arquivo para subir todos os containers necessários para a execução dessa aplicação.
