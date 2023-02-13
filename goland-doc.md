## Массив

Создание массива это по сути просто выделение куска памяти нужного размера.

```go
var arr [5]int{} // [0, 0, 0, 0, 0]
```

```go
// newarray allocates an array of n elements of type typ.
func newarray(typ *_type, n int) unsafe.Pointer {
    if n < 0 || uintptr(n) > maxSliceCap(typ.size) {
        panic(plainError("runtime: allocation size out of range"))
    }
    return mallocgc(typ.size*uintptr(n), typ, true)
}
```

## Слайс

```go
foo := make([]int, 3, 5)
foo := []int{}
```
```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```


## Особенности функций

1. В функцию аргументы передаются копированием.
   Передача сложных типов (слайсы, массивы, структуры) дорогая операция, необходимо передавать указатели.
   ```go
    func Foo(p *[]int) {}
    ```
   
## Конкуретность

### Каналы - механизм коммуникации, который, помимо прочего, позволяет обмениваться данными между горутинами
```go
 ch := make(chan int)
 
 ch <- 1
 
 go func() {
   ch <- 2
   ch <- 3 
 }()
 
 for v := range ch {
	 fmt.Println(v)
 }
 close(ch)
```

### Pipeline