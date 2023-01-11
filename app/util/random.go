package util

import (
	"math/rand"
	"time"
)

/*
線形探索で重み付抽選する
@return 当選した要素のインデックス
*/
func LinearSearchLottery(weights []int) int {
	//  重みの総和を取得する
	var total int
	for _, weight := range weights {
		total += weight
	}

	// 乱数取得
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	rnd := rand.Intn(total)

	var currentWeight int
	for i, w := range weights {
		// 現在要素までの重みの総和
		currentWeight += w

		if rnd < currentWeight {
			return i
		}
	}

	// たぶんありえない
	return len(weights) - 1
}

/*
二分探索で重み付抽選する
@return 当選した要素のインデックス
*/
func BinarySearchLottery(weights []int) int {
	// 重みの総和
	var total int
	// 重みの累積
	var cumulative []int

	for _, w := range weights {
		total += w
		cumulative = append(cumulative, total)
	}

	// 乱数取得
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	rnd := rand.Intn(total)

	// 探索
	minIndex := 0
	maxIndex := len(weights) - 1
	for minIndex < maxIndex {
		// 探索範囲の中央要素取得
		centerIndex := (minIndex + maxIndex) / 2
		centerPoint := cumulative[centerIndex]

		if rnd > centerPoint {
			// 乱数値が現在要素の範囲の外
			minIndex = centerIndex + 1
		} else {
			// 現在要素の範囲の下限取得
			var prevPoint int
			if centerIndex > 0 {
				prevPoint = cumulative[centerIndex-1]
			} else {
				prevPoint = 0
			}

			// 乱数値が現在要素と上限の間なら確定
			if rnd >= prevPoint {
				return centerIndex
			}

			maxIndex = centerIndex
		}
	}
	// 探索範囲が1要素しかなくなったら確定
	return maxIndex
}
