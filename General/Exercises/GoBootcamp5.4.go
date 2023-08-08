package main

import "fmt"

var (
	coins        = 50
	distribution = make(map[string]int, len(users))
	users        = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	conversionTable = map[string]int{
		"a": 1,
		"A": 1,
		"e": 1,
		"E": 1,
		"i": 2,
		"I": 2,
		"o": 3,
		"O": 3,
		"u": 4,
		"U": 4,
	}
)

const (
	MAX_COINS_PER_USER = 10
)

func main() {
	countUserCoins := func(userName string) (userCoins int) {
		for _, letter := range userName {
			userCoins += conversionTable[string(letter)]
			if userCoins > MAX_COINS_PER_USER {
				return MAX_COINS_PER_USER
			}
		}
		return
	}
	for _, user := range users {
		userCoins := countUserCoins(user)
		coins -= userCoins
		distribution[user] = userCoins
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
