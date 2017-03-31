package main

import "log"

var (
	yy []Hoge
	zz map[string]*Hoge = make(map[string]*Hoge)
)

type Hoge struct {
	a int
	b int
}

func main() {
	d1 := Hoge{}
	d1.a = 10
	d1.b = -20
	yy = append(yy, d1)

	d2 := Hoge{}
	d2.a = 5
	d2.b = -8
	yy = append(yy, d2)

	zz["d1"] = &yy[0]
	zz["d2"] = &yy[1]

	log.Println("yy[0] = ", yy[0], "*zz[d1] = ", *zz["d1"])
	log.Println("yy[1] = ", yy[1], "*zz[d2] = ", *zz["d2"])

	yy[0].a = -9
	yy[0].b = 888
	yy[1].a = -6
	yy[1].b = 444

	// zzにも変更が反映されると思い込んでた
	log.Println("yy[0] = ", yy[0], "*zz[d1] = ", *zz["d1"])
	log.Println("yy[1] = ", yy[1], "*zz[d2] = ", *zz["d2"])
}
