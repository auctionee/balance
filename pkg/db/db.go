package db

import "github.com/auctionee/balance/pkg/data"

func GetInfo() data.Balance {
	return data.Balance{}
}
func Modify(info data.Modify) (data.Balance, error) {
	return data.Balance{}, nil
}
