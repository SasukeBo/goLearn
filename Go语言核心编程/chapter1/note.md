# 基础知识

## 初识GO程序

```go
// hello.go
package main

import "fmt"

func main() {
  fmt.Printf("Hello, world. 你好，世界！\n")
}
```

- 包名为main，main是可执行程序的包名，所有的Go源程序文件头部都必须有一个包生命语句，Go通过包来管理命名空间。
- import 引用一个外部包，可以是标准库的包，也可以是第三方或自定义的包。
- 和C语言一样，main函数代表程序的入口。

## Go词法单元

### token

token是构成源程序的基本不可再分割的单元。
Go语言的token可以分为关键字、标识符、操作符、分隔符和字面常量等。

### 标识符

Go的标识符构成规则是：开头一个字符必须是字母或下划线，后面跟任意多个字符、数字或下划线。并且区分大小写。

**内置数据类型标识符20个**

- 数值（16个）
  - 整形（12个）
    > byte int int8 int16 int32 int64
    > uint uint8 uint16 uint32 uint64 uintptr

  - 浮点型（2个）
    > float32 float64

  - 复数型（2个）
    > complex64 complex128

- 字符和字符串型（2个）
  > string rune

- 接口型（1个）
  > error

- 布尔型（1个）
  > bool

Go是一种强类型静态编译语言。Go语言支持自动类型推导。

**内置函数15个**

```go
make new len cap append copy delete panic recover close complex real image print println
```
内置函数也是高级语言的一种语法糖，由于是语言内置的，不需要import任何包就可以直接调用。

**常量值标识符4个**

```go
true false
iota // 用在连续的枚举类型的声明中
nil // 指针/引用型的变量的默认值就是nil
```

**空白标识符1个**

```go
_
```
空白标识符有特殊的含义，用来声明一个匿名变量。该变量在赋值表达式的左端，空白标识符引用通常被用作占位，比如忽略函数多个返回值中的一个和强制编译器做类型检查。

### 字面常量

字面常量表示固定值，简称字面量。
Go的字面量可以出现在两个地方：
- 用于常量和变量的初始化。
- 用在表达式里或作为函数调用参数。


## 变量和常量

### 变量

使用一个名称来绑定一块内存地址，该内存地址中存放的数据类型由定义变量时指定的类型决定，该内存地址存放的内容可以改变。

1. 显示的完整声明
```go
var varName dataType [ = value ]
```
2.短类型声明
```go
varName := value
```
*`:=`声明只能出现在函数内（方法内）*

### 常量

常量使用一个名称来绑定一块内存地址，该地址中存放的数据类型由定义常量时指定的类型决定，而且该内存地址里面存放的内容不可以改变。
常量存储在程序的只读段内。

预声明标识符iota用在常量声明中，其初始值为0。

```go
// 类似枚举的iota

const (
  c0 = iota // c0 == 0
  c1 = iota // c1 == 1
  c2 = iota // c2 == 2
)

// 简写模式
const (
  c0 = iota // c0 == 0
  c1
  c2
)

// 注意iota逐行增加
const (
  a = 1 << iota // a == 1 (iota == 0)
  b = 1 << iota // b == 2 (iota == 1)
  c = 3         // c == 3 (iota == 2, unused)
  d = 1 << iota // d == 8 (iota == 3)
)

const (
  u = iota * 42
  v float64 = iota * 42
  w = iota * 42
)

// 分开的const语句，iota每次都是从0开始
const x = iota // x == 0
const y = iota // x == 0
```

## 基本数据类型

### rune类型

Go内置两种字符类型：
- 一种是byte的字节类类型（byte 是 uint 的别名）
- 另一种是标识Unicode编码的字符rune。rune在Go内部是int32类型的别名，占用4个字节。

## 复合数据类型

指针、数组、切片、字典、通道、结构和接口。

### 指针

1. Go语言支持指针，同样支持多级指针`**T`。

2. 结构体指针访问结构体字段仍然使用'.'操作符，不同于C/C++
```go
type User struct {
  name string
  age int
}

andes := User{
  name: "andes",
  age: 18,
}

p := &andes
fmt.Println(p.name) // p.name通过'.'`操作符访问成员变量。
```

3. Go 不支持指针的运算。
Go由于支持垃圾回收，如果支持指针运算，则会给垃圾回收的实现带来很多不方便。因此Go在语言层面禁止使用指针运算。

4. 函数中允许返回局部变量的地址
Go编译器使用“栈逃逸”机制将这种局部变量的空间分配在堆上。
```go
  func sum (a, b int) *int {
    sum := a+b
    return &sum // 允许， sum会分配在heap上
  }
```

### 数组

数组的类型名是`[n]elementType`。单独声明一个数组类型变量而不进行初始化时没有意义的。
```go
var arr [2]int // 声明一个有两个整形的数组， 但元素默认值都是0，一般很少这样使用。
array := [...]float64{7.0, 8.5, 9.1} // [...]后面跟字面量初始化列表。
```

**数组初始化**

```go
a := [3]int{1,2,3} // 指定长度和初始化字面量
a := [...]int{1,2,3} // 不指定长度，但是由后面的初始化列表元素数量来决定数组长度。
a := [3]int{1:1, 2:3} // 指定总长度，并通过索引值进行初始化，没有初始化元素时使用类型默认值。
a := [...]int{1:1, 2:3} // 不指定总长度，通过索引值进行初始化，数组长度由最后一个索引值确定，没有指定索引的元素被初始化为类型的零值。
```

**数组的特点**

- 数组创建完长度就固定了，不可以再追加元素
- 数组是值类型，数组赋值或作为函数参数都是值拷贝
- 数组长度是数组类型的组成部分，[10]int 和 [20]int 表示不同类型
- 可以根据数组创建切片

**数组相关操作**

- 数组元素访问

```go
a := [...]int{1,2,3}
b := a[0]

for i, v := range a {

}
```

- 数组长度。

```go
a := [...]int{1,2,3}
alengh := len(a)

for i:=0; i < alengh; i++ {

}
```

### 切片

Go语言的数组的定长性和值拷贝限制了其使用场景，Go提供另一种数据类型slice，这是一种变长数组，其数据结构中有指向数组的指针，所以是一种引用类型。

```go
type slice struct {
  array unsafe.Pointer // 指向底层数组的指针
  len int // 切片的元素数量
  cap int // 底层数组的容量
}
```

1. 切片的创建

- 由数组创建

```go
// slice.go
var array = [...]int{0, 1, 2, 3, 4, 5, 6}
s1 := array[0:4]
s2 := array[:4]
s3 := array[2:]
fmt.Printf("%v\n", s1)
fmt.Printf("%v\n", s2)
fmt.Printf("%v\n", s3)
```

- 通过内置函数make创建切片

注意：由make创建的切片各元素被默认初始化为切片元素类型的零值。

```go
// len = 10, cap = 10
a := make([]int  10) // [0 0 0 0 0 0 0 0 0 0]

// len = 10, cap = 15
b := make([]int  10  15) // [0 0 0 0 0 0 0 0 0 0]
```

2. 切片支持的操作

- 内置函数len()返回切片的长度。
- 内置函数cap()返回切片底层数组容量。
- 内置函数append()对切片追加元素。
- 内置函数copy()用于复制一个切片。

```go
a := [...]int{0, 1, 2, 3, 4, 5, 6}
b := make([]int, 2, 4)
c := a[0:3]

fmt.Println(len(b))
fmt.Println(cap(b))
b = append(b, 1)
fmt.Println(b)
fmt.Println(len(b))
fmt.Println(cap(b))

b = append(b, c...)
fmt.Println(b)
fmt.Println(len(b))
fmt.Println(cap(b))

d := make([]int, 2, 2)
copy(d, c)
fmt.Println(d)
fmt.Println(len(d))
fmt.Println(cap(d))
```

### map

Go 语言内置的字典类型叫map。map的类型格式是：`map[K]T`，其中K可以是任意可以进行比较的类型，T是值类型。map也是一种引用类型。

1. map的创建

- 使用字面量创建。

```go
ma := map[string]int{"a":1, "b":2}
fmt.Println(ma["a"])
fmt.Println(ma["b"])
```

- 使用内置的make函数创建。

```go
// map.go
make(map[K]T) // map的容量使用默认值
make(map[K]T, len) // map的容量使用给定的len值

mp1 := make(map[int]string)
mp2 := make(map[int]string, 10)
mp1[1] = "tom"
mp2[1] = "pony"

fmt.Println(mp1[1])
fmt.Println(mp2[1])
```

2. map支持的操作

- map的单个键值访问格式为`mapName[key]`，更新某个key值时只需要将`mapName[key]`放到等式左边。
- 可以使用range遍历一个map类型变量，但是不保证每次迭代元素的顺序。
- 删除map中的某个键值，使用如下语法，delete(mapName, key)。delete是内置函数。
- 可以使用内置的len()函数返回map中的键值对个数。

```go
// map2.go
mp := make(map[int]string)
mp[1] = "tom"
mp[1] = "pony"
mp[2] = "jaky"
mp[3] = "andes"
delete(mp, 3)

fmt.Println(map[1])
fmt.Println(len(mp))

for k, v := range mp {
  fmt.Println("key = ", k, " value = ", v)
}
```

**注意**

- Go内置的map不是并发安全的，并发安全的map可以使用标准包sync中的map。
- 不要直接修改map value内某个元素的值，如果想修改map的某个键值，则必须整体赋值。

```go
// map3.go
type User struct {
  name string
  age int
}

ma := make(map[int]User)
andes := User {
  name: "andes",
  age: 18,
}

ma[1] = andes
// ma[1].age = 19 // ERROR, 不能通过map引用直接修改结构体成员
andes.age = 19
ma[1] = andes // 必须整体替换value

fmt.Printf("%v\n", ma)
```

### struct

Go的struct类型和C类似
- struct结构中的类型可以是任意类型
- struct的存储空间是连续的，其字段按照声明时的顺序存放。

struct有两种形式：
1. struct类型字面量

```go
struct {
  FieldName FieldType
  FieldName FieldType
  FieldName FieldType
}
```

2. 使用type声明的自定义struct类型

```go
type TypeName struct {
  FieldName FieldType
  FieldName FieldType
  FieldName FieldType
}
```

**struct类型变量的初始化**

```go
type Person struct {
  Name string
  Age int
}

type Student struct {
  *Person
  Number int
}

// 按照类型声明顺序，逐个赋值
// 不推荐这种初始化方式，一旦struct增加字段，则整个初始化语句会报错
a := Person{"Tom", 21}

// 推荐使用这种
a := &Person{
  Name: "tata",
  Age: 12,
}

s := Student{
  Person: p,
  Number: 110,
}
```

**其他复合类型**

接口后续介绍。


## 控制结构

### if 语句

特点

- if后面的条件判断子句不需要用小括号括起来。
- {必须放在行尾，和if或if else放在一行。
- if后面可以带一个简单的初始化语句，并以分号分割，该简单语句声明的变量的作用域是整个if语句块，包括后面的else if和else 分支。
- Go语言没有条件运算符（a>b?a:b)

```go
if x <= y {
  return y
} else {
  return x
}
```
一个完整的if else示例
```go
if x := f(); x < y { // 初始化语句中的声明变量x
  return x
} else if x > z { // x 在else if 里面一样可以被访问
  return z
} else {
  return y
}
```

**最佳实践**

- 尽量减少条件语句的复杂度
- 尽量减少if语句的嵌套层次

例如：
```go
if err, file := os.Open("xxx"); err == nil {
  defer file.close()
  // do someting
} else {
  return nil, err
}
```

重构为：
```go
err, file := os.Open("xxx")
if err != nil {
  return nil, err
}

defer file.Close()
// do something
```

### switch

**特性**

- switch语句也可以带一个可选的初始化语句
- switch后面的表达式可是可选的，如果没有表达式，则case子句是一个布尔表达式。
- 条件表达式可以是任意支持相等比较运算的类型变量，C只可以是整数。
- 通过fallthrough语句来强制执行下一个case，不在判断下一个case子句的条件是否满足。
- switch和.(type)结合可以进行类型的查询。

```go
// switch.go

switch i := "y"; i { // switch后面可以带上一个初始化语句
  case "y", "Y": // 多个case值使用逗号分隔
    fmt.Println("yes")
    fallthrough

  case "n", "N":
    fmt.Println("no")
}

score := 85
grade := ' '

if score >= 90 {
  grade = 'A'
} else if score >= 80 {
  grade = 'B'
} else if score >= 70 {
  grade = 'C'
} else if score >= 60 {
  grade = 'D'
} else {
  grade = 'F'
}

// 上面的if else 可以改写为下面的switch语句
switch {
  case score >= 90:
    grade = 'A';
  case score >= 80:
    grade = 'B';
  case score >= 70:
    grade = 'C';
  case score >= 60:
    grade = 'D';
  default:
    grade = 'F';
}

fmt.Printf("grade = %v", grade)
```


### for

Go 语句仅支持一种循环语句就是for
三种使用场景：
- 类似C里面的for循环语句
```go
for init; condition; post {}
```
- 类似C里面的while
```go
for condition {}
```
- 死循环
```go
for {}
```

for对数组、切片、字符串、map和通道的访问。

```go
// 访问map
for key, value := range map {}
for ket := range map {}

// 访问数组
for index, value := range array{}
for index := range array{}
for _, value := range array{}

// 访问切片
for index, value := range slice{}
for index := range slice{}
for _, value := range slice{}

// 访问通道
for value := range channel{}
```

### 标签和跳转

**标签**

Go语言使用标签Label来标识一个语句的位置，用于goto、break、continue语句的跳转，标签的语法是：
```go
Label: Statement
```
**goto**

goto语句用于函数内部的跳转。需要配合标签一起使用。
goto语句不能只能跳到同级作用域，或者上层作用域，而不能跳到内部作用域。
```go
if n%2 == 1 {
  goto L1
}
for n > 0 {
  f()
  n--
}

L1:
  f()
  n--
```

**break**

break用于函数内跳出for、switch、select语句的执行;
和标签一起使用可以跳出标签所标识的for、switch、select语句的执行。可用于跳出多重循环，但标签和break必须在同一个函数内。

```go
L1:
  for i := 0; ; i++ {
    for j := 0; ; j++ {
      if i >= 5 {
        // 跳出L1标签所在的for循环
        break L1
      }
      if j > 10 {
        // 默认仅跳出离break最近的内存循环
        break
      }
    }
  }
```

**continue**

除了C语言中的用法还可以结合标签使用，跳出标签指示的for循环的本次循环

```go
L1:
  for i := 0; ; i++ {
    for j := 0; ; j++ {
      if i >= 5 {
        // 跳出L1标签所在的for循环i++处执行
        continue L1
        // the following is not executed
      }
      if j > 10 {
        // 默认仅跳到离break最近的内层循环j++处执行
        continue
      }
    }
  }
```

**return和函数调用**

return 语句也能引发控制流程的跳转，用于函数和方法的退出。函数和方法的调用也能引发程序控制流的跳转，这些在后续章节中会详细介绍。
