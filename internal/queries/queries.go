package queries

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB_Client struct {
	Client *mongo.Client
}

// post object which will be saved in db
type posts struct {
	ID     primitive.ObjectID `bson:"_id"`
	PostID int                `bson:"id"`
}

// search post id, to prevent duplicated posts
func (m *MongoDB_Client) PostExists(id int32) bool {
	var post posts

	err := m.Client.Database("main").Collection("posts").FindOne(context.TODO(), bson.M{"id": id}).Decode(&post)

	if err != nil {
		return false
	}

	if post.PostID == 0 {
		return false
	}
	return true
}

func (m *MongoDB_Client) Insert(id int64) error {
	_, err := m.Client.Database("main").Collection("posts").InsertOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
