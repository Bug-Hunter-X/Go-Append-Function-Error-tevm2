```go
func main() {
    var i int
    fmt.Println(i)
    i = append([]int{1, 2, 3}, i) //Error: cannot use i (type int) as type []int in argument to append
    fmt.Println(i)
}
```