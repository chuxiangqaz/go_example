package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

// flag
// flag 实现了命令行参数的解析
// 可以通过三种方式去解析 flag.Int flag.IntVar customVar 等方式去解析
// 可以使用 flag.FlagSet 类型去解析非命令行参数, 函数参见: flagSet
// tips: 当存在为解析的参数就会发生 panic
func main() {
	flagToType()
	//flagToVar()
	//customType()
	//customStruct()
	//flagSet()
}

/*
$ go run main.go  --help
Usage of C:\Users\admin\AppData\Local\Temp\go-build3327055719\b001\exe\main.exe:
-age int
	请输入你的年龄 (default 18)
-name string
	请输入你的姓名
-sex
	是否是男生
*/

// 下面两个结果一样
// go run main.go  -name=chuxiang -age=180 -sex=true
// go run main.go  --name=chuxiang --age=180 --sex=true
//
//	go run main.go  --name chuxiang --age 180 --sex true
func flagToType() {
	age := flag.Int("age", 18, "请输入你的年龄")
	name := flag.String("name", "", "请输入你的姓名")
	sex := flag.Bool("sex", false, "是否是男生")
	flag.Parse()
	fmt.Printf("user info \n\tname:%s \n\tage:%d\n\tis man:%t\n", *name, *age, *sex)
}

// 同 flagToType 一样,这是解析函数不一样
func flagToVar() {
	var name string
	var age int
	var sex bool
	flag.StringVar(&name, "name", "", "请输入你的姓名")
	flag.IntVar(&age, "age", 18, "请输入你的年龄")
	flag.BoolVar(&sex, "sex", false, "是否是男生")
	flag.Parse()
	fmt.Printf("user info \n\tname:%s \n\tage:%d\n\tis man:%t\n", name, age, sex)
}

type customVar int

func newCustomVar(val customVar, p *customVar) *customVar {
	*p = val
	return (*customVar)(p)
}

func (c *customVar) String() string {
	return strconv.Itoa(int(*c))
}

func (c *customVar) Set(s string) error {
	t, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*c = customVar(t)
	return nil
}

func customType() {
	var data customVar
	value := newCustomVar(-1, &data)
	flag.Var(value, "data", "请输入一个正整数")
	flag.Parse()
	println(data)
}

type studentValue struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Sex  bool   `json:"sex,omitempty"`
}

func newStudentValue(def *studentValue, p *studentValue) *studentValue {
	*p = *def
	return p
}

func (s *studentValue) String() string {
	if s == nil {
		return ""
	}

	data, _ := json.Marshal(s)
	return string(data)
}

func (s *studentValue) Set(s2 string) error {
	err := json.Unmarshal([]byte(s2), s)
	return err

}

// go run main.go -student='{"name":"cx","age":180,"sex":true}'
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

func flagSet() {
	commandLine := flag.NewFlagSet("my_flag", flag.ExitOnError)
	name := commandLine.String("name", "defalut", "请输入你的姓名")
	commandLine.Parse([]string{"--name=chuxiang"})
	fmt.Println(*name)
}
