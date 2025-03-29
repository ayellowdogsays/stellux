package domain

import (
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Posts struct {
	ID          bson.ObjectID `bson:"_id"`
	CreatedAt   time.Time     `bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"`
	DeletedAt   time.Time     `bson:"deleted_at,omitempty"`
	Title       string
	Content     string
	Author      string
	Description string
	Category    string
	Tags        []string
	Cover       string
	IsTop       *bool `bson:"is_top"`
	IsPublish   *bool `bson:"is_publish"`
}

func ToPtr(posts []Posts) []*Posts {
	return lo.Map(posts, func(item Posts, _ int) *Posts {
		return &item
	})
}
