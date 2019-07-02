# 第3章 类型系统

## 类型简介

### 命名类型和未命名类型

**命名类型**
类型可以通过标识符来表示，这种类型称为命名类型。

**未命名类型**
一个类型由预声明类型、关键字和操作符组合而成，这个类型称为未命名类型。未命名类型又称为类型字面量。
前面所说的结构和接口是未命名类型，举个例子：

```go
package main

import "fmt"

// 使用type声明的是命名类型
type Person struct {
  name String
  age int
}

func main() {
  // 使用struct字面量声明的是未命名类型
  a := struct {
    name string
    age int
  }{"andes", 18}

  fmt.Printf("%T\n", a) // struct { name string; age int }
  fmt.Printf("%v\n", a) // { andes 18 }

  b := Person("tom", 21)
  fmt.Printf("%T\n", b) // main.Person
  fmt.Printf("%v\n", b) // { tom 21 }
}
```
1. `未命名类型`和`类型字面量`是等价的。Go基本类型中的`复合类型`就是类型字面量。所以未命名类型、类型字面量和Go语言基本类型中的复合类型三者等价。

2. 通常所说的Go语言基本类型中的`简单类型`就是这20个预声明类型，他们都属于`命名类型`。

3. 预声明类型是命名类型的一种，另一类命名类型是`自定义类型`（用户type定义的类型）。

### 底层类型

所有类型都有一个underlying type（底层类型）。

1. 预声明类型（Pre-declared types）和类型字面量（type literals）的底层类型使它们自身。

1. 自定义类型 type newtype oldtype 中 newtype 的底层类型是逐层递归向下查找的，直到查到的oldtype是预声明类型（Pre-declared types）或类型字面量（type literals）为止。

```go
type T1 string
type T2 T1
type T3 []string
type T4 T3
type T5 []T1
type T6 T5
```

T1 和 T2 的底层都是string，T3 和 T4 的底层类型都是`[]string`类型，T6和T5的底层类型都是`[]T1`。
特别注意这里的T6、T5与T4、T3的底层类型是不一样的，一个是`[]T1`，另一个是`[]string`。

底层类型在类型赋值和类型强制转换时会使用。

### 类型相同和类型赋值

**类型相同**

Go是强类型的语言，编译器在编译时会进行严格的类型校验。两个命名类型是否相同参考如下：

- 两个命名类型相同的条件是两个类型声明的语句完全相同。
- 命名类型和未命名类型永远不相同。
- 两个未命名类型相同的条件是它们的类型声明字面量的结构相同，并且内部元素的类型相同。
- 通过类型别名语句声明的两个类型相同。

Go 1.9 引入了类型别名语法 type T1 = T2，T1 的类型完全和T2一样。引入别名主要有一下原因:

- 为了解决新旧包的迁移兼容问题。
- Go的按包进行隔离的机制不太精细，有时我们需要将大包划分为几个小包进行开发，但需要在大包里面暴露全部的类型给使用者。
- 解决新旧类型的迁移问题，新类型显示旧类型的别名，后续的软件都基于新类型编程，在合适的时间将新类型升级为和旧类型不兼容，常用于软件的柔性升级。

**类型可直接赋值**

不同类型的变量之间一般是不能直接相互赋值的，除非满足一定的条件。

类型为T1的变量a可以赋值给类型为T2的变量b，称为类型T1可以赋值给类型T2，伪代码表述如下：

```go
// a 是类型为T1的变量，或者a本身就是一个字面常量或nil
// 如果如下语句可以执行，则称之为类型T1可以赋值给类型T2
var b T2 = a
```

a 可以赋值给变量b必须要满足如下条件中的一个：
- T1和T2的类型相同
- T1和T2具有相同的底层类型，并且T1和T2里面至少有一个是未命名类型。
- T2是接口类型，T1是具体类型，T1的方法集是T2方法集的超集。
- T1和T2都是通道类型，它们拥有相同的元素类型，并且T1和T2中至少有一个是未命名类型。
- a是预声明标识符nil，T2是pointer、function、slice、map、channel、interface类型中的一个。
- a是一个字面常量值，可以用来表示类型的T值。

example:
```go
package main

import "fmt"

// Map doc false
type Map map[string]string

// Print doc false
func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

// 只要底层类型是slice、map等支持range的类型字面量，新类型仍然可以使用range迭代

// Print doc false
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"

	// mp 与 ma 有相同的底层类型 map[string]string, 并且mp是未命名类型变量
	// 所以 mp 可以直接赋值给 ma
	var ma Map = mp

	// 不能赋值，如下语句不能通过编译：
	// var im iMap = ma
	// im.Print()
	// im 与 ma 虽然有相同的底层类型 map[string]string, 但它们中没有一个是未命名类型

	ma.Print()

	// Map 实现了 Print()，所以其可以赋值给接口类型变量
	var i interface {
		Print()
	} = ma

	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()
}
```

### 类型强制转换

由于Go是强类型语言，如果不满足自动转换的条件，则必须进行强制类型转换。
转换格式：
```go
var a T = (T)(b)
```

非常量类型的变量x可以强制转化并传递给类型T，需要满足如下任一条件：
- x可以直接赋值给T类型变量

- x的类型和T具有相同的底层类型：
> example:
```go
package main

import "fmt"

// Map doc false
type Map map[string]string

// Print doc false
func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

// 只要底层类型是 slice、map 等支持 range 的类型字面量，新类型仍然可以使用 range 迭代
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	// mp 与 ma 有相同的底层类型 map[string]string，并且mp是未命名类型
	var ma Map = mp

	// im 与 ma 虽然有相同的底层类型，但是二者中没有一个是字面量类型，不能直接赋值，
	// 可以强制进行类型转换
	// var im iMap = ma
	var im iMap
	im = (iMap)(ma)

	ma.Print()
	im.Print()
}
```

- x的类型和T都是未命名的指针类型，并且指针指向的类型具有相同的底层类型。
- x的类型和T都是整数，或者都是浮点型。
- x的类型和T都是复数类型。
- x是整数值或`[]byte`类型的值，T是string类型。
- x是一个字符串，T是`[]byte`或`[]rune`

字符串和字节切片之间的转换最常见，示例如下：
[string_to_byte_slice.go](string_to_byte_slice.go)

注意：
1. 数值类型和string类型之间的相互转换可能造成值部分丢失。其他的转换仅是类型转换，不会造成值的改变。string和数字之间的转换可使用标准库strconv。
2. Go语言没有语言机制支持指针和integer之间的直接转换，可以使用标准库中的unsafe包进行处理。


## 类型方法

为类型增加方法是Go语言实现面向对象编程的基础。

### 自定义类型

用户自定义类型使用关键字type，语法格式为：
```go
type newtype oldtype
```
其中oldtype可以是自定义类型、预声明类型、未命名类型中的任意一种。
新类型与旧类型具有相同的底层类型，并且都继承了底层类型的操作集合。
这里的操作不是方法，比如底层类型是map，支持range迭代访问，则新类型也可以使用range迭代。

除此之外，`newtype`和`oldtype`是两个完全不同的类型，`newtype`不会继承`oldtype`的方法。

无论`oldtype`是什么类型，使用type声明的新类型都是一种命名类型，也就是说自定义类型一定是命名类型。

#### 自定义struct类型

struct 类型是Go语言自定义类型的普遍的形式，是Go语言类型拓展的基石，也是Go语言面向对象承载的基础。

如果使用type语句声明，则struct类型定义的新类型也是命名类型。Eg: [struct_type.go](struct_type.go)

#### struct 初始化

以Person结构为例

```go
type Person struct {
	name string
	age int
}
```

1. 按照字段顺序进行初始化
```go
a := Person{"andes", 18}
```
这不是一种推荐的方法，一旦结构增加字段，则不得不修改顺序初始化的语句。

2. 指定字段名进行初始化
```go
a := Person{name: "andes", age: 18}
```

3. 使用new创建内置函数，字段默认初始化为其类型的零值，返回值是指向结构的指针。
```go
p := new(Person)
// 此时name为""，age是0
```
这种方法不常用，一般结构体不会被初始化零值

4. 一次初始化一个字段
```go
p := Person{}

p.name = "andes"
p.age = 18
```
这种方法不常用，这是一种结构化的编程思维，没有封装，违背了struct本身抽象封装的理念。

5. 使用构造函数进行初始化

这是推荐的一种方法，当结构发生变化时，构造函数可以屏蔽细节。下面是标准库errors的构造函数New：
```go
// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}
```

#### 结构字段的特点

结构的字段可以是任意的类型，基本类型、接口类型、指针类型、函数类型都可以作为struct的字段。
结构字段的类型名必须唯一，struct字段类型可以是普通类型，也可以是指针。另外，结构支持内嵌自身的指针，
这也是实现树形和链表等复杂数据结构的基础：
```go
// 标准库 container/list

type Element struct {
	// 指向自身类型的指针
	next, prev *Element
	list *List
	Value interface{}
}
```

#### 匿名字段

在定义struct过程中，字段只给出字段类型，没有给出字段名，则称这样的字段为”匿名字段“。
被匿名嵌入的字段必须是命名类型或命名类型指针，类型字面量不能作为匿名字段使用。

匿名字段的字段名默认就是类型名，如果匿名字段是指针类型，则默认的字段名就是指针指向的类型名。
但一个结构体里面不能同时存在某一类型及其指针类型的匿名字段，原因是二者的字段名相等。
如果嵌入的字段来自其他包，则需要加上包名，并且必须是其他包可以导出的类型。
```go
type File struct {
	*file // os specific
}
```

#### 自定义接口类型

自定义接口类型同样使用type关键字声明。
例如：
```go
// interface{} 是接口字面量类型标识，所以i是非命名类型变量
var i interface{}

// Reader 是自定义接口类型，属于命名类型
type Reader interface {
	Read(p []byte)(n int, err error)
}
```

### 方法

Go语言的类型方法是一种对类型行为的封装。
其显式的将对象示例或指针作为函数的第一个参数，并且参数名可以自己指定。
语法规则：

```go
// 类型方法接受者是值类型
func (t TypeName)MethodName(ParamList)(ReturnList) {
	// method body
}

// 类型方法接受者是指针
func (t *TypeName)MethodName(ParamList)(ReturnList) {
	// method body
}
```
说明：
- t是接收者，可以自由指定名称。
- TypeName为命名类型的类型名。
- MethodName为方法名，是一个自定义标识符。
- ParamList是形参列表。
- ReturnList是返回值列表。

Go语言的类型方法本质上就是一个函数，没有使用隐式的指针。
[type_method.go](type_method.go)

类型方法有如下特点：
- 可以为命名类型增加方法（除了接口），非命名类型不能自定义方法。
- 为类型增加方法有一个限制，就是*方法的定义必须和类型的定义在同一个包中*。
- 方法的命名空间的可见性和变量一样，大写开头的方法可以在包外被访问，否则只能包内可见。
- 使用type定义的自定义类型是一个新类型，*新类型不能调用原有类型的方法*，但是底层类型支持的运算可以被新类型继承。
[type_method2.go](type_method2.go)

## 方法调用

### 一般调用

类型方法的一般调用方式：
```go
TypeInstanceName.MethodName(ParamList)
```
[type_method3.go](type_method3.go)

### 方法值(method value)

变量x的静态类型是T，M是类型T的一个方法，x.M被称为方法值，x.M是一个函数类型变量，可以赋值给其他变量，并像普通的函数名一样使用。
```go
f := x.M
f(args...)
```
等价于
```go
x.M(arg...)
```
方法值其实是一个带有闭包的函数变量，其底层实现原理和带有闭包的匿名函数类似，接收值被隐式地绑定到方法值（method value）的闭包环境中。
Eg: [type_method4.go](type_method4.go)

### 方法表达式

方法表达式相当于提供一种语法将类型方法调用显式地转换为函数调用，接收者（receiver）必须显式地传递进去。下面定义一个类型T，增加两个方法，
方法Get的接收者为T，方法Set的接收者类型为*T。

```go
type T struct {
	a int
}

func (t *T) Set(i int) {
	t.a = i
}

func (t T) Get() int {
	return t.a
}

func (t *T) Print() {
	fmt.Printf("%p, %v, %d\n", t, t, t.a)
}
```
表达式T.Get和(*T).Set被称为方法表达式，方法表达式可以看做函数名，只不过这个函数的首个参数是接收者的实例或指针。
```go
// 如下方法表达式调用都是等价的
t := T{a:1}

// 普通方法调用
t.Get(t)

// 方法表达式调用
(T).Get(t)

// 方法表达式调用
f1 := T.Get
f1(t)

// 方法表达式调用
f2 := (T).Get
f2(t)

//如下方法表达式调用都是等价的
(*T).Set(&t, 1)
f3 := (*T).Set
f3(&t, 1)
```

### 方法集(method set)

命名类型方法接收者有两种类型，一个是值类型，另一个是指针类型，
这个和函数是一样的，前者的形参是值类型，后者的形参是指针类型。
[method_set.go](method_set.go)

新类型Int虽然不能继承int的方法，但底层类型支持的操作可以被上层类型继承。这是Go系统类型的一个特点。

接收者是Int类型的方法集
```go
func (i Int) Print()
func (i Int) Max(b Int) Int
```

接收者是*Int类型的方法集
```go
func (i *Int) Set(a Int)
```

将接收者（receiver）为值类型`T`的方法的集合记录为`S`，将接收者（receiver）为指针类型`*T`的方法的集合统称为`*S`:
1. `T`类型的方法集是`S`
2. `*T`类型的方法集是`S`和`*S`

在直接使用类型实例调用类型的方法时，无论值类型变量还是指针类型变量，都可以调用类型的所有方法，原因是编译器在编译期间
能够识别出这种调用关系，做了自动转换。

比如 `a.Set()` 使用值类型实例调用指针接收者方法，编译器会自动将其转换为`(&a).Set()`,`(&a).Print()`使用指针类型实例调
用值类型接收者方法，编译器自动将其转化为`a.Print()`。

### 值调用和表达式调用的方法集

具体类型实例变量直接调用其方法时，编译器会为调用的方法进行自动转换，即使接受者是指针方法，仍然可以使用值类型变量进行调用。

下面讨论在以下两种情况下编译器是否会进行方法的自动转换：

1. 通过类型字面量显式的进行值调用和表达式调用，可以看到在这种情况下编译器不会做自动转换，会进行严格的方法集检查。
```go
type Data struct { }

func (Data) TestValue() {}
func (*Data) TestPointer() {}

// 这种字面量显式调用，无论值调用，还是表达式调用，
// 编译器都不会进行方法集的自动转换，编译器会严格校验方法集。

// *Data方法集是TestPointer和TestValue
// Data 方法集只有TestValue

(*Data)(&struct{}{}).TestPointer() // 显式的调用
(*Data)(&struct{}{}).TestValue() // 显式的调用

(Data)(struct{}{}).TestValue() // method value
Data.TestValue(struct{}{}) // method expression

// 如下调用因为方法集合不匹配而失败
// Data.TestPointer(struct{}{}) // type Data has no method TestPointer
// (Data)(struct{}{}).TestPointer() // cannot call pointer method on Data(struct{} literal)
```

2. 通过类型变量进行值调用和表达式调用，在这种情况下，使用`值调用`（method value）方式调用时编译器`会进行自动转换`，
使用`表达式调用`（method expression）方式调用时编译器`不会进行转换`，会进行严格的方法集检查。
[method_value.go](method_value.go)
```go
type Data struct {}

func (Data) TestValue() {}
func (*Data) TestPointer() {}

// 声明一个类型变量a
var a Data = struct{}{}

// 表达式调用编译器不会进行自动转换
Data.TestValue(a)
// Data.TestValue(&a)
(*Data).TestPointer(&a)
// Data.TestPointer(&a)

// 值调用编译器会进行自动转换
f := a.TestValue
f()

y := (&a).TestValue // 编译器帮助转换a.TestValue
y()

g := a.TestPointer // 会转换为(&a).TestPointer
g()

x := (&a).TestPointer
x()
```

## 组合和方法集

### 组合

命名结构类型可以嵌套其他的命名类型的字段，外层的结构类型是可以调用嵌入字段类型的方法，
这种调用既可以是显式的也可以是隐式的。这就是Go的继承，准确的说是组合。Go是没有继承语义的。

#### 内嵌字段的初始化和访问

struct的字段访问使用点操作符”.“，struct的字段可以嵌套很多层，只要内嵌的字段是唯一的即可，
不需要使用全路径进行访问。
[combination.go](combination.go)

在struct多层嵌套中，不同嵌套层次可以有相同的字段，此时最好使用完全路径进行访问和初始化。

#### 内嵌字段的方法调用

struct类型方法调用也使用点操作符。
不同嵌套层次的字段可以有相同的方法，外层变量调用内嵌字段的方法时也可以像嵌套字段的访问一样使用简化模式。
如果外层字段和内层字段有相同的方法，则使用简化模式访问外层的方法会覆盖内层的方法。

这个特性类似面向对象编程中子类方法覆盖父类的同名方法。

[example.go](example.go)
不推荐在多层的struct类型中内嵌多个同名的字段；但是并不反对struct定义和内嵌字段同名方法的用法，
因为这提供了一种编程技术，使得struct能够重写内嵌字段的方法，提供面向对象编程中子类覆盖父类的同名方法的功能。

### 组合的方法集

组合结构的方法集有如下规则：
- 若类型S包含匿名字段T，则S的方法集包含T的方法集。
- 若类型S包含匿名字段`*T`，则S的方法集包含T和`*T`方法集。
- 不管类型S中嵌入的匿名字段是T还是`*T`，`*S`方法集总是包含T和`*T`方法集。

下面的示例使用方法表达式的调用方式，阻止Go编译器对方法调用进行自动转换，更清楚地理解方法集的规约。
[example2.go](example2.go)

编译器的自动转换仅适用于直接通过类型实例调用方法时才有效，类型实例传递给接口时，编译器不会进行自动转换,
而是会进行严格的方法集校验。

## 函数类型

#### 函数字面量类型

函数字面量类型的语法表达式格式是func(InputTypeList)OutputTypeList，可以看出”有名函数“和”匿名函数“的类型都属于函数字面量类型。

#### 函数命名类型

类似命名变量，也可以使用type定义函数命名类型，语法上支持这么做，但是很少这么用。

```go
type NewFuncType OldFuncType
```

#### 函数签名

所谓函数签名就是函数字面量类型，但是不包括函数名。


#### 函数声明

Go语言没有C语言中函数声明的语义，准确地说，Go代码调用Go编写的函数不需要声明，可以直接调用，但Go调用汇编语言编写的
函数还是要使用函数声明语句:

```go
// 函数声明 = 函数名 + 函数签名

// 函数签名
func (InputTypeList)OutputTypeList

// 函数声明
func FuncName (InputTypeList)OutputTypeList
```

[functions.go](functions.go)
