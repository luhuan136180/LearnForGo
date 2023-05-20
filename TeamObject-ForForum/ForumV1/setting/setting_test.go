package setting

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	if err := Init(); err != nil {
		fmt.Printf("init settings failed,err:%#v\n", err)
		return
	}
	fmt.Println(Conf.Port)
	fmt.Println(Conf.Host)
	fmt.Println(Conf.Name)
	fmt.Println(Conf.Version)
	fmt.Println(Conf.StartTime)
	fmt.Println(Conf.MachineID)

}
