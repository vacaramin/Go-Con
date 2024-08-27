package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println("Total time = ", time.Since(start))
	}()

	var WG sync.WaitGroup
	enemyNinja := map[string]string{
		"Sasuke":     "Uchiha",
		"Naruto":     "Uzumaki",
		"Itachi":     "Uchiha",
		"Neji":       "Hyuga",
		"Madara":     "Uchiha",
		"Orochimaru": "Sannin",
		"Kisame":     "Hoshigaki",
		"Nagato":     "Uzumaki",
		"Obito":      "Uchiha",
		"Kakuzu":     "Takigakure",
		"Deidara":    "Iwagakure",
		"Zabuza":     "Momochi",
		"Kabuto":     "Yakushi",
		"Gaara":      "Kazekage",
		"Kankuro":    "Sabaku",
	}
	for i, ninja := range enemyNinja {
		WG.Add(1)
		go AttackNinja(i+" "+ninja, &WG)
	}
	delete(enemyNinja, "Sasuke")
	WG.Wait()

}

func AttackNinja(Ninja string, wg *sync.WaitGroup) {
	fmt.Println("Attacking ", Ninja, "!")
	fmt.Println(Ninja, "Down!")
	fmt.Println()

	defer wg.Done()
}
