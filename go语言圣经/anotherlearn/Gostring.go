package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person2) Gostring() string {
	return "&Person{Name is \"+this.Name+\", Age is \"+strconv.Itoa(this.Age)+\", Sex is \"+strconv.Itoa(this.Sex)+\"}"
}
func main() {
	p := &Person2{"polaris", 28, 0}
	fmt.Printf("%#v", p)
}
