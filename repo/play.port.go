package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type PlayRepo struct {
	redis *redis.Client
}

func NewPlayRepo(redis *redis.Client) *PlayRepo {
	return &PlayRepo{redis: redis}
}

func (r *PlayRepo) CreateRoom(playerName, roomName string) any {
	getRoom := r.redis.Get(context.Background(), `room:`+roomName)
	result, err := getRoom.Result()
	if err != nil {
		fmt.Println("error", err)
	}
	if result != "" {
		return "sudah ada roomnya"
	} else {
		setRoom := r.redis.Set(context.Background(), `room:`+roomName, playerName, 300*time.Second)
		setFightRoom := r.redis.Set(context.Background(), `fightroom:`+roomName, "", 300*time.Second)
		if setRoom.Err() != nil {
			return setRoom.Err()
		}
		if setFightRoom.Err() != nil {
			return setFightRoom.Err()
		}
	}
	return nil
}

func (r *PlayRepo) JoinRoom(playerName, roomName string) error {

	return nil
}
