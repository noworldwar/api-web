package entity

import (
	"github.com/noworldwar/api-web/internal/model"
)

func InsertPlayer(m model.Player) error {
	_, err := model.MyDB.Insert(m)
	return err
}

func GetPlayer(playerID string) (*model.Player, error) {
	m := new(model.Player)
	_, err := model.MyDB.Where("PlayerID=?", playerID).Get(m)
	return m, err
}

func UpdatePlayer(m model.Player) (int64, error) {
	affected, err := model.MyDB.Where("PlayerID=?", m.PlayerID).Update(m)
	return affected, err
}
