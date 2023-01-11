package util_test

import (
	"JY8752/gacha-app/util"
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

func TestLinearSearchLottery(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("LinearSearchLottery", func() {
		g.It("execute 100 times", func() {
			weights := []int{10, 20, 70}

			i := 0
			results := make(map[int]int)
			for i < 10000 {
				index := util.LinearSearchLottery(weights)
				count := results[index]
				results[index] = count + 1
				i++
			}
			fmt.Printf("0: %d 1: %d 2: %d\n", results[0], results[1], results[2])
		})
	})
}

func TestBinarySearchLottery(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("BinarySearchLottery", func() {
		g.It("execute 1000 times", func() {
			weights := []int{10, 20, 70}

			i := 0
			results := make(map[int]int)
			for i < 10000 {
				index := util.BinarySearchLottery(weights)
				count := results[index]
				results[index] = count + 1
				i++
			}
			fmt.Printf("0: %d 1: %d 2: %d\n", results[0], results[1], results[2])
		})
	})
}
