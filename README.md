# mad-profile-service
MAD Service for user profiles
 Based on Go arch - https://github.com/restuwahyu13/go-rest-api/blob/main/middlewares/authJwt.go
## JWT Token not correct 

```
func verifyJWT(JWTtoken string, HMACsecret []byte) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(JWTtoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return HMACsecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("Token wasn't verified correctly")
}

``` 

## JWT Token correct 

```
func verifyJWT(JWTtoken string, HMACsecret []byte) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(JWTtoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return HMACsecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token wasn't verified correctly")
}

```