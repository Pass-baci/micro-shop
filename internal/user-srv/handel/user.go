package handel

import (
	"context"
	"micro-shop/internal/pkg/rand"
	"micro-shop/internal/user-srv/proto"
	"micro-shop/internal/user-srv/store"
	"micro-shop/internal/user-srv/svc"
)

type UserHandle interface {
	CreateUser(ctx context.Context, in *proto.CreateUserInfo) (*proto.UserInfo, error)
}

type userHandle struct {
	svcCtx *svc.Svc
}

func NewUserHandle(svcCtx *svc.Svc) UserHandle {
	return &userHandle{
		svcCtx: svcCtx,
	}
}

func (u *userHandle) CreateUser(ctx context.Context, in *proto.CreateUserInfo) (*proto.UserInfo, error) {
	var (
		err error
	)
	userId := rand.GetRandString(5)
	if _, err = u.svcCtx.Store.CreateUser(ctx, store.CreateUserParams{
		ID:       userId,
		Username: in.Name,
		Password: in.Password,
	}); err != nil {
		u.svcCtx.Logger.Errorf("【用户服务】创建用户失败，错误信息：%s", err.Error())
		return nil, err
	}

	var user store.User
	user, err = u.svcCtx.Store.GetUser(ctx, userId)
	if err != nil {
		u.svcCtx.Logger.Errorf("【用户服务】查询用户失败，错误信息：%s", err.Error())
		return nil, err
	}

	return UserRemoveToProto(&user), nil

}

func UserRemoveToProto(user *store.User) *proto.UserInfo {
	return &proto.UserInfo{
		ID:         user.ID,
		Name:       user.Username,
		Password:   user.Password,
		CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
