package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mammal.net/reccomendationQueue/models"
)

//Handle inserting a reccommendation into a user's queue
func (a application) createRecommendation(res http.ResponseWriter, req *http.Request) {
	// Display create recommendation page on anything other than post
	if req.Method != "POST" {
		a.templates.Execute(res, "add-rec.html")
		return
	}

	// Get information from user
	temp := models.Recomendation{
		ID:     primitive.NewObjectID(),
		Title:  req.FormValue("title"),
		Medium: req.FormValue("medium"),
	}

	// Insert into database
	collection := a.client.Database("queues").Collection("test")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection.InsertOne(ctx, temp)

	// Display confirmation
	fmt.Fprint(res, "Inserted a recomendation into the database: ", temp)
}

//Handle reading reccommendations in a user's queue
func (a application) readRecommendation(res http.ResponseWriter, req *http.Request) {

}

//Handle updating a reccommendation in a user's queue
func (a application) updateRecommendation(res http.ResponseWriter, req *http.Request) {

}

//Handle removing a reccommendation from a user's queue
func (a application) deleteRecommendation(res http.ResponseWriter, req *http.Request) {

}
