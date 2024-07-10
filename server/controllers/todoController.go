package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/tejasp2003/go-react-todo/config"
	"github.com/tejasp2003/go-react-todo/models"
	"github.com/tejasp2003/go-react-todo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func errorHandler(err error, c *fiber.Ctx, message string) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": message})
	}
	return nil
}

// CreateTodo creates a new todo for a user
func CreateTodo(c *fiber.Ctx) error {
	// Parse token from the Authorization header
	token := string(c.Request().Header.Peek("Authorization"))

	// Validate token to get user ID
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")

	// Parse the request body to get the new todo
	todo := new(models.Todo)
	errorHandler(c.BodyParser(todo), c, "Cannot parse JSON")

	// Generate a new todo ID
	todo.ID = primitive.NewObjectID().Hex()

	// Convert userId to ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	// Get the users collection
	collection := config.GetCollection(config.DB, "users")

	// Update the user document to add the new todo
	filter := bson.M{"_id": userObjID}
	update := bson.M{"$push": bson.M{"todos": todo}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	errorHandler(err, c, "Cannot create todo")

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// GetTodos retrieves all todos for a user
func GetTodos(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))

	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")
	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	collection := config.GetCollection(config.DB, "users")

	// Find the user document
	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": userObjID}).Decode(&user)
	errorHandler(err, c, "Cannot find user")

	return c.JSON(user.Todos)
}

// GetTodoById retrieves a todo for a user

func GetTodoById(c *fiber.Ctx) error {
	// Parse user ID and todo ID from the URL
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")

	todoId := c.Params("todoId")

	// Convert userId to ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	// Get the users collection
	collection := config.GetCollection(config.DB, "users")

	// Find the user document
	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": userObjID}).Decode(&user)
	errorHandler(err, c, "Cannot find user")

	// Find the todo by ID
	var todo models.Todo
	for _, t := range user.Todos {
		if t.ID == todoId {
			todo = t
			break
		}
	}

	return c.JSON(todo)
}

// UpdateTodo updates a todo for a user
func UpdateTodo(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")
	todoId := c.Params("todoId")
	todo := new(models.Todo)
	errorHandler(c.BodyParser(todo), c, "Cannot parse JSON")

	// Convert userId to ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	// Get the users collection
	collection := config.GetCollection(config.DB, "users")

	// Update the user document to update the todo
	filter := bson.M{"_id": userObjID, "todos._id": todoId}

	if todo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}
	update := bson.M{
		"$set": bson.M{
			"todos.$.title": todo.Title,
		},
	}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	errorHandler(err, c, "Cannot update todo")
	return c.JSON(todo)
}

func ToggleTodoCompletion(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")
	todoId := c.Params("todoId")

	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	collection := config.GetCollection(config.DB, "users")

	// Fetch the user and todo to get the current completion status
	var user models.User
	filter := bson.M{"_id": userObjID, "todos._id": todoId}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	errorHandler(err, c, "Cannot find user")

	var currentTodo *models.Todo
	for _, todo := range user.Todos {
		if todo.ID == todoId {
			currentTodo = &todo
			break
		}
	}

	if currentTodo == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	// Toggle the completion status
	toggleStatus := !currentTodo.Done
	update := bson.M{"$set": bson.M{"todos.$.done": toggleStatus}}

	_, err1 := collection.UpdateOne(context.Background(), filter, update)
	errorHandler(err1, c, "Cannot update todo")

	return c.JSON(fiber.Map{"message": "Todo completion status toggled"})
}

// DeleteTodo deletes a todo for a user
func DeleteTodo(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")

	todoId := c.Params("todoId")

	// Convert userId to ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	// Get the users collection
	collection := config.GetCollection(config.DB, "users")

	// Update the user document to remove the todo
	filter := bson.M{"_id": userObjID}
	update := bson.M{"$pull": bson.M{"todos": bson.M{"_id": todoId}}}
	result , err := collection.UpdateOne(context.Background(), filter, update)
	errorHandler(err, c, "Cannot delete todo")

    if result.ModifiedCount == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
    }

    //send status 200 if the todo is deleted

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo deleted"})

}





func GetUserFromToken(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.ValidateToken(token)
	errorHandler(err, c, "Unauthorized")

	userObjID, err := primitive.ObjectIDFromHex(userId)
	errorHandler(err, c, "Invalid user ID")

	collection := config.GetCollection(config.DB, "users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": userObjID}).Decode(&user)
	errorHandler(err, c, "Cannot find user")

	return c.JSON(user)
}