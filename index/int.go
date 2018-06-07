package main

// import (
// 	"fmt"
// 	"github.com/google/btree"
// )

// type Int int

// func (a Int) Less(b btree.Item) bool { return a < b.(Int) }

// func main() {
// 	tr := btree.New(32)
// 	for i := Int(0); i < 10; i++ {
// 		tr.ReplaceOrInsert(i)
// 	}
// 	fmt.Println("len:       ", tr.Len())
// 	fmt.Println("get3:      ", tr.Get(Int(3)))
// 	fmt.Println("get100:    ", tr.Get(Int(100)))
// 	fmt.Println("del4:      ", tr.Delete(Int(4)))
// 	fmt.Println("del100:    ", tr.Delete(Int(100)))
// 	fmt.Println("replace5:  ", tr.ReplaceOrInsert(Int(5)))
// 	fmt.Println("replace100:", tr.ReplaceOrInsert(Int(100)))
// 	fmt.Println("min:       ", tr.Min())
// 	fmt.Println("delmin:    ", tr.DeleteMin())
// 	fmt.Println("max:       ", tr.Max())
// 	fmt.Println("delmax:    ", tr.DeleteMax())
// 	fmt.Println("max:       ", tr.Max())
// 	fmt.Println("len:       ", tr.Len())
// }

// // Output:
// // len:        10
// // get3:       3
// // get100:     <nil>
// // del4:       4
// // del100:     <nil>
// // replace5:   5
// // replace100: <nil>
// // min:        0
// // delmin:     0
// // max:        100
// // delmax:     100
// // len:        8
