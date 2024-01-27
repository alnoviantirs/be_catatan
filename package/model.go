package _package

import "go.mongodb.org/mongo-driver/bson/primitive"

type Catatan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID2         int                `bson:"id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Note        string             `bson:"note,omitempty" json:"note,omitempty"`
	Date        string             `bson:"date,omitempty" json:"date,omitempty"`
	StartTime   string             `bson:"startTime,omitempty" json:"startTime,omitempty"`
	EndTime     string             `bson:"endTime,omitempty" json:"endTime,omitempty"`
	Remind      *int               `bson:"remind,omitempty" json:"remind,omitempty"`
	Repeat      string             `bson:"repeat,omitempty" json:"repeat,omitempty"`
	IsCompleted *int               `bson:"isCompleted,omitempty" json:"isCompleted,omitempty"`
	CompletedAt string             `bson:"completedAt,omitempty" json:"completedAt,omitempty"`
	CreatedAt   string             `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   string             `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Color       *int               `bson:"color,omitempty" json:"color,omitempty"`
	User				string          	 `bson:"user,omitempty" json:"user,omitempty"`
}
