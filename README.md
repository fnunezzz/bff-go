[![Run tests](https://github.com/fnunezzz/bff-go/actions/workflows/tests.yml/badge.svg)](https://github.com/fnunezzz/bff-go/actions/workflows/tests.yml)

# BFF Go

The code base is in PT-BR as the POC was made with the intention of being read by developers with various levels of english proficiency but full Portuguese fluency.

-   [PT-BR Readme](https://github.com/fnunezzz/bff-go/blob/main/docs/PT-BR.md)

This is a POC (Proof Of Concept) of a BFF (Backend For Frontend) in GO. It implements the "DDD" concept and folders structure, following both visually and functionally it's concepts. Because it's a BFF, it does not need Aggregates and Entities, only services and interfaces.

The `go` language has some recommended structures specific to the language itself, such as the `cmd` and `internal` folders.

-   The `cmd` folder is where the application's input files are grouped (equivalent to `main.js` in NestJS/Node.js).
-   The `internal` folder has some language specific uses but generally it is equivalent to the `src` folder.

```text
cmd/
├─ bff/
│  ├─ main.go
internal/
├─ controller/
├─ domain/
│  ├─ service/
├─ infra/
│  ├─ service/
├─ shared/
│  ├─ middleware/
```

At the moment, the implementation of the application boils down to just the `ProductService | ProdutoService` interface implementing the `FetchProductData | BuscarDadosProduto` method which returns a `FetchProductDataResponse | BuscarDadosProdutoResponse` response merging and unifying the `media` and `price` outbound request results.

```json
{
    "data": [
        {
            "sku": "1515",
            "preco": 10.0,
            "imagens": [
                {
                    "url": "http://teste.com.br",
                    "principal": true,
                    "ordem": 1
                }
            ]
        }
    ]
}
```

## How to use

It was done using version `1.21.3`

As the application is only a POC at the moment, the functionality is being tested in the `product_test | produto_test` file. It tests the implementation and grouping of data from the `media` and `price` services in various use cases.

To run:

```bash

# Running test cases
$ make test

```

![drawing](docs/img/image.png)
