package model

type PlayerWallet struct {
	ID         int64  `xorm:"pk autoincr"`
	PlayerID   string `xorm:"varchar(20)  notnull unique(playerID_constraint)"`
	OperatorID string `xorm:"varchar(20)  notnull unique(playerID_constraint) DEFAULT ''"`
	Bal        int64  `xorm:"bigint       notnull DEFAULT 0"`
	Point      int64  `xorm:"bigint       notnull DEFAULT 0"`
	Currency   string `xorm:"varchar(4) DEFAULT 'RMB'`
}
