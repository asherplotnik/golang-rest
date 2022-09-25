package repository

import (
	"context"
	"errors"
	"github.com/asherplotnik/golang-rest/model"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const dbName = "goTest"
const colName = "todos"
const connectionString = "mongodb://localhost:27017"

var collection *mongo.Collection

func Init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil{
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Mongodb Connection success.")

}

func InsertOneTodo(todo model.Todo) (model.Todo, error) {
	inserted, err := collection.InsertOne(context.Background(), todo)
	if err != nil{
		return model.Todo{}, model.TodoError{Message: err.Error(), Code: 500}
	}
	fmt.Println("Inserted 1 todo in db with id: ", inserted.InsertedID)
	return todo, nil
} 

func UpdateOneTodo(todo model.Todo) (int64, error){
	filter := bson.M{"_id": todo.ID}
	update := bson.M{"$set": bson.M{"completed": todo.Completed, "item": todo.Item}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, errors.New("error updating record")
	}

	fmt.Println("modified count: ", result.ModifiedCount)
	return result.ModifiedCount, nil
}

func DeleteOneTodo(todoId string) (int, error){
	filter := bson.M{"_id": todoId}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return 0, model.TodoError{Message: "not found", Code: 404}
	}

	fmt.Println("deleted count: ", result.DeletedCount)
	return int(result.DeletedCount), nil
}

func DeleteAllTodo() (int, error) {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return 0, model.TodoError{Message: "error - no records has been deleted", Code: 500}
	}

	fmt.Println("deleted all: ", result.DeletedCount)
	return int(result.DeletedCount), nil
}

func GetAllTodos() ([]model.Todo, error) {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, model.TodoError{Message: "error fetching data", Code: 500} 
	}

	var todos []model.Todo

	for cursor.Next(context.Background()) {
		var todo model.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	defer cursor.Close(context.Background())

	if len(todos) == 0 {
		return nil, model.TodoError{Message: "no content", Code: 404}
	}
	return todos, nil

}

func GetTodoById(todoId string) (model.Todo, error) {
	filter := bson.M{"_id": todoId}
	
	var todo model.Todo
	err := collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		return todo, model.TodoError{Message: "not found", Code: 404} 
	}

	return todo, nil
}
