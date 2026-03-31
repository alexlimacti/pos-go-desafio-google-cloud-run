# Desafio Cloud Run

Este é um projeto de desafio que implementa um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin).

## URL do sistema no Cloud Run

Ainda não foi feito o deploy.

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
