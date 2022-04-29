package proxy

import "errors"

type ProxyDataBase struct {
	Users map[string]bool
	Db    *DataBase
}

func (pDb ProxyDataBase) GetData(user string) ([]string, error) {
	if !pDb.Users[user] {
		return nil, errors.New("Insufficient access rights")
	}
	return pDb.Db.GetData(user)
}
