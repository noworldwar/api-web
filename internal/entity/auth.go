package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/noworldwar/api-web/internal/model"
)

func SetAuth(token string, in *model.Player) error {
	m := &model.Auth{
		PlayerID: in.PlayerID,
		Email:    in.Email,
		NickName: in.NickName,
		Balance:  in.Balance,
		Token:    token,
	}

	out, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return model.RedisDB.Set(token, string(out), 1*time.Hour).Err()
}

func ReSetAuth(m *model.Auth) error {
	out, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return model.RedisDB.Set(m.Token, string(out), 1*time.Hour).Err()
}

func GetAuth(token string) (string, error) {
	val, err := model.RedisDB.Get(token).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func ToAuth(val interface{}) *model.Auth {

	if val == nil {
		return nil
	}

	val_str := fmt.Sprintf("%s", val)

	out := &model.Auth{}
	err := json.Unmarshal([]byte(val_str), out)
	if err != nil {
		fmt.Println("Parse UserAuth Error:", err)
		return nil
	}
	return out
}
