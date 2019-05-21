package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addTask(task string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("todo_cli").Collection("tasks")
	currentTime := time.Now()
	collection.InsertOne(ctx, bson.M{"task": task, "created_on": currentTime.Local()})
}

func main() {
	var i int

	fmt.Println("welcome julio what can jarvis help you with today")
	fmt.Println(`
	  I can help you with your tasks please enter what you would like to do:
		1. Create a new task
		2. View all your tasks
		3. Mark a task as complete
		4. Remove a task`)

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println("Please enter a NUMBER from 1 - 4")
	}

	switch i {
	case 1:
		var task string
		fmt.Println("Please enter your task")
		_, err := fmt.Scanf("%d", &task)
		if err != nil {
			panic(err)
		}
		addTask(task)
		break
	case 2:
		fmt.Print("You have entered 2: ", i)
		break
	case 3:
		fmt.Print("You entered 3: ", i)
		break
	case 4:
		fmt.Print("You entered 4: ", i)
		break
	}
}
