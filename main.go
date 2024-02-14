package main

import (
	"fmt"
	"net/http"
	"time"

	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/satyalohit/mongo-golang/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	//r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		panic(err)
	}else{
		fmt.Println("Connected to Database")
	}
/* 	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}() */
	return client
}
