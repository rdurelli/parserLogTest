



<!-- PROJECT LOGO -->
<br />
<p align="center">

<h3 align="center">Logger Parser</h3>

  <br align="center">
    Parser created with ANTLR 4 to import logs in Go using goroutines. 
    We also provide an end-point (GET) to fetch the <i>n's</i> saved logs. 
    <br />
    <br />
    <br />
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">Sobre o Projeto</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
<!-- ## Sobre o Projeto

Uma das melhores formas de se aprender alguma coisa é fazendo.

Dessa forma, eu resolvi realizar um desafio que de acordo com esse canal (maiores descrições sobre o problema pode ser encontrada em: https://www.youtube.com/watch?v=u3GAYK8hvTI) é um desafio para uma vaga Senior em Golang.

O desafio é o seguinte:
Criar um micro service em Golang que leia um arquivo de log e realize um PARSER de cada linha. E caso essa linha seja do método GET ou POST, deve-se armazenar esse log em uma tabela do MYSQL. Além disso,  a API deverá retornar as X entradas de log salvas filtrando pelo método GET ou POST.

Para criar o PARSER eu utilizei o ANTLR — onde foi possível criar uma gramática para reconhecer as linhas do arquivo de log.. filtrando os métodos GET e/ou POST.

Para maximizar e fazer o processo rápido, foi utilizado goroutines no código. Assim, é possível ler um arquivo de log extremamente grande em pouco tempo.

Boas práticas de programação foram também utilizadas - injeção de dependência (FX), .env variables (Viper), logger (Zap) e como Web Service escolhe-se utilizar o Gin

O código-fonte desse projeto pode ser visualizado em: https://github.com/rdurelli/parserLogTest

 -->

### Built With

* [ANTLR](https://www.antlr.org/)
* [GIN](https://github.com/gin-gonic/gin)
* [Viper](https://github.com/spf13/viper) 
* [Fx](https://github.com/uber-go/fx)
* [Mysql Go Driver](https://github.com/go-sql-driver/mysql)
* [Zap](https://github.com/uber-go/zap)
* [Docker](https://hub.docker.com/)



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps. 

### Prerequisites

* Firstly install Go and Docker in your machine then install the following:

### Installation

* go mod
  ```sh
  go get go.uber.org/fx@v1
  ```
  ```sh
  go get github.com/antlr/antlr4/runtime/Go/antlr
  ```
  ```sh
  go get -u github.com/gin-gonic/gin
  ```

  ```sh
  go get github.com/spf13/viper
  ```

  ```sh
  go get -u github.com/go-sql-driver/mysql
  ```

  ```sh
  go get -u go.uber.org/zap
  ```
  

## Usage

1. Clone the repo
   ```sh
   git clone https://github.com/rdurelli/parserLogTest.git
   ```
2. Install all dependencies listed before
3. Create a .env file with the following:
```sh
   LOG_PATH="...path/logFile.log"
   SERVER_PORT=3333
   
   DB_USER=user
   DB_PASS=pass
   DB_HOST=127.0.0.1
   DB_PORT=port
   DB_NAME=schema_name
   
   ENV=development
   
   LOG_OUTPUT=./server.log
   ```


<!-- USAGE EXAMPLES -->
## Usage

To execute the source code use the Makefile... run directly the main.go file.
if everything was good we will see a message: 
```
Time taken to process all logs
```

Then, an end-point is available: 
```
GET   /api/logs
QUERY PARAMS: offset and limit 
```

One can specify two query params in this end-point: offset and limit



<!-- CONTACT -->
## Contact

Rafael S. Durelli - rafael.durelli@ufla.br

Project Link: [https://github.com/rdurelli/parserLogTest.git](https://github.com/rdurelli/parserLogTest.git)
