package authtoken

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	UserId primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	jwt.StandardClaims
}
