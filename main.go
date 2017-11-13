package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Rank struct {
	id    int
	score int
}

type Ranking []Rank

type Testtype struct {
	infos Ranking
}

func (r Ranking) Len() int {
	return len(r)
}

func (r Ranking) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranking) Less(i, j int) bool {
	return r[i].score < r[j].score
}

func main() {
	rand.Seed(time.Now().UnixNano())
	datas := &Testtype{}

	start := time.Now()
	for i := 0; i < 300000; i++ {
		datas.infos = append(datas.infos, Rank{i, rand.Intn(100000)})
	}
	end := time.Now()

	fmt.Println((end.Sub(start)).Seconds())

	sort.Sort(sort.Reverse(datas.infos))

	start = time.Now()

	for i := 0; i < 300000; i++ {
		for j := 0; j < 300000; j++ {
			if datas.infos[j].id == i {
				break
			}
		}
	}
	end = time.Now()

	fmt.Println((end.Sub(start)).Seconds())
}
