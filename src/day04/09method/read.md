```go
// 这个int是go标准库中定义的
func (i int) hello() {
}
```

`错误用法`，我们只能给自定义的结构体添加方法。
我们可以借助关键字`type`定义我们自己的类型，从而给 int 添加方法。

```go
type myInt int

func (m myInt) test {
}
```
