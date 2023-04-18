package main

import "testing"

//func TestHello(t *testing.T) {
//	got := Hello("Chris")
//	want := "Hello,Chris"
//
//	if got != want {
//		t.Errorf("got '%q' want '%q'", got, want)
//	}
//}

/*
testing库的t.Run函数是用来在单元测试中运行一组子测试的函数。t.Run函数可以接收一个子测试的名字和一个函数，然后运行该函数。

t.Run函数可以用于将一些相似的测试用例组织到一起，方便统一管理和分组。此外，t.Run函数中的子测试可以被其他测试函数所调用。

当使用t.Run函数时，如果子测试失败会信息会单独输出，方便查找错误原因。同时，t.Run函数也提供了一个可选的函数参数t.Parallel()，可以让子测试并发地执行，从而加速整个测试过程的执行。
*/
func TestHello2(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	// Run运行f作为t的子测试，名为name。
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello,Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello,World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour,Elodie"
		assertCorrectMessage(t, got, want)
	})
}
