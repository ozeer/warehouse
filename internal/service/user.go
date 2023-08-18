package service

import (
	"context"
	"fmt"

	pb "warehouse/api/git"
	"warehouse/helper"
	"warehouse/models"

	"github.com/go-kratos/kratos/v2/metadata"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	var username, identity string

	if md, ok := metadata.FromServerContext(ctx); ok {
		// 获取经过auth中间件校验通过后的解析信息
		username = md.Get("username")
		identity = md.Get("identity")

		fmt.Println(username, identity)
	}

	user := new(models.UserBasic)
	err := models.DB.Where("username = ? AND password = ?", req.Username, helper.GetMd5(req.Password)).First(&user).Error

	if err != nil {
		return nil, err
	}

	token, err := helper.GenerateToken(user.Identity, user.UserName, 1)

	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{Token: token}, nil
}
