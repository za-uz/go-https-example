package config

import (
	"../types"
	"time"
)

const MAX_DELAY = 10 * time.Second

func GetMap() map[string]types.Entity {
	return map[string]types.Entity{
		// users
		"user1": {"123456", 100000},
		"user2": {"123456", 100000},
		"user3": {"123456", 100000},
		"user4": {"123456", 100000},
		"user5": {"123456", 100000},
		"user6": {"123456", 100000},
		"user7": {"123456", 100000},
		// Projects
		"Project1": {},
		"Project2": {},
		"Project3": {}}
}
