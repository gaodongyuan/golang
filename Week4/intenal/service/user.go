package service

import (
	"week4/intenal/dao"
)

/**
获取名称
*/
func GetNameById(id int) (name string, err error) {


	name, err = dao.GetNameById(id)

	if err != nil {
		return "", err
	}
	return name, nil
}
