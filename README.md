# excercicios_go_lang
Repositório contendo alguns exemplos de exercicios com </br>
GOLang e algumas restAPI </br>

<h3>importanto as dependencias do GIN (rodar no terminal)</h3>

```
go get -u github.com/gin-gonic/gin
```
aqui o repositório do GIN: </br>
https://github.com/gin-gonic/gin   </br>
Nesse Link o endereço da primeira api. Pode ser acessada pelo navegador.
http://localhost:8080/ping

<h3>imports do mongo (rodar no terminal)</h3>

```

go get go.mongodb.org/mongo-driver
go get go.mongodb.org/mongo-driver/x/mongo/driver@v1.12.0
go get go.mongodb.org/mongo-driver/x/mongo/driver/topology@v1.12.0
go get go.mongodb.org/mongo-driver/x/mongo/driver/auth@v1.12.0
go get go.mongodb.org/mongo-driver/mongo/options@v1.12.0
go get go.mongodb.org/mongo-driver/x/mongo/driver/ocsp@v1.12.0
go get go.mongodb.org/mongo-driver/internal/aws/credentials@v1.12.0

```

api3_crud_mongo, o crud funciona.</br>
Para esse projeto rodar vc precisa ter o mongoDB rodando no seu PC </br>

Aqui dois curls para uso:</br>
POST</br>

```

curl --location --request POST 'http://localhost:3031/pessoas' \
--header 'Authorization: 123' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nome":"teste",
    "endereco":[{
        "cidade":"EV",
        "rua":"abc",
        "numero":123
    }]
}'

```

GET</br>

```

curl --location --request GET 'http://localhost:3031/pessoas' \
--header 'Authorization: 123'

```
