package mutex

import (
	"fmt"
	"sync"
)

type Pool struct {
	New func() interface{}
}

func (p *Pool) Get() interface{}
func (p *Pool) Put(x interface{})

type Dog struct{ Name string }

func (d *Dog) Bark() { fmt.Printf("%s", d.Name) }

var dogPack = sync.Pool{
	New: func() interface{} { return &Dog{} },
}

func main() {
	dog := dogPack.Get().(*Dog)
	dog.Name = "Rex"
	dog.Bark()
	dogPack.Put(dog)
}
