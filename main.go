// Echoフレームワークを使用してAPIサーバーを起動するためのエントリーポイント
// DB接続やバリデーション、データの永続化など、アプリケーションの各層を初期化
// それぞれの層で作成されたコンストラクタを呼び出し、依存関係を解決

package main

import (
	"twitter_echo_backend/controller"
	"twitter_echo_backend/db"
	"twitter_echo_backend/repository"
	"twitter_echo_backend/router"
	"twitter_echo_backend/usecase"
	"twitter_echo_backend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)                   // repositoryで作ったコンストラクタを起動
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator) // usecaseで作ったコンストラクタを起動
	userController := controller.NewUserController(userUsecase)          // controllerで作ったコンストラクタを起動
	e := router.NewRouter(userController)                                // routerで作ったコンストラクタを起動
	e.Logger.Fatal(e.Start(":1323"))                                     // e.Startでサーバ起動、エラーの場合、ログ情報を出力し強制終了
}
