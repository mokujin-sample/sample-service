package sampleService

import (
	svc "sample-service"
)

type service struct {
	repository svc.Repository
}

func NewService(objectRepo svc.Repository) svc.ObjectService {
	return &service{
		repository: objectRepo,
	}
}

func (s *service) Get(id uint64) (object svc.Object, err error) {
	object, err = s.repository.Get(id)
	if err != nil {
		return object, err
	}

	return s.repository.Get(id)
}

func (s *service) GetAll(as []uint32, ip []string, limit uint32) ([]svc.Object, int64, error) {
	return s.repository.GetAll(as, ip, limit)
}
