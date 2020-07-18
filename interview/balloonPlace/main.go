package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	ColorR = "R"
	ColorG = "G"
	ColorB = "B"
)

var (
	ErrNoBalloons   = errors.New("no balloons")
	ErrEmptyColor   = errors.New("empty color")
	ErrInvalidCount = errors.New("invalid count")
)

func main() {
	arrangement, err := balloonPlace([]Balloon{
		{ColorR, 3},
		{ColorG, 3},
		{ColorR, 3}, // 可重复指定颜色
		{ColorB, 2},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(arrangement)
}

// 某种气球的颜色及数量
type Balloon struct {
	Color  string
	Number int
}

type Balloons []Balloon

func (bs Balloons) Len() int {
	return len(bs)
}

func (bs Balloons) Less(i, j int) bool {
	return bs[i].Number > bs[j].Number
}

func (bs Balloons) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

// 小明有一些气球想挂在墙上装饰，他希望相同颜色的气球不要挂在一起，
// arrangement:
// 		为空表示无法按要求排列
// 		否则表示排列的方式（如： ["red", "green", "red"]
// err：
// 		可能的出错， 如参数格式有误等
// 理解： 多个气球挂在一条线上（直线排列）
func balloonPlace(bs Balloons) (sequence []string, err error) {
	if len(bs) == 0 {
		err = ErrNoBalloons
		return
	}

	var total int
	m := make(map[string]int) // color => count

	for _, b := range bs {
		if b.Color == "" { // 颜色不能为空
			err = ErrEmptyColor
			return
		}
		if b.Number <= 0 { // 数量应该为正数
			err = ErrInvalidCount
			return
		}
		total += b.Number

		m[b.Color] += b.Number // += 表示可以给出重复的颜色
	}

	order := make(Balloons, 0) // 颜色去重复后再按Count倒序排列
	for c, n := range m {
		order = append(order, Balloon{c, n})
	}
	sort.Sort(order)

	// O X O Y O Z O X O
	// 以O表示按按颜色分组后 数量最多的气球
	// 上面就是一个期望的排列
	// 且O再多一个就会导致无法生成期望的排列
	// 所以 O应该满足不大于 (total + 1) / 2
	if order[0].Number > (total+1)/2 {
		return
	}

	// 生成一个合适的排列
	sequence = method1(m, order, total)
	//sequence = method2(m, total)
	return
}

// 方法一， 插空法
// 尽一种颜色的先取完， 再取下一种颜色。隔一个空位放一个气球
func method1(m map[string]int, order Balloons, total int) (sequence []string) {
	sequence = make([]string, total)
	offset := 0 // 当前要放在哪个位置
	for _, o := range order {
		for m[o.Color] > 0 {
			m[o.Color]--
			sequence[offset] = o.Color
			offset += 2

			if offset >= total {
				offset = 1
			}
		}
	}
	return sequence
}

// 方法二
// 每次找到数量最多的前两种颜色的气球
// 交错排列
func method2(m map[string]int, total int) (sequence []string) {
	sequence = make([]string, 0, total)

	b1, b2 := top2(m)
	for {
		if b2.Number > 0 {
			for i := 0; i < b2.Number; i++ {
				sequence = append(sequence, b1.Color, b2.Color)
			}
			delete(m, b2.Color)
			m[b1.Color] -= b2.Number
		} else if b1.Number > 0 {
			sequence = append(sequence, b1.Color)
			delete(m, b1.Color)
		} else {
			return
		}
		b1, b2 = top2(m)
	}
}

func max(m map[string]int) (b *Balloon) {
	if len(m) == 0 {
		b = &Balloon{}
		return
	}

	colorOfMax := ""
	numberOfMax := 0
	for c, n := range m {
		if n > numberOfMax {
			colorOfMax = c
			numberOfMax = n
		}
	}
	b = &Balloon{
		Color:  colorOfMax,
		Number: numberOfMax,
	}
	return
}

func top2(m map[string]int) (b1 *Balloon, b2 *Balloon) {
	newM := make(map[string]int)
	for c, n := range m {
		newM[c] = n
	}
	b1 = max(newM)
	delete(newM, b1.Color)
	b2 = max(newM)
	return
}
