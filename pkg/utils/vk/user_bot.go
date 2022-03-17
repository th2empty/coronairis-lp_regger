package vk

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

func GetUserId(token string) (int, error) {
	vk := api.NewVK(token)
	userInfo, err := vk.AccountGetProfileInfo(nil)
	if err != nil {
		return 0, err
	}

	return userInfo.ID, err
}
