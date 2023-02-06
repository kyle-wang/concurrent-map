# concurrent-map
结合Golang的sync.Map和新的泛型，封装的map库

# Usage

```go
type testStruct struct {
    a int32
}

func main() {
    testMap := concurrent_map.CreateMap[int, *testStruct]()
    testMap.Store(1, &testStruct{a: 1})
    testMap.Store(2, &testStruct{a: 2})
    fmt.Println(testMap.Load(2))
    b := testMap.Load(3)
    if b == nil {
        fmt.Println("Nil")
    } else {
        fmt.Println("Not Nil")
    }
    fmt.Println(testMap.Len())
}
```