package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	_Id         primitive.ObjectID `json: "_id"`
	Id          int32              `json: "id"`
	Title       string             `json: "title"`
	Url         string             `json: "url"`
	ImageUrl    string             `json: "imageUrl"`
	NewsSite    string             `json: "newsSite"`
	Summary     string             `json: "summary"`
	PublishedAt string             `json: "publishedAt"`
	UpdatedAt   string             `json: "updatedAt"`
	Featured    bool               `json: "featured"`
	Launches    []Provider         `json: "launches"`
	Events      []Provider         `json: "events"`
}

type Articles struct {
	Articles []Article `json: "articles"`
}

type Provider struct {
	id       string `json: "id"`
	provider string `json: "provider"`
}
