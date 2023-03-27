package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

//1.1.3
//func main() {
//	var name string
//	flag.StringVar(&name, "name", "GO语言编程之旅", "帮助信息")
//	flag.StringVar(&name, "n", "GO语言编程之旅", "帮助信息")
//	flag.Parse()
//
//	log.Printf("name:%s", name)
//
//}

//1.1.3
//var name string
//
//func main() {
//	flag.Parse()
//	args := flag.Args()
//	if len(args) <= 0 {
//		return
//	}
//
//	switch args[0] {
//	case "go":
//		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
//		goCmd.StringVar(&name, "name", "GO语言", "帮助信息")
//		_ = goCmd.Parse(args[1:])
//	case "php":
//		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
//		phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")
//		_ = phpCmd.Parse(args[1:])
//
//	}
//
//	log.Printf("name:%s", name)
//
//}

//1.2
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execte err :%v", err)
	}
}
