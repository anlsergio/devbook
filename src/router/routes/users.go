package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequiresAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetUser,
		RequiresAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequiresAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequiresAuth: true,
	},
	{
		URI: "/users/{userID}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:    "/users/{userID}/unfollow",
		Method: http.MethodDelete,
		Function: controllers.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:    "/users/{userID}/followers",
		Method: http.MethodGet,
		Function: controllers.GetFollowers,
		RequiresAuth: true,
	},
	{
		URI:    "/users/{userID}/following",
		Method: http.MethodGet,
		Function: controllers.GetFollowing,
		RequiresAuth: true,
	},
	{
		URI:    "/users/{userID}/update-password",
		Method: http.MethodPost,
		Function: controllers.UpdatePassword,
		RequiresAuth: true,
	},
	{
		URI:    "/users/{userID}/posts",
		Method: http.MethodGet,
		Function: controllers.GetUserPosts,
		RequiresAuth: true,
	},
}
