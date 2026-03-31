# Desafio Cloud Run

Este é um projeto de desafio que implementa um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin).

## Objetivo

Desenvolver um sistema em Go que receba um CEP, identifique a cidade correspondente e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O requisito final é que este sistema esteja publicado e acessível no Google Cloud Run.

## Requisitos Funcionais

-   **Entrada:** O sistema deve receber um CEP válido de 8 dígitos.
-   **Identificação de Localização:** O sistema deve realizar a busca do CEP para encontrar o nome da localização (cidade).
-   **Consulta de Clima:** A partir da localização, o sistema deve consultar a temperatura atual.
-   **Conversão:** O sistema deve retornar as temperaturas formatadas em Celsius, Fahrenheit e Kelvin.

## Especificações da API (Contrato)

### Cenário 1: Sucesso

-   **Código HTTP:** 200
-   **Response Body:**
    ```json
    { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }
    ```

### Cenário 2: Falha (Formato Inválido)

-   **Condição:** Caso o CEP não tenha 8 dígitos ou contenha caracteres inválidos.
-   **Código HTTP:** 422
-   **Mensagem:** `invalid zipcode`

### Cenário 3: Falha (CEP não encontrado)

-   **Condição:** Caso o CEP tenha o formato correto, mas não seja encontrado na base de dados (ex: CEP inexistente).
-   **Código HTTP:** 404
-   **Mensagem:** `can not find zipcode`

## Fórmulas de Conversão

-   **Celsius para Fahrenheit:** `F = C * 1.8 + 32`
-   **Celsius para Kelvin:** `K = C + 273`

## APIs Externas Utilizadas

-   **Localização:** [ViaCEP](https://viacep.com.br/)
-   **Temperatura:** [WeatherAPI](https://www.weatherapi.com/)

## Requisitos de Infraestrutura e Deploy

-   **Docker:** O projeto possui um `Dockerfile` para containerização.
-   **Cloud Run:** A aplicação deve ser implantada no Google Cloud Run.
-   **Testes:** Foram implementados testes automatizados.

## Entregável

-   **Código Fonte:** O código fonte está neste repositório.
-   **URL de Acesso:** O deploy no Google Cloud Run não foi efetuado por problemas na conta.
-   **Testes:** O projeto contém testes automatizados.

## Como rodar os testes

Para rodar os testes, execute o seguinte comando na raiz do projeto:

```
go test ./...
```

## Como rodar a aplicação localmente via Docker

1.  **Construa a imagem Docker:**

    ```
    docker build -t desafio-cloud-run .
    ```

2.  **Rode o container:**

    ```
    docker run -p 8080:8080 desafio-cloud-run
    ```

3.  **Acesse a aplicação:**

    Abra o seu navegador e acesse `http://localhost:8080/weather?cep=SEU_CEP`.
