package logic

import (
	"acana/dao/mysql"
	"acana/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// query all data about community from sql
	return mysql.GetCommunityList()

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
