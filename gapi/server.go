package gapi

import (
	"github.com/LastBit97/go-mongodb-api/config"
	"github.com/LastBit97/go-mongodb-api/pb"
	"github.com/LastBit97/go-mongodb-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config         config.Config
	authService    services.AuthService
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcServer(config config.Config, authService services.AuthService,
	userService services.UserService, userCollection *mongo.Collection) (*Server, error) {

	server := &Server{
		config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
	}

	return server, nil
}
