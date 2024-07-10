package controllers

import (
	"context"


	"github.com/gofiber/fiber/v2"
	"github.com/tejasp2003/go-react-todo/config"
	"github.com/tejasp2003/go-react-todo/models"
	"github.com/tejasp2003/go-react-todo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	//step 1: get the collection
	//step 2: parse the request body
	//step 3: hash the password
	//step 4: insert the user into the database
	//step 5: generate a JWT token
	//step 6: return the token
	
    collection :=config.GetCollection(config.DB, "users") // Getting the users collection from the database

    user := new(models.User) // Creating a new user model

	

	err := c.BodyParser(user);  // Parsing the request body into the user model



    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

	// check if the user already exists
	filter := bson.D{{"email", user.Email}} // Creating a filter to check if the user already exists
	count, _ := collection.CountDocuments(context.Background(), filter) // Counting the number of documents that match the filter
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}


    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14) // Hashing the password
    user.Password = string(hashedPassword) // Setting the password to the hashed password

	
    result,  err1 := collection.InsertOne(context.Background(), user) // Inserting the user into the database



    if err1 != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot register user"})
    }


	// &{ObjectID("668c257a8b9d982f7603e087")}

	// fmt.Println("InsertedId", result.InsertedID)
	userId := result.InsertedID.(primitive.ObjectID).Hex() // Getting the ID of the inserted user

	// fmt.Println(userId)
	 
    token, err := utils.GenerateJWT(userId) // Generating a JWT token
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}

func Login(c *fiber.Ctx) error {
    collection := config.GetCollection(config.DB, "users")

    data := new(models.User)
	err := c.BodyParser(data);
    if  err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    user := new(models.User)
    filter := bson.D{{"email", data.Email}}
    err1 := collection.FindOne(context.Background(), filter).Decode(&user) // Finding the user by email
    if err1 == mongo.ErrNoDocuments {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}



