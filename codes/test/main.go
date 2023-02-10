package main

import "fmt"

type CunZhuang struct {
	n       int
	x, y, t []int
}

func (cz *CunZhuang) Sort() CunZhuang {
	len := len(cz.t)
	flag := true
	for i := 0; i < len && flag; i++ {
		flag = false
		for j := 0; j < len-i-1; j++ {
			if cz.t[j] > cz.t[j+1] {
				cz.t[j], cz.t[j+1] = cz.t[j+1], cz.t[j]
				cz.x[j], cz.x[j+1] = cz.x[j+1], cz.x[j]
				cz.y[j], cz.y[j+1] = cz.y[j+1], cz.y[j]
				flag = true
			}
		}
	}

	return *cz
}

type DangerZone struct {
	dangerX []int
	dangerY []int
}

// func (dz *DangerZone) AddDangerZone(cz CunZhuang) DangerZone{
// 	len:=len(cz.t)
// 	for i:=0; i<len;i++{

// 	}

// 	return *dz
// }

func CanGo(cz CunZhuang, dz DangerZone) int {
	// person := CunZhuang{}
	n := -1
	len := len(cz.t)
	for i := 0; i < len; i++ {
		return 0
	}

	return n
}

func main() {
	var guaishou = CunZhuang{}
	guaishou.n = 4
	guaishou.t = []int{2, 0, 2, 4}
	guaishou.x = []int{0, 3, 2, 0}
	guaishou.y = []int{0, 3, 1, 2}
	guaishou.Sort()
	fmt.Println(guaishou.Sort())

	// var yuyanjia = CunZhuang{}

}
