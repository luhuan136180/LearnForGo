package main

import "fmt"

func main() {
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%L", p)

	//

}
func (this *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(this.String()))
		f.Write([]byte(" Person has three fields."))
	} else {
		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		f.Write([]byte(fmt.Sprintln(this.String())))
	}
}
