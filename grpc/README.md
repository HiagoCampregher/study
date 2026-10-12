# gRPC com Go

Projeto de estudo de uma API gRPC para categorias, usando Go e SQLite.

## Conceitos praticados

- contrato de serviço com Protocol Buffers;
- chamadas unary;
- client streaming;
- bidirectional streaming;
- persistência local com SQLite;
- server reflection;
- testes da camada de dados.

## Executar

Requisitos: Go 1.27 ou superior. O banco e as tabelas são criados automaticamente.

Em um terminal, inicie o servidor:

```bash
make server
```

Em outro, execute o cliente de exemplo:

```bash
make client
```

O endereço padrão é `localhost:50051`. Para usar outra porta:

```bash
go run ./cmd/grpcServer -address=:50052
go run ./cmd/grpcClient -address=localhost:50052
```

Também é possível personalizar os dados enviados:

```bash
go run ./cmd/grpcClient -name="Go" -description="Linguagem de programação"
```

## Testes

```bash
make test
```

## Gerar novamente o código protobuf

Instale `protoc`, `protoc-gen-go` e `protoc-gen-go-grpc`, depois execute:

```bash
make generate
```

O contrato está em [`proto/course_category.proto`](./proto/course_category.proto).
