[TOC]

# flag

## WHAT
> 要清楚某个东西是什么？它提供了哪些数据处理和数据分析的方法？这些方法怎么使用，语法是什么？

### 什么是flag
`flag` 是用来解析命令行参数的, 也可以用来解析命令行参数的。如 `mycmd args1=var1 arg2=var2` 可以通过该包来解析

### 分类
- 基础类型：`flag.Bool` `flag.Int` `flag.Float` `flag.String` `flag.Duration` `flag.Uint` `flag.Uint8` `flag.Uint16` `flag.Uint32` `flag.Uint64` `flag.Int8` `flag.Int16` `flag.Int32` `flag.Int64` `flag.String` `flag.Complex64` `flag.Complex128`
- 自定义类型：只需要实现 `flag.Value` 接口，然后使用 `flag.Var` 注册即可。


### 如何使用flag

#### 1.基础类型使用
> 基础类型可以的使用方式一般有两种，直接返回参数值，或者通过指针返回参数值。如 `flag.Int` 和 `flag.IntVar`。
 
下面是使用方式：
```shell
# 结果是一样的

go run main.go  -age=180 
go run main.go  --age=180
go run main.go  --age 180
```
 
1. `flag.Int`
    ```go
    age := flag.Int("age", 18, "请输入你的年龄")
    flag.Parse()
    ```
2. `flag.IntVar`
    ```go
    var age int
    flag.IntVar(&age, "age", 18, "请输入你的年龄")
   flag.Parse()
    ```



#### 2.自定义结构体
1. 实现 `flag.Value` 接口
2. 使用 `flag.Var` 注册
3. 使用 `flag.Parse` 来解析命令行参数

下面是示例代码


```go
// go run main.go -student='{"name":"cx","age":180,"sex":true}'
package main

import (
   "encoding/json"
   "flag"
   "fmt"
)

func main() {
   customStruct()
}

// 自定义结构体
type studentValue struct {
   Name string `json:"name,omitempty"`
   Age  int    `json:"age,omitempty"`
   Sex  bool   `json:"sex,omitempty"`
}

// 复制一个指针内容到另外一个指针
func newStudentValue(def *studentValue, p *studentValue) *studentValue {
   *p = *def
   return p
}

// 实现 flag.Value 接口,该函数的作用是显示默认值时候调用,如执行 go run main.go --help, 输出如下 
// -student value
// 请以json的结构输入信息 (default {"name":"default"})

func (s *studentValue) String() string {
   if s == nil {
      return ""
   }

   data, _ := json.Marshal(s)
   return string(data)
}

// 实现 flag.Value 接口,该函数就是将命令行中的字符串转化到自定义结构体的方法
func (s *studentValue) Set(s2 string) error {
   err := json.Unmarshal([]byte(s2), s)
   return err

}

func customStruct() {
   def := &studentValue{
      Name: "default",
      Age:  0,
      Sex:  false,
   }

   data := &studentValue{}
   flag.Var(newStudentValue(def, data), "student", "请以json的结构输入信息")
   flag.Parse()
   fmt.Printf("%#v", data)
}

```


### 3. 使用 `flagSet` 解析自定义字符串
flag包的 `flagSet` 结构体可以用来解析自定义字符串。

```go
// go run main.go 
package main

import (
   "flag"
   "fmt"
)

func main() {
   flagSet()
}
func flagSet() {
   commandLine := flag.NewFlagSet("my_flag", flag.ExitOnError)
   name := commandLine.String("name", "defalut", "请输入你的姓名")
   // 解析该字符串
   commandLine.Parse([]string{"--name=chuxiang"})
   fmt.Println(*name)
}

```







## HOW
> 确认实现机制



## WHY

> 有什么优先的代码值得学习，