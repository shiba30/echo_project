// Echoフレームワークを使用してHTTPルーティングを設定するための関数

package router

import (
	"os"
	"twitter_echo_backend/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(uc controller.UserController) *echo.Echo {
	e := echo.New() // echoのインスタンス作成
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("SERVER_ADDRESS")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// UserのAPIエンドポイント
	user := e.Group("/users")
	user.POST("", uc.CreateUser)
	user.GET("", uc.GetAllUsers)
	user.GET("/:userId", uc.GetUserById)
	user.PUT("/:userId", uc.UpdateUser)
	user.DELETE("/:userId", uc.DeleteUser)

	/*
		// TweetのAPIエンドポイント
		tweet := e.Group("/tweets")
		tweet.GET("", tc.GetAllTweets)
		tweet.GET("/:tweetId", tc.GetTweetById)
		tweet.POST("", tc.CreateTweet)
		tweet.PUT("/:tweetId", tc.UpdateTweet)
		tweet.DELETE("/:tweetId", tc.DeleteTweet)
	*/

	return e
}
