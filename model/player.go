package model

type Player struct {
	PlayerID  string `xorm:"varchar(20)    pk"`
	Email     string `xorm:"varchar(255)   notnull unique(Email_constraint)"`
	Password  string `xorm:"varchar(255)   notnull"`
	NickName  string `xorm:"varchar(255)   notnull"`
	CreatedAt int64  `xorm:"bigint         notnull"`
	Balance   int64  `xorm:"bigint         notnull"`
}
