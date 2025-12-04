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

func (r *PlayRepo) CreateRoom(playerName, roomName string) error {
	ctx := context.Background()

	// Check if room already exists
	result, err := r.redis.Get(ctx, `room:`+roomName).Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("failed to check room: %w", err)
	}

	if result != "" {
		return fmt.Errorf("sudah ada roomnya")
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
	ctx := context.Background()

	// Check if room exists
	result, err := r.redis.Get(ctx, `room:`+roomName).Result()
	if err == redis.Nil {
		return fmt.Errorf("room not found")
	}
	if err != nil {
		return fmt.Errorf("failed to check room: %w", err)
	}

	if result == "" {
		return fmt.Errorf("room not found")
	}

	// Add player to room
	if err := r.redis.Set(ctx, `player:`+roomName, playerName, 300*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to join room: %w", err)
	}

	return nil
}

func (r *PlayRepo) LeaveRoom(playerName, roomName string) error {
	ctx := context.Background()

	//check user exists in room
	resultCheck, err := r.redis.Get(context.Background(), `player:`+roomName).Result()
	if err == redis.Nil {
		return fmt.Errorf("player not found in room", err)
	}
	if err != nil {
		return fmt.Errorf("failed to check player in room: %w", err)
	}
	if resultCheck != playerName {
		return fmt.Errorf("player not found")
	}
	// Check if room exists
	result, err := r.redis.Get(ctx, `room:`+roomName).Result()
	if err == redis.Nil {
		return fmt.Errorf("room not found")
	}
	if err != nil {
		return fmt.Errorf("failed to check room: %w", err)
	}

	if result == "" {
		return fmt.Errorf("room not found")
	}

	// Delete player
	if err := r.redis.Del(ctx, `player:`+roomName).Err(); err != nil {
		return fmt.Errorf("failed to remove player: %w", err)
	}

	if playerName == result {
		if err := r.redis.Del(ctx, `room:`+roomName).Err(); err != nil {
			return fmt.Errorf("failed to delete room: %w", err)
		}
	}

	// Delete fight room
	if err := r.redis.Del(ctx, `fightroom:`+roomName).Err(); err != nil {
		return fmt.Errorf("failed to delete fight room: %w", err)
	}

	return nil
}
