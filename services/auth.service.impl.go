package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/LastBit97/go-mongodb-api/models"
	"github.com/LastBit97/go-mongodb-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}

func (as *AuthServiceImpl) SignUpUser(user *models.SignUpInput) (*models.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	res, err := as.collection.InsertOne(as.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := as.collection.Indexes().CreateOne(as.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = as.collection.FindOne(as.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (as *AuthServiceImpl) VerifyEmail(verificationCode string) error {
	code := utils.Encode(verificationCode)

	query := bson.D{{Key: "verificationCode", Value: code}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verified", Value: true}}}, {Key: "$unset", Value: bson.D{{Key: "verificationCode", Value: ""}}}}
	result, err := as.collection.UpdateOne(as.ctx, query, update)
	if err != nil {
		fmt.Print(err)
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("could not verify email address")
	}
	fmt.Println(result)
	return nil
}

func (as *AuthServiceImpl) UpdatePasswordResetTokenByEmail(email string, passwordResetToken string) error {
	// Update User in Database
	query := bson.D{{Key: "email", Value: strings.ToLower(email)}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "passwordResetToken", Value: passwordResetToken}, {Key: "passwordResetAt", Value: time.Now().Add(time.Minute * 15)}}}}
	result, err := as.collection.UpdateOne(as.ctx, query, update)

	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("there was an error sending email")
	}
	return nil
}

func (as *AuthServiceImpl) ResetPassword(resetToken string, password string) error {
	hashedPassword, _ := utils.HashPassword(password)

	passwordResetToken := utils.Encode(resetToken)

	// Update User in Database
	query := bson.D{{Key: "passwordResetToken", Value: passwordResetToken}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashedPassword}}}, {Key: "$unset", Value: bson.D{{Key: "passwordResetToken", Value: ""}, {Key: "passwordResetAt", Value: ""}}}}
	result, err := as.collection.UpdateOne(as.ctx, query, update)

	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("token is invalid or has expired")
	}
	return nil
}
