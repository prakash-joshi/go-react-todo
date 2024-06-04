package main

import (
	"context"
	"fmt"
	"go-react-todo/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongodb")

	collection = client.Database("golang-db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", GetAllTodo)
	app.Post("/api/todo", CreateTodo)
	app.Patch("/api/todo/:id", UpdateTodo)
	app.Delete("/api/todo/:id", DeleteTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	log.Fatal(app.Listen(":" + PORT))

}

func GetAllTodo(c *fiber.Ctx) error {
	var todos []models.Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.Status(200).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo id"})
	}

	filteredData := bson.M{"_id": objectID}
	updateData := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filteredData, updateData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	deleteData := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), deleteData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
