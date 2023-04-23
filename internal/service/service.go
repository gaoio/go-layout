package service

import (
	"github.com/google/wire"
	"go-layout/internal/conf"
	"go-layout/internal/dao"
)

var ProviderSet = wire.Struct(new(Service), "*")

type Service struct {
	Dao  *dao.Dao
	Conf *conf.Conf
}

func (s *Service) Close() {
	s.Dao.Close()
}

func (s *Service) formatPageAndPageSize(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 20
	}

	return page, pageSize
}
