package like

import (
	"context"
	"github.com/Peterliang233/go-blog/databases/mysql"
	RedisDatabase "github.com/Peterliang233/go-blog/databases/redis"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// CheckLike 根据邮箱验证是否点过赞,true表示点过赞，false表示没有点过赞
func CheckLike(email string) (bool, error) {
	// 检查是否在redis里面
	if ok := CheckRedisEmail(email); ok {
		return true, nil
	}
	// 检查mysql数据库里面是否有这个邮箱
	ok, err := CheckMysqlEmail(email)

	if err != nil {
		return false, err
	}

	return ok, nil
}

// CheckRedisEmail 检查redis里面的email集合里面是否有这个邮箱
func CheckRedisEmail(email string) bool {
	return RedisDatabase.
		RedisClient.
		SIsMember(context.Background(), "email", email).
		Val()
}

// CheckMysqlEmail 检查在数据库里面是否有这个邮箱
func CheckMysqlEmail(email string) (bool, error) {
	var e model.Email
	if err := mysql.Db.Where("email = ?", email).First(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// CreateEmail 持久化点赞的邮箱数据
func CreateEmail(email string) (int, error) {
	e := &model.Email{
		Email: email,
	}

	if err := mysql.Db.Create(&e).Error; err != nil {
		return errmsg.ErrLikeCreate, err
	}

	return errmsg.Success, nil
}

// RedisCreateEmail 将email放进redis里面
func RedisCreateEmail(email string) error {
	return RedisDatabase.
		RedisClient.
		SAdd(context.Background(), "email", email).
		Err()
}
