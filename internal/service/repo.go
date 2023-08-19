package service

import (
	"context"
	"errors"
	"os"
	"os/exec"

	pb "warehouse/api/git"
	"warehouse/define"
	"warehouse/helper"
	"warehouse/models"

	"gorm.io/gorm"
)

type RepoService struct {
	pb.UnimplementedRepoServer
}

func NewRepoService() *RepoService {
	return &RepoService{}
}

func (s *RepoService) CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.CreateRepoReply, error) {
	// 查重
	var cnt int64
	err := models.DB.Model(new(models.RepoBasic)).Where("path = ?", req.Path).Count(&cnt).Error

	if err != nil {
		return nil, err
	}

	if cnt > 0 {
		return nil, errors.New("存在同名仓库")
	}

	// 落库
	repo := &models.RepoBasic{
		Identity: helper.GetUUID(),
		Name:     req.Name,
		Path:     req.Path,
		Desc:     req.Desc,
		Type:     req.Type,
	}

	// 事务处理
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(repo).Error

		if err != nil {
			return err
		}

		// init repo path
		// mkdir path
		gitRepoPath := define.RepoPath + string(os.PathSeparator) + req.Path
		err = os.MkdirAll(gitRepoPath, 0755)
		if err != nil {
			return err
		}
		// git init --bare
		cmd := exec.Command("/bin/bash", "-c", "cd "+gitRepoPath+" ; git init --bare")
		err = cmd.Run()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateRepoReply{}, nil
}
func (s *RepoService) UpdateRepo(ctx context.Context, req *pb.UpdateRepoRequest) (*pb.UpdateRepoReply, error) {
	repo := models.RepoBasic{
		Name: req.Name,
		Desc: req.Desc,
		Type: req.Type,
	}
	err := models.DB.Model(new(models.RepoBasic)).Where("identity = ?", req.Identity).Updates(repo).Error

	if err != nil {
		return nil, err
	}

	return &pb.UpdateRepoReply{}, nil
}
func (s *RepoService) DeleteRepo(ctx context.Context, req *pb.DeleteRepoRequest) (*pb.DeleteRepoReply, error) {
	var repo *models.RepoBasic

	err := models.DB.Model(new(models.RepoBasic)).Where("identity = ?", req.Identity).Delete(&repo).Error

	if err != nil {
		return nil, err
	}

	return &pb.DeleteRepoReply{}, nil
}
func (s *RepoService) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoReply, error) {
	var repo *models.RepoBasic
	err := models.DB.Model(new(models.RepoBasic)).Where("identity = ?", req.Identity).First(&repo).Error

	if err != nil {
		return nil, err
	}

	return &pb.GetRepoReply{
		Name: repo.Name,
		Desc: repo.Desc,
		Path: repo.Path,
		Star: repo.Star,
		Type: repo.Type,
	}, nil
}
func (s *RepoService) ListRepo(ctx context.Context, req *pb.ListRepoRequest) (*pb.ListRepoReply, error) {
	repos := make([]*models.RepoBasic, 0)
	var cnt int64
	err := models.DB.Model(new(models.RepoBasic)).Count(&cnt).Offset(int(req.Page-1) * int(req.PageSize)).Limit(int(req.PageSize)).Find(&repos).Error

	if err != nil {
		return nil, err
	}

	list := make([]*pb.ListRepoItem, 0, len(repos))
	for _, v := range repos {
		list = append(list, &pb.ListRepoItem{
			Identity: v.Identity,
			Name:     v.Name,
			Desc:     v.Desc,
			Path:     v.Path,
			Star:     v.Star,
		})
	}

	return &pb.ListRepoReply{List: list, Cnt: cnt}, nil
}
