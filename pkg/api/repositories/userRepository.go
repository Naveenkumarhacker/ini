package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ini/configs/database"
	"ini/pkg/api/models"
	"ini/pkg/utils"
	"time"
)

var UserRepo UserRepositoryI = new(userRepository)

type UserRepositoryI interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
	GetUserByUsernameAndPassword(username, password string) (models.User, error)
}

type userRepository struct {
}

var userCollection *mongo.Collection = database.GetCollection(database.DB, "user")

func (u *userRepository) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return nil, err
		}
		users = append(users, singleUser)
	}
	return users, nil
}

func (u userRepository) GetUser(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u userRepository) CreateUser(user models.User) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if _, err := userCollection.InsertOne(ctx, &user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u userRepository) UpdateUser(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(user.Id.String())

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": update})
	if err != nil {
		return models.User{}, err
	}
	if result.MatchedCount != 0 && result.ModifiedCount != 0 {
		return models.User{}, utils.UserNotFoundErr
	}
	var updatedUser models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": user.Id}).Decode(&updatedUser)
	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}

func (u userRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("User with specified ID not found!")
	}
	return nil
}

func (u userRepository) GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"name": username, "password": password}).Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}
