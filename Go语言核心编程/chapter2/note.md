# 第2章 函数

近几年函数式语言因其数值不变性在高并发场景备受青睐。

Go不是一门纯函数式编程语言，但是函数在Go中是第一公民：
- 函数是一种类型，函数类型变量可以像其他类型变量一样使用，可以作为其他函数的参数或返回值，也可以直接调用。
- 函数支持多值返回。
- 支持闭包。
- 函数支持可变参数。

## 基本概念

### 函数定义

- 函数是Go程序源代码的基本构造单位。
- 函数名遵循标识符的命名规则，首字母的大小写决定该函数在其他包的可见性。
- 函数的参数和返回值需要使用()包裹，如果只有一个返回值，而且使用的是非命名的参数，则返回参数的()可以省略。
```go
func funcName(param-list)(result-list) {
  function-body
}
```

**函数的特点**

1. 函数可以没有输入参数，也可以没有返回值。默认返回0
2. 多个相邻的相同类型的参数可以使用简写模式。
```go
func add(a, b int) int { // a int, b int 简写为 a, b int
  return a + b
}
```
3. 支持有名的返回值，参数名就相当于函数体内最外层的局部变量，命名返回值变量会被初始化为类型零值，最后的return可以不带参数名直接返回。
```go
// sum 相当于函数内的局部变量，被初始化为零
func add(a, b int) (sum int) {
  sum = a + b
  return // return sum的简写
  // sum := a + b // 如果是sum := a + b，则相当于新声明一个sum变量命名返回变量sum覆盖
  // return sum // 最后需要显式地调用return sum
}
```
4. 不支持默认值参数
5. 不支持函数重载
6. 不支持函数嵌套，严格地说是不支持命名函数的嵌套定义，但支持嵌套匿名函数。
```go
func add(a, b int) (sum int) {
  anonymous := func(x, y int) int {
    return x + y
  }

  return anonymous(a, b)
}
```

### 多值返回

Go函数支持多值返回，定义多值返回的返回参数列表时要使用()包裹，支持命名参数的返回。
```go
func swap(a, b int) (int, int) {
  return b, a
}
```
习惯用法，如果多值返回有错误类型，则一般将错误类型作为最后一个返回值。

### 实参到形参的传递

Go函数实参到形参的传递永远是拷贝。

```go
package main

import "fmt"

func chvalue(a int) int {
  a = a + 1
  return a
}

func chpointer(a *int) {
  *a = *a + 1
  return
}

func main() {
  a := 10
  chvalue(a) // 实参传递给形参是值拷贝
  fmt.Println(a)

  chpointer(&a) // 实参传递给形参仍然是值拷贝，只不过复制的是a的地址值
  fmt.Println(a)
}
```

### 不定参数

Go函数支持不定数目的形式参数，不定参数声明使用`param ...type`的语法格式。
- 所有不定参数类型必须是相同的。
- 不定参数必须是函数的最后一个参数。
- 不定参数名在函数体内相当于切片，对切片的操作同样适合对不定参数的操作。
```go
func sum(arr ...int) (sum int) {
  for _, v := range arr { // 此时arr就相当于切片，可以使用range访问
    sum += v
  }
  return
}
```
- 切片可以作为参数传递给不定参数，切片名后面加上“...”。
```go
func sum(arr ...int) (sum int) {
  for _, v := range arr {
    sum += v
  }
  return
}

func main() {
  slice := []int{1, 2, 3, 4}
  // array := [...]int{1, 2, 3, 4} // 数组不可以作为实参传递给不定参数的函数

  sum(slice...)
}
```

## 函数签名和匿名函数

### 函数签名

函数类型又叫函数签名。可以使用`fmt.Printf("%T", func)`来打印函数的类型。
可以使用type定义函数类型，函数类型变量可以作为函数的参数或返回值。

```go
package main

import "fmt"

func add(a, b int) int {
  return a + b
}

func sub(a, b int) int {
  return a - b
}

type Op func(int, int) int // 定义一个函数类型

func do(f Op, a, b, int) int { // 定义一个函数，第一个参数是函数类型Op
  return f(a, b) // 函数类型变量可以直接用来进行函数调用
}

func main() {
  a := do(add, 1, 2)
  fmt.Println(a)
  s := do(sub, 1, 2)
  fmt.Println(s)
}
```

函数类型和map、slice、chan一样，实际函数类型变量和函数名都可以当作指针变量，该指针指向函数代码的开始位置，通常说函数类型变量是一种引用类型，未初始化的函数类型的变量的默认值是nil。

Go中函数是“第一公民”，有名函数的函数名可以看作函数类型的常量，可直接使用函数名调用该函数，也可以直接赋值给函数类型变量，后续通过该变量来调用该函数。
```go
package main

func sum(a, b int) int {
  return a + b
}

func main() {
  sum(3, 4)
  f := sum
  f(1, 2)
}
```

### 匿名函数

Go提供两种函数；匿名函数可以看作函数字面量，所有直接使用该函数类型变量的地方都可以由匿名函数代替，匿名函数可以直接赋值给函数变量，可以当做实参，也可以作为返回值。

```go
package main

import "fmt"

var sum = func(a, b int) int {
  return a + b
}

func doinput(f func(int, int) int, a, b int) int {
  return f(a, b)
}

func wrap(op string) func(int, int) int {
  switch op {
    case "add":
      return func(a, b int) int {
        return a + b
      }

    case "sub":
      return func(a, b int) int {
        return a - b
      }
    default:
      return nil
  }
}

func main() {
// 匿名函数直接被调用

  defer func() {
    if err := recover(); err != nil {
      fmt.Println(err)
    }
  }()

  sum(1, 2)

  // 匿名函数作为实参
  dopoint(func(x, y int) int {
    return x + y
  }, 1, 2)

  opFunc := wrap("add")
  re := opFun(2, 3)

  fmt.Printf("%\n", re)
}
```

## defer

Go函数里提供了defer关键字，可以注册多个延迟调用，这些调用可以采用堆栈调用。先注册的后调用。
常用于保证一些资源最终一定能够得到回收和释放。

```go
package main

func main() {
  defer func() {
    println("first")
  }()

  defer func() {
    println("second")
  }()

  println("function body")
}
```
主动调用`os.Exit(int)`退出程序时，defer函数不会被执行。

```go
package main

import "os"

func main() {
  defer func() {
    println("deger")
  }()

  println("func body")

  os.Exit(1)
}
```

defer 的好处是可以在一定程度上避免资源泄露，特别是在有很多return语句，有多个资源需要关闭的场景，很容易漏掉资源的关闭操作。

```go
func CopyFile(dst, stc string) (w int64, err error) {
  src, err := os.Open(src)
  if err != nil {
    return
  }
  defer src.Close()

  dst, err := os.Create(dst)
  if err != nil {
    return
  }

  defer dst.Close()

  w, err = io.Copy(dst, src)

  return
}
```

## 闭包

### 概念

闭包是由函数及其相关引用环境组合而成的实体，一般通过在匿名函数中引用外部函数的局部变量或包全局变量组成。

> 闭包 = 函数 + 引用环境

闭包对闭包外的环境引入是直接引用，编译器检测到闭包，会将闭包引用的外部变量分配到堆上。
如果函数返回的闭包引用了该函数的局部变量（参数或函数内部变量）：
1. 多次调用该函数，返回的多个闭包所引用的外部变量是多个副本，原因是每次调用函数都会为局部变量分配內存。
2. 用一个闭包函数多次，如果该闭包修改了其引用的外部变量，则每一次调用该闭包对该外部变量都有影响，因为闭包函数共享外部引用。

```go
package main

func fa(a int) func(i int) int {
  return func(i int) int {
    println("address a:", &a, " value:", a)
    a = a + i
    return a
  }
}

func main() {
  f := fa(1) // f 引用的外部的闭包环境包括本次函数调用的形参a的值1
  g := fa(1) // g 引用的外部的闭包环境包括本此函数调用的形参a的值1

  // 此时f、g引用的闭包环境中的a的值并不是同一个，而是两次函数调用产生的副本
  println(f(1))
  // 多次调用f引用的是同一个副本a
  println(f(1))

  // g 中的a的值仍然是1
  println(g(1))
  println(g(1))
}
```

如果一个函数调用返回的闭包引用修改了全局变量，则每次调用都会影响全局变量。

### 闭包的价值

闭包最初的目的是减少全局变量，在函数调用的过程中隐式传递共享变量。
*对象是附有行为的数据，而闭包是附有数据的行为*



## panic和recover

### 基本概念

panic 和 recover 的函数签名如下：
```go
panic(i interface{})
recover() interface{}
```

引发panic有两种情况，一种是程序主动调用panic函数，另一种是程序产生运行时错误，由运行时检测并抛出。

发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，知道被recover捕获或运行到最外层函数而退出。

panic 不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。

recover()用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，但是recover()只有在defer后面的`函数体内`被直接调用才能捕获panic终止异常，否则返回nil，异常继续向外传递。

```go
// 这个会捕获失败
defer recover()

// 这个会捕获失败
defer fmt.Println(recover())

// 这个嵌套两层也会捕获失败
defer func() {
  func() {
    println("defer inner")
    recover() // 无效
  }()
}()

// 如下场景会捕获成功
defer func() {
  println("defer inner")
  recover()
}()

func except() {
  recover()
}

func test() {
  defer except()
  panic("test panic")
}
```

可以有连续多个panic被抛出，连续多个panic的场景只能出现在延迟调用里面，否则不会出现多个panic被抛出的场景。但只有最后一次panic能被捕获。

```go
package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 只有最后一次panic调用能够被捕获
	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")
}
```

## 错误处理

**TODO 后续回来重新看**
***

### error

Go语言内置错误接口类型error。
Go语言典型的错误处理方式是将error作为函数最后一个返回值。

### 错误和异常

Go 程序需要处理的错误可以分为两类：
1. 运行时错误。
2. 程序逻辑错误。
