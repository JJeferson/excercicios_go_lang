package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Address struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Cidade string `json:"cidade"`
}

type Person struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	Nome           string    `json:"nome"`
	DataNascimento string    `json:"dt_nascimento"`
	Endereco       []Address `json:"endereco"`
}

var collection *mongo.Collection

func main() {
	// Conexão com o MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Selecionar a coleção
	collection = client.Database("crud_1_go_lang").Collection("pessoa")

	router := gin.Default()

	// Middleware para verificar o header de autenticação
	router.Use(authMiddleware)

	// Rotas
	router.POST("/pessoas", createPerson)
	router.GET("/pessoas", getAllPeople)
	router.GET("/pessoas/:id", getPersonByID)
	router.GET("/pessoas/nome/:nome", getPeopleByName)
	router.PUT("/pessoas/:id", updatePerson)
	router.DELETE("/pessoas/:id", deletePerson)

	// Iniciar o servidor
	router.Run(":3031")
}

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

func createPerson(c *gin.Context) {
	var newPerson Person

	// Bind do JSON da requisição para a struct Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Inserir a nova pessoa no MongoDB
	result, err := collection.InsertOne(context.Background(), newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
}

func getAllPeople(c *gin.Context) {
	var people []Person

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		people = append(people, person)
	}

	c.JSON(http.StatusOK, people)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	var person Person
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func getPeopleByName(c *gin.Context) {
	nome := c.Param("nome")

	var people []Person
	cursor, err := collection.Find(context.Background(), bson.M{"nome": nome})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		people = append(people, person)
	}

	c.JSON(http.StatusOK, people)
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")

	var updatedPerson Person
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": updatedPerson})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
