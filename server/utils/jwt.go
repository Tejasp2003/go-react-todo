package utils

import (
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

// GenerateJWT generates a new JWT token

func GenerateJWT(userId string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": userId,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })  //jwt.MapClaims is a type alias for map[string]interface{} in the jwt-go library. It is used to represent the claims in a JWT token. Claims 
	//are the statements about an entity (typically, the user) and additional data. The claims are encoded in the JWT token and are used to verify the authenticity of the token. 

	// fmt.Println("Token: ", token)

    tokenString, err := token.SignedString(jwtSecret) // this will sign the token with the secret key and return the token string

	// fmt.Println("TokenString: ", tokenString)

    if err != nil {
        return "", err
    }

    return tokenString, nil //why return nil? Because we are not returning any error
}


// ValidateToken validates a JWT token

func ValidateToken(tokenString string) (string, error) {

	//step 1: Parse the token
	//step 2: Check if the token is valid
	//step 3: Extract the userId from the token
	//step 4: Return the userId

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

	// fmt.Println("Token: ", token)

	claims, ok := token.Claims.(jwt.MapClaims); //claims is a map of claims in the token

    if ok && token.Valid {
        userId := claims["userId"].(string) //extract the userId from the claims map
        return userId, nil
    } else {
        return "", err
    }
}