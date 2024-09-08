package converter

import (
	"rwa/internal/utils"
	"rwa/pkg/model"
	"rwa/pkg/msg"
	"time"
)

func ToProfile(user *model.User) *msg.UserProfile {
	return &msg.UserProfile{
		Inner: msg.InnerContent{
			Email:     user.Email,
			Bio:       user.Bio,
			Following: user.Following,
			Image:     user.Image,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func ToUser(register *msg.RegisterMessage) *model.User {
	return &model.User{
		Email:        register.Inner.Email,
		Username:     register.Inner.Username,
		PasswordHash: utils.HashPass(register.Inner.Password, utils.RandStringRunes(8)),
	}
}

func FromLogon(logon *msg.LogonMessage) *model.User {
	return &model.User{
		Email:        logon.Inner.Email,
		PasswordHash: []byte(logon.Inner.Password),
	}
}

func Merge(user *model.User, profile *msg.UserProfile) model.User {
	result := model.User{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		Bio:          user.Bio,
		Image:        user.Image,
		Following:    user.Following,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    time.Now(),
	}
	if profile.Inner.Email != "" {
		result.Email = profile.Inner.Email
	}
	if profile.Inner.Username != "" {
		result.Username = profile.Inner.Username
	}
	// if profile.Inner.Bio != "" {
	// 	result.Bio = profile.Inner.Bio
	// }
	// if profile.Inner.Image != "" {
	// 	result.Image = profile.Inner.Image
	// }
	result.Bio = profile.Inner.Bio
	result.Image = profile.Inner.Image
	result.Following = profile.Inner.Following
	return result
}
