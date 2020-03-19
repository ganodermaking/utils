# utils
simplify go utils func

# install
```shell
go get github.com/ganodermaking/utils
```

# example
```go
// 截取字符串
utils.Substr("0123456789X", 5, 4)

// 截取字符串（包含中文字符）
utils.MbSubstr("01234中国6789X", 5, 4)

// 元素是否在切片或数组内
if utils.InArray([]int{1, 2, 3, 4}, 1) {
    // ...
}

// 切片或数组是否在切片或数组内
if utils.InArray([]int{1, 2, 3, 4}, []int{1, 4}) {
    // ...
}

// 简化字符串转int64
int64 := utils.ParseInt(str)

// 简化字符串转int32
int32 := utils.Atoi(str)

// 简化复制
utils.Copy(&employees, &user)
```
