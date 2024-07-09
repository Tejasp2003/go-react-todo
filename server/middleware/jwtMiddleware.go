package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/tejasp2003/go-react-todo/utils"
)

func JWTMiddleware(c *fiber.Ctx) error {


	//step 1: Get the token from the Authorization header in the request
	//step 2: Check if the token is empty
	//step 3: Validate the token
	//step 4: Extract the userId from the token
	//step 5: Set the userId in the Locals
	//step 6: Return the control to the next handler

    tokenString := c.Get("Authorization") //get the token from the Authorization header in the request
    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token"})
    }

    userId, err := utils.ValidateToken(tokenString)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    c.Locals("userId", userId) // Locals is a map that can be used to store values that are specific to the current request.
    return c.Next()
}
