package service

import "guntingbatukertas/repo"

type PlayService struct {
	repo repo.PlayRepo
}

func NewPlayService(r repo.PlayRepo) *PlayService {
	return &PlayService{
		repo: r,
	}
}

func (ps *PlayService) CreateRoom(playerName string, roomName string) any {
	// anggaplah validasi sudah dilakukan disini
	return ps.repo.CreateRoom(playerName, roomName)
}

func (ps *PlayService) JoinRoom(playerName string, roomName string) error {
	// anggaplah validasi sudah dilakukan disini
	return ps.repo.JoinRoom(playerName, roomName)
}
