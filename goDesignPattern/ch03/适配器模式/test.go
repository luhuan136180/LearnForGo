package main

import "fmt"

//适配目标  抽象的技能
type Attack interface {
	Fight()
}

//具体的技能
type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了技能：大保健")
}

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用率技能")
	h.attack.Fight()
}

func main() {
	gailun := Hero{Name: "GAILUN", attack: new(Dabaojian)}
	gailun.Skill()
}
