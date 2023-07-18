package dto

import (
	"ginEssential/model"
)

// UserDto是用户数据传输对象,只需要给前端返回用户的姓名和手机号即可
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUserDto方法将User转换为UserDto
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
