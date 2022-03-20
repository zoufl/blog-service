package service

import (
	opgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/dao"
	"golang.org/x/net/context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(opgorm.WithContext(svc.ctx, global.DBEngine))

	return svc
}
