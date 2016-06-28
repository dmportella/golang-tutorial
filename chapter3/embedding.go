package chapter3

import "fmt"
import "time"
import "math/rand"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type entity struct {
	name    string
	life    int
	defence int
	skill   int
}

func roll() int {
	return r.Intn(19) + 1
}

func (attacker *entity) attack(defender *entity) {
	if (roll() + attacker.skill) >= defender.defence {
		defender.life -= attacker.skill
		fmt.Printf("%s hits %s\n\r", attacker.name, defender.name)
		return
	}
	fmt.Printf("%s misses %s\n\r", attacker.name, defender.name)
}

type Orc struct {
	entity
}

type OrcMage struct {
	entity
	fireball int
}

func (orcMage *OrcMage) attack(defender *entity) {
	if (roll() + orcMage.skill + orcMage.fireball) >= defender.defence {
		defender.life -= orcMage.fireball
		fmt.Printf("%s hits %s with fireball\n\r", orcMage.name, defender.name)
		return
	}
	fmt.Printf("%s misses %s\n\r", orcMage.name, defender.name)
}

func Embedding() {
	orc := Orc{entity{"Brakt", 15, 14, 6}}
	orcMage := OrcMage{entity{"Shorkta", 15, 10, 2}, 4}

	var round = 1
	for {
		time.Sleep(time.Second * 1)

		fmt.Println("round #", round)

		orc.attack(&orcMage.entity)
		orcMage.attack(&orc.entity)

		fmt.Println("orc life: ", orc.life, "orc mage life: ", orcMage.life)

		if orc.life <= 0 || orcMage.life <= 0 {
			break
		}
		round++
	}

}
