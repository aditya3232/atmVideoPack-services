package middleware

// func AuthMiddleware(usersService users_model.Service) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get Authorization header from request
// 		authHeader := c.GetHeader("Authorization")

// 		// Get user ID from JWT token
// 		userID, err := jwt.GetUserIDFromToken(authHeader)
// 		if err != nil {
// 			// Return 401 Unauthorized if JWT token is invalid or missing
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		// var
// 		var UsersGetOneByIdInput users_model.UsersGetOneByIdInput{}

// 		user, err := usersService.GetOne(UsersGetOneByIdInput)
// 		if err != nil {
// 			// Return 401 Unauthorized if user doesn't exist in database
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		// Check if user.Token is different from token headers
// 		if user.Token != authHeader {
// 			// Return 401 Unauthorized if user.Token is different from token headers
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		c.Set("currentUser", user)

// 		// Call next middleware/handler function
// 		c.Next()
// 	}
// }
