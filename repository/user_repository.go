// DBへのアクセスGORMを用いて実装
// コンストラクタは後から作成するGo言語のエントリーポイントのmain.goファイル内で呼び出す

package repository

import (
	"fmt"
	"twitter_echo_backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IFの定義
type UserRepository interface {
	GetAllUsers(users *[]model.User) error
	GetUserById(user *model.User, userId uint) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, userId uint) error
	DeleteUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

// コンストラクタの作成、main.goで呼び出し
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// サインアップ機能 ユーザ追加
func (tr *userRepository) CreateUser(user *model.User) error {
	if err := tr.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザ全取得
func (tr *userRepository) GetAllUsers(users *[]model.User) error {
	if err := tr.db.Order("created_at").Find(users).Error; err != nil {
		return err
	}
	return nil
}

// ユーザ取得
func (tr *userRepository) GetUserById(user *model.User, userId uint) error {
	if err := tr.db.First(user, userId).Error; err != nil {
		return err
	}
	return nil
}

// ユーザ更新
func (tr *userRepository) UpdateUser(user *model.User, userId uint) error {
	result := tr.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Update("email", user.DisplayName)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// ユーザ削除
func (tr *userRepository) DeleteUser(userId uint) error {
	result := tr.db.Where("id=?", userId).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
