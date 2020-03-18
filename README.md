# utils
go utils func

# install
```shell
go get github.com/ganodermaking/utils
```

# example
```go
// 元素是否在切片或数组内
if utils.In([]int{1, 2, 3, 4}, 1) {
    // ...
}

// 简化字符串转int64
int64 := utils.ParseInt(str)

// 简化字符串转int32
int32 := utils.Atoi(str)

// 简化复制
utils.Copy(&employees, &user)
```
