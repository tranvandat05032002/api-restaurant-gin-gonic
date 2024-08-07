package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectId [12]byte

type RoleType struct {
	Owner    string `json:"owner" bson:"owner"`
	Employee string `json:"employee" bson:"employee"`
}
type AccountModel struct {
	id     primitive.ObjectID `json:"id" bson:"_id"`
	name   string             `json:"name" bson:"name" binding:"required,min=2,max=50"`
	email  string             `json:"email" bson:"email" binding:"required,email"`
	role   RoleType           `json:"role_type" bson:"role_type"`
	avatar string             `json:"avatar" bson:"avatar"`
}
