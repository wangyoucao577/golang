# [Go] my_gopl_test  
Sample and exercise codes from learning  [The Go Programming Language](http://gopl.io). 

## 实验平台
- Linux: `Ubuntu 16.04.4 LTS`
    - `Kernel 4.13.0-36-generic`
    - `go version go1.9.4 linux/amd64`


## 学习笔记
### 基本点
- `Go`是一门编译型语言，且仅支持静态编译链接（不支持动态链接）
- `Go`原生支持`unicode`
- `Go`的编译没有警告, 要么pass, 要么error
    - `Go`语言不允许 unused local variable, 否则会报编译错误
    - `import`但未被引用的包, 会导致编译错误
- `Go`的代码通过`package`组织(类似于其他语言的modules 或 libraries)
    - 一个`package`由位于单个目录下的一个或多个.go生成
    - 通常目录的名字即为`package`的名字
    - 而`package xxx`中的`xxx`为命名空间(引用名)
    - 按照惯例是导入路径的最后一段与命名空间一致
- 每个源文件都应
    - 首先以 `package xxx` 开始，以定义此文件属于哪个package.
    - 然后`import xxx` 导入所需要链接的`package`(`import`必须在`package`之后)
    - 再然后才是此文件中的代码实现
- `package main`，以及 `func main`
    - `package main`定义了独立的可执行程序
    - `func main`则定义了程序的入口函数
- 注释： `//`
- `Go`中 函数和包级别(package level entities)的变量/函数可以任意顺序声明, 并不影响其调用
- 变量的几种声明/初始化形式(声明即初始化)
    - `var s1, s2 string` (声明2个`string`变量, 隐式初始化为`""`字符串. 若类型为`int`, 则隐式初始化为`0`)
	- `s1,s2 := "",""` (声明2个变量s1 s2, 以空字符串`""`初始化, 于是即推导出其为`string`类型. 此种方式术语叫做 short variable declaration)
	- `var s1 = ""` (冗余的 `var`, 使用的较少)
	- `var s1 string = ""` (冗余的`var` 和`string`, 用的也比较少)
- Go中的循环语句仅有for一种, 其有几种形式
    - `for initialization, condition, post { `
        - 此种形式下, initialization/condition/post 均可以省略, 左大括号"{" 必须在 post 的同一行
	- `for index, arg := range os.Args[1:] {`
        - 此种区间遍历形式, 提供 索引 和 值 两个参数. 若不需要其中某个, 经常是不需要 index, 可以用 blank identifier "_" 即下划线来代替(Go语言不允许unused local variable). 如 `for _, arg := range os.Args[1:] {`






### Go的代码组织 
(摘自[How to Write Go Code](https://golang.org/doc/code.html)的Code origanization章节原文, 非常好的入门文档, 强烈推荐!!)  
- Overview
    - Go programmers typically keep all their Go code in a single workspace.
    - A workspace contains many version control repositories (managed by Git, for example).
    - Each repository contains one or more packages.
    - Each package consists of one or more Go source files in a single directory.
    - The path to a package's directory determines its import path.
    - Note that this differs from other programming environments in which every project has a separate workspace and workspaces are closely tied to version control repositories.

- Workspaces
    - A workspace is a directory hierarchy with three directories at its root:
        - src contains Go source files,
        - pkg contains package objects, and
        - bin contains executable commands.
    - The go tool builds source packages and installs the resulting binaries to the pkg and bin directories.

    - The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.

    - To give you an idea of how a workspace looks in practice, here's an example:
```
bin/    
    hello                          # command executable    
    outyet                         # command executable    
pkg/    
    linux_amd64/    
        github.com/golang/example/    
            stringutil.a           # package object    
src/    
    github.com/golang/example/    
        .git/                      # Git repository     metadata    
	hello/    
	    hello.go               # command source    
	outyet/    
	    main.go                # command source    
	    main_test.go           # test source    
	stringutil/    
	    reverse.go             # package source    
	    reverse_test.go        # test source    
    golang.org/x/image/    
        .git/                      # Git repository     metadata    
	bmp/    
	    reader.go              # package source    
	    writer.go              # package source    
    ... (many more repositories and packages omitted) ...   
```    




### Go的常用工具命令    
Go提供了一系列的工具命令，都可以通过一个单独的go命令调用    
- go run：编译一个或多个 .go, 链接库文件, 并运行最终生成的可执行文件 (不会保留可执行文件)   
- go build: 编译一个或多个 .go, 链接库文件, 生成可执行程序或package   
- go install: 编译一个或多个 .go, 链接库文件, 生成可执行程序或package, 并将其对应的安装到 bin/pkg 目录下供执行或其他程序链接   
- gofmt: 格式化源代码，**强制无参数的命令来统一go的代码格式**, 默认行为为将diff的内容写到stdout，而要直接格式化源文件本身的话，加上 `-w` 选项   
    - `gofmt -l -w .`   
- goimports: 根据代码需要自动地添加或删除import   
- go doc: 在cmd中看go的文档   


### 细节与杂项
- Go中的区间索引：
    - 左闭右开原则. 即区间包括第一个索引元素, 不包括最后一个。（比如 a = [1, 2, 3, 4, 5], a[0:3]=[1,2,3], 即包含左边第一个元素a[0], 但不包含右侧的索引元素a[3]. ）
    - 区间索引的左、右参数分别可以省略。左参数省略则默认为0, 右参数省略则为len(a)
- 名字的作用域
    - 函数内部定义的名字，只在函数内部有效
    - 函数外部定义的名字(包级名字), 在整个包的所有文件中都可以访问
        - 包级名字，若首字母大写(包括函数名和变量名)，那么就是导出的名字，即可以被外部的包访问
        - 包级名字，若首字母小写(包括函数名和变量名)，那么就是属于包内部的名字，可以在包的所有文件中访问
    - 包本身的名字，一般总是用小写
- 命名风格
    - 倾向于不要太长的名字
    - 倾向于驼峰命名法(优先大小写分隔，而不是下划线分隔)
- 四种声明
    - `var`: 变量声明
    - `const`: 常量声明
    - `type`: 类型声明
    - `func`: 函数声明
- 变量的默认零值    
(零值初始化机制可以确保每个声明的变量总是有一个良好定义的值, 因此在`Go`中不存在未初始化的变量)
    - `bool`: `false`
    - 数值型: `0`
    - `string`: `""`
    - 接口或引用类型(包括`slice/map/chan/func`):`nil`
    - 指针: `nil`  
- 简短变量声明(例如 `a:= 0`)
    - 简短变量声明`:=`是声明语句, 而`=`是赋值语句
    - 简短变量声明语句中必须至少要声明一个新的变量, 否则会编译失败(已声明过的变量在简短变量声明语句中等价于赋值)
- 指针
    - `C`风格, 即通过`&`取地址, 通过`*`取值, 类型为`*T`
- 变量的生命周期与内存分配
    - `Go`中作用域的概念与`C/C++`中的作用域的概念不同
        - `Go`中的声明语句的作用域是指源代码中可以有效使用这个名字的范围
        - `Go`中的作用域是编译时概念, 而生命周期是运行时概念
            - 前者是变量名字的可见范围, 后者是实际变量的可引用时间段
        - 注意`for/if/switch`等的隐式词法域
            - 比如`if`条件中声明的变量, 在`else block`中也可见
    - 变量的回收由`Go`垃圾回收器负责, 而是否可以垃圾回收的唯一标志为变量是否仍然可达
        - 所以局部变量的地址返回后也依然有效, 因为依然可达
    - `new`在`Go`中只是一个内建的函数
        - `new`函数返回的是指针
        - 通过`new`函数创建变量和直接声明的变量没有什么本质区别(后者需要一个临时变量, 然后才能取地址)
        - `new`和变量是在栈上分配还是堆上分配没有任何关系
        - 通常情况下`new`的使用较少
    - 包级变量的生命周期和整个程序的运行一致
    - 局部变量的声明周期, 从声明开始, 到不可达结束
    - 变量的内存分配
        - 由`Go`编译器自动决定是在栈上分配还是堆上分配
        - 对于局部变量来说
            - 从函数中逃逸的变量, 必须在堆上分配
            - 不从函数中逃逸的变量, 由编译器自动决定在栈上分配还是堆上分配
- 命令行参数处理
    - `Go`中一般通过`flag`包, 类似于`python`中的`argparse`
- 包的初始化
    - 依次初始化(会按照变量的初始化依赖顺序)
    - 包中复杂变量的初始化, 可以通过特殊的`init()`函数来进行(进入`main()`之前自动被调用)
    - 每个包只会被初始化一次, 不会重复初始化
    - `package main`最后被初始化
- 内置的复数类型(`complex64/complex128`)
    - 用于构建复数, 其中`real()`和`imag()`函数分别返回复数的实部与虚部
- 字符串
    - 一个字符串是一个不可改变的字节序列(只读的)
    - `rune`: 对应于`utf-32`编码, 由于定长, 方便于索引
    - `Go`中`range`循环时, **会自动隐式解码`utf-8`, 故索引更新的步长将会超过1个字节, 应特别注意!!**
    - 字符串字面值
        - 通常用双引号`"内容"`来表达
        - 原生的字面值形式用`` `内容` ``来表达, 内部所有的字符都会字面解释而不会转义(会忽略回车符)
    - 标准库中常用的几个字符串处理的包
        - `bytes`
        - `strings`
        - `strconv`
        - `unicode`
- 常量
    - 常量表达式的值在编译期计算, 而不是在运行期. 
    - 常量间的所有算数运算、逻辑运算和比较运算的结果也都是常量, 对常量的类型转换操作或以下函数调用都是返回常量结果:
        - `len/cap/real/imag/complex/unsafe.Sizeof`
    - 常量的初始化
        - 批量声明时, 除了第一个以外其他的常量的右边的初始化表达式都可以省略, 即沿用上一个的初始化
        - 可通过`itoa`进行批量按一定规则初始化一堆常量(有点类似于其他语言中的`enum`)
    - `Go`中的常量可以无类型
        - 即若是没有显式明确类型, 则编译器为这些常量提供比基础类型更高精度的算术运算, 可以认为至少有256bit的运算精度
        - 有六种未明确类型的常量类型
            - 无类型的布尔型
            - 无类型的整数
            - 无类型的字符
            - 无类型的浮点数
            - 无类型的复数
            - 无类型的字符串 
- 数组
    - 数组和结构都是有固定内存大小的数据结构, 相比之下`slice`和`map`则是动态内存大小的数据结构
    - 数组的初始化
        - 默认情况下, 数组的每个元素都被初始化为元素类型对应的零值
        - 可以使用数组字面量的形式来进行初始化
        - 可以用`...`来初始化数组的长度, 即根据数组字面量初始化的数量来推导数组长度, 如`s := [...]int{1, 2, 3}`
        - 数组的长度是数组类型的组成部分, 故不同长度的数组可以认为是不同的类型
        - 数组的长度必须是常量表达式, 编译时确定
        - 也可以使用索引+对应值的形式来进行数组的初始化
    - `Go`中数组作为函数传参时, 会进行值传递(拷贝整个数组), 而不是许多语言中的指针/引用传递. 
        - 可以显式的以数组的指针作为函数参数进行传递
    - `Go`与`C`中数组的主要区别
        - `Go`中的数组是值类型, 将一个数组赋值给另一个数组, 会拷贝所有元素
        - `Go`中的函数传递数组参数时, 也是拷贝整个数组传递; 而`C`中传递的是指向数组的指针
        - `Go`中数组长度为数组类型的一部分, 声明即不可修改. 不同长度的数组可认为类型不同. 
- Slice
    - `slice`与数组
        - 主要区别为, 数组为定长的(编译时确定), 而`slice`是动态长度的
        - 语法上来讲, `[]T`代表`slice`, `[len]T`代表数组
        - 数组可以使用`==`, `!=`进行比较, `slice`则不行
            - `slice`仅支持与`nil`之间通过`==`, `!=`比较
            - 两个`slice`之间的比较, 需要通过循环进行深度比较(若元素类型为`byte`, 可以用`bytes.Equal`)
    - 理解`slice`
        - `slice`底层引用数组实现
        - 一个`slice`有三个部分构成: 指针、长度、容量(可以看做一个由指针、长度、容量组成的结构体来理解)
            - 指针: 指向当前`slice`的第一个元素的地址
                - 注意: 在底层数组中未必是第一个元素, 因为底层数组经常会是复用的
                - 传递`slice`时由于指针的存在, `slice`的底层数组元素内容是可修改的, 相当于传递一个`slice`的别名
            - 长度: `slice`中的元素个数, 内置的`len()`函数可以返回长度
            - 容量: `slice`能容纳的元素的总个数, 内置的`cap()`函数可以返回容量
        - 多个`slice`可以共享底层的数据, 并且引用的数组部分区间可能重叠
    - 内置的`make()`函数可以用于创建一个指定元素类型、指定长度和指定容量的`slice`
    - 内置的`append()`函数可以用于向一个`slice`追加元素
        - 对于任何可能修改`slice`的函数, 应将更新后的`slice`直接赋值给原`slice`, 以保证`len`, `cap`和底层数组元素的正确更新
    - `slice`的内存小技巧
        - 输入的`slice`和输出的`slice`共享一个底层数组结构, 从而避免了不必要的内存分配, 多用于过滤/合并`slice`中的元素
- Map
    - `map`的类型可写为**map[K]V**, 其中**K**和**V**分别对应于**Key**和**Value**
        - **Key**必须是支持`==`比较运算符的数据类型
        - **Value**类型则没有限制, 完全可以支持`slice/map`或自定义的聚合类型
    - `map`的创建
        - 可以使用内置的`make()`函数进行创建, 如`ages := make(map[string]int)`
        - 也可以使用字面值语法进行创建, 如`ages := map[string]int{}`
    - `map`的元素插入、查找、修改值、删除
        - 插入新元素或查找元素或修改已有元素的值: `ages[K]=V`
        - 删除元素: 使用内置的`delete()`函数
    - `map`的迭代
        - `map`的迭代顺序是不确定的, 并且不同的哈希函数实现可能导致不同的迭代顺序
            - 在实践中, 遍历的顺序是随机的, 每一次遍历的顺序都不一样. 这是故意的. 
            - 如果要按顺序遍历, 则必须显式地对**Key**进行排序. 比如用一个`slice`存储所有的**Key**, 排序后遍历`slice`取出**Key**, 再从`map`中取出对应的**Value**
    - `map`中的元素禁止取地址, 原因是`map`可能随着元素的增长而重新分配内存, 从而可能导致之前的地址无效
    - 和`slice`一样, `map`之间也不能进行`==`比较, 除了与`nil`比较
    - 由于`map`中的**Key**总是不同的, 必要时可以使用`map`模拟`set`的功能(`Go`没有内置`set`类型)
    - `map`上的大部分操作, 包括查找、删除、`len`和`range`循环都可以安全地工作在`nil`的`map`上, 它们的行为和一个空的`map`类似. 但是向一个`nil`的`map`插入新元素则会导致panic异常.  
- Struct 结构体
    - `struct`是一种聚合的数据类型. 
    - `struct`成员的定义顺序是有意义的, 交换了成员的顺序可以认为创建了不同的`struct`类型. 
    - `struct`成员的导出规则满足`Go`变量/类型导出的一般规则: 
        - 大写字母开头的成员名字为导出的
        - 小写字母开头的成员名字为未导出的
        - 一个结构体可能同时包含导出的和未导出的成员
    - 一个命名为`S`的`struct`类型将不能再包含`S`类型的成员, 因为一个聚合的值不能包含它自身. 但是`S`类型的结构体可以包含`*S`指针类型的成员, 这样我们就可以创建递归的数据结构了. 
    - 如果结构体没有任何成员的话，就是空结构体，写作`struct{}`, 大小为0, 也不包含任何信息. 但有时候也仍然是有价值的, 比如用`map`模拟`set`时用作**Value**以节约内存(虽然很少). 
    - 结构体字面值初始化的两种写法(如对于结构体`type Point struct{ X, Y int }`): 
        - 按照顺序初始化, 如 `p := Point{1, 2}`
        - 以成员名字和相应的值来初始化，如 `p := Point{ X: 1, Y: 2}`. 这种方式可以仅写明部分成员, 顺序也不影响. 
            - 注：对于未导出的成员, 在包外部这两种方式都不能使用, 因为无法访问未导出的成员. 
    - 如果结构体的成员都是可比较的, 那么结构体也是可比较的, 这样的话两个结构体将可以使用`==`或`!=`进行比较, 行为为比较两个结构体的每个成员. 
    - 匿名成员
        - 为`Go`的一个语言特性, 即声明一个成员时仅写明类型, 而不指明成员的名字. 实际上匿名成员还是有名字的, 名字就是类型名. 
        - 好处: 可以直接访问匿名成员的叶子属性(也同样可以通过显式的类型名字访问叶子属性)
- 函数(`func`)    
    - 函数的类型被称为函数的标识符. 如果两个函数形参列表和返回值列表都一一对应, 那么这两个函数被认为有相同的类型和标识符.    
    - 函数调用时, `Go`语言没有默认参数值.    
    - 实参总是值传递的, 因此函数的形参是实参的拷贝.   
    - 没有函数体的声明, 表示该函数不是以`Go`语言实现的(比如汇编实现).    
    - `Go`使用可变栈大小, 栈的大小按需增加(初始化时很小), 因此不会有栈溢出问题(尤其是递归时).   
    - 在`Go`中, 一个函数可以有多个返回值.    
    - 如果一个函数将所有的返回值都显式地命名变量名, 那么该函数的`return`语句可以省略操作数, 这称之为 bare return. bare return可能会使代码变得难以被理解, 不应过度使用.    
    - 在`Go`中, 函数被看做第一类值(first-class values):     
        - 函数像其他值一样, 拥有类型, 可以被赋值给其他变量, 传递给函数, 从函数返回.    
            - 对函数值(function value)的调用也类似于函数调用.    
            - 函数类型的零值是`nil`. 函数值可以与`nil`比较, 但函数值之间不可比较.   
        - 与`C`中的函数指针的概念非常类似, 对函数值的调用可看做`C`中对函数指针所指向的函数的调用.    
    - 拥有函数名的函数只能在包级语法块中被声明, 通过函数值字面量(function literal)也即匿名函数(anonymous function)可以绕过这一限制.    
    - 匿名函数的语法与普通函数的声明类似, 区别仅在于`func`关键字后没有函数名.    
        - 更为重要的是, 通过这种方式定义的函数可以访问完整的词法环境（lexical environment), 这意味着在函数中定义的内部函数可以引用该函数的变量.    
            - e.g. 函数值不仅仅是一串代码, 还记录了状态(变量引用).    
            - `Go`使用闭包(closures)技术实现函数值.    
        - 当匿名函数需要被递归调用时, 我们必须首先声明一个变量, 再把匿名函数赋值给这个变量.    
    - 可变参数函数    
        - 参数数量可变的函数称为可变参数函数, 典型的例子就是`fmt.Printf`.    
        - 在声明可变参数函数时, 需要在参数列表的最后一个参数类型之前加上省略符号`...`, 这表示该函数会接收任意数量的该类型参数. i.e. `func sum(vals ...int) int {}`    
        - 若原始参数已经是`slice`, 那么只需要在调用时最后一个参数后加上省略符`...`, 即可直接传递给可变参数函数. i.e. `vals := []int{1,2,3}; fmt.Println(sum(vals...))`    
    - `defer`机制    
        - 语法: 在调用普通函数或方法前加上关键字`defer`即可    
        - 当`defer`语句被执行时, 跟在`defer`后面的函数会被延迟执行, 直到包含该`defer`语句的函数执行完毕时, `defer`后的函数才会被执行. 不论包含`defer`语句的函数时通过`return`正常结束, 还是由于`panic`异常结束.    
        - 可以在一个函数中执行多条`defer`语句, 他们的执行顺序与声明顺序相反.   
        - 常应用于:    
            - 确保资源在退出函数时总是被关闭, 防止资源泄露. 如文件句柄、锁等.    
            - 记录进入和退出函数    
            - 在函数每次调用时输出参数和返回值, 甚至修改最终的返回值.    
        - 在循环中应用`defer`时要十分谨慎, 容易造成资源消耗过多甚至耗尽的风险    
        - `defer`机制有点类似于面向对象语言中局部`object`变量退出作用域时的自动析构, 可以实现类似的效果.    
- 方法(`method`)    
    - 属于某个特定类型的函数, 或者说绑定到某个特定类型的函数, 即是方法.    
    - 声明: 在函数声明时, 在其名字前放上一个变量, 即是一个方法. 这个附加的参数会将该函数附加到这种类型上, 相当于为这种类型定义了一个独占的方法.   
        - 示例函数: `func Distance(p, q Point) float64 {...} `    
        - 绑定到类型`Point`的相同功能的方法: `func (p Point) Distance(q Point) float64 {...} `    
    - 方法接收器(Receiver)的概念    
        - 早期的面向对象语言将调用一个方法称为"向一个对象发送消息"    
        - 大多语言中会使用`this`或`self`作为方法的接收器    
        - `Go`中可以任意选择接收器的名字, 通常使用其类型的第一个字母    
    - `Go`中我们可以很容易为一些简单的数值、字符串、`slice`、`map`等内置类型来定义一些附加行为.    
        - 方法可以被声明到任意类型, 只要不是一个指针或者一个`interface`    
        - 不需要像大多语言中那样派生出新的类型来, 而只需要为类型声明一个新的方法即可.    
    - 如果`method`需要更新`receiver`的内容, 或者`receiver`对象太大希望避免调用时的拷贝, 那么就可以用其指针而不是对象来声明`method`.    
        - 声明`method`的`receiver`该是指针还是非指针类型的原则:    
            - 对象本身是否特别大, 从而是否需要避免调用时的拷贝传值    
            - 是否需要通过此方法更新`receiver`的内容    
        - 不管`method`的`receiver`是指针类型还是非指针类型, 都是可以通过指针/非指针类型进行调用的, 编译器会帮我们做好取地址或解引用的转换.    
    - `nil`也是一个合法的`receiver`, 类似于给`func`传递了一个为`nil`的值    
    - 可以直接通过`struct`对象调用`struct`匿名内嵌成员的方法. 从实现的角度看, 可以理解为内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法.    
    - `method value`: 也即`method`对应的`function value`    
- 接口(`interface`)    
    - 接口类型是一种抽象的类型. 可以换个角度理解, 当你看到一个接口类型时, 你不知道它是什么, 唯一知道的就是可以通过它的方法来做什么. 也即方法接口的约定.    
    - 接口类型具体描述了一系列方法的集合, 一个实现了这些方法的具体类型就是这个接口类型的实例.    
    - 接口类型声明时, 方法的顺序变化没有影响, 也支持内嵌方式(类似结构的内嵌)声明, 唯一重要的就是这个集合里面的方法.     
    - 一个类型如果拥有一个接口需要的所有方法, 那么这个类型就实现了这个接口. 也即这个类型属于这个接口.    
    - 空接口类型`interface{}`对实现它的类型没有任何要求, 所以我们可以将任意一个值赋给空接口类型.    
    - 接口值(`interface value`)由一个具体的类型和一个此类型的值两个部分组成.     
        - 两部分皆为`nil`时 `interface value == nil`, 称为空接口值.    
        - 调用一个空接口值上的任何方法, 都会造成panic.    
        - 比较接口值或者包含了接口值的聚合类型时, 必须要意识到潜在的panic风险.    
        - 注意: 一个不包含任何值的`nil`接口值和一个刚好包含`nil`指针的接口值是不同的.    
    - 类型断言(`Type Assersion`)    
        - 可用于帮助判断接口值在运行时的实际类型    
        - 语法为`x.(T)`, 其中`x`为接口值, `T`为一个具体类型或接口类型. 返回值可以是一个或者两个
            - 一个返回值的情况, 如`f := x.(T)`, 若失败则抛出`panic`异常    
            - 两个返回值的情况, 如`f, ok := x.(T)`, 若失败则`ok`为`false`, 从而方便程序错误处理    
        - 可以通过`Type Assertion`询问行为(非常有价值的用法, 详见《The Go Programming Language》 Ch7.12)    
    - `Type Switch`
        - 通过`switch`以及`Type Assertion`来根据不同类型进行不同处理的方便写法, 本质上有点语法糖, 因为通过`if/else`加上`Type Assertion`完全可以实现, 但略显啰嗦. 用`switch`看起来优雅些.    
        - 书上把`Type Switch`翻译成了类型开关, 虽然是直译, 但看起来够够的, 非常的不 make sense, 还不如保留原文`Type Switch`     
- Concurrency: `Goroutine`, `Channel`, `sync`
    - `goroutine`+`channel` 支持术语为"顺序通信进程"(communicating sequential processes, 简称为CSP)的并发模型. 而更传统的并发模型为"多线程共享内存".     
    - 当一个程序启动时, `main`函数即在一个单独的`goroutine`中运行, 称为`main goroutine`.    
    - 通过`go`语句来创建新的`goroutine`, 语法上为普通的函数或方法调用前加上`go`关键字. e.g. `go f()`    
    - `channel`是一个通信机制, 它可以让一个`goroutine`通过它给另一个`goroutine`发送值信息. 每个`channel`都有一个特殊的类型, 也就是`channel`可发送的数据类型. 
        - e.g. `ch := make(chan int)` 即创建一个`channel`可发送`int`型值.    
        - e.g. `ch := make(chan string 3)` 即创建一个`channel`可发送`string`型值, 并最多缓存3个元素.    
    - 和`map`或`slice`类似, `channel`变量对应的也是一个底层数据结构的引用. 两个相同类型的`channel`可以使用`==`运算符比较.    
    - 一个`channel`有发送和接收两个主要操作, 都是通信行为. 语法为(`ch`为一个`channel`):    
        - 发送: `ch <- x`
        - 接收: `x <- ch` (不写`x`时, 如` <- ch` 则为丢弃接收的内容)
    - 使用`make()`创建一个`channel`, 使用`close()`关闭一个`channel`    
        - 通常不需要显式关闭
            - 首先, `close()`一个`channel`意义为不能再对此`channel`发送数据, 所以一般仅在需要告诉接收者`goroutine`, 要向`channel`发送的数据已经全部完成的时候才显式调用`close()`. 在接收`channel`数据的`goroutine`中可通过第二个返回值判断`channel`是否已经被关闭. e.g. ` x, ok := <- ch`  
            - 如上所述, 若要显式调用`close()`, 也仅应在发送的`goroutine`中调用     
            - 其次, `channel`不再被引用后会像普通变量一样自动被垃圾回收    
            - 试图重复`close()`一个`channel`或关闭一个`nil`的`channel`将导致`panic`异常
    - `Channel`默认行为为阻塞
        - 一个基于无缓存的`Channel`的发送操作将导致阻塞, 直到另一个`gorouting`在相同的`Channel`上执行接收操作. Vice Versa.    
            - 注: 当通过一个无缓存`Channel`发送数据时, 接收者收到数据发生在唤醒发送者`goroutine`之前("Happen Before").    
        - 带缓存的`Channel`, 则是在缓存用满后开始阻塞.    
    - 单向`channel`
        - 典型应用场景: 通常在当`channel`定义为函数参数时, 且其在函数中仅会被用于发送或仅被用于接收(发送`goroutine`和接收`goroutine`调用不同的函数)    
        - `out chan<- int`: `out`代表一个仅允许发送操作且支持的类型为`int`的`channel`
        - `in <-chan int`: `in`代表一个仅允许接收操作且支持的类型为`int`的`channel`
        - 单向`channel`的限制将在编译期检查. 对一个只接收的`channel`调用`close()`将会是编译错误.    
        - 任何双向`channel`向单向`channel`变量的赋值操作将会是隐式转换, 而反向并不能转换, 即不能将单向`channel`转换为双向.    
    - 带缓存的`channel`内部持有一个元素队列, 向`channel`的发送操作就是向内部缓存队列的尾部插入元素, 接收操作就是从缓存队列的头部取出元素.
    - 多个`goroutine`并发的向同一个`channel`发送数据, 或从同一个`channel`接收数据都是常见的用法.   
    - `goroutine`泄露: `goroutine`卡住而永远不会返回(如从一个不会再有数据的空的不带缓存的`channel`中接收), 称为`goroutine`泄露, 类似于线程卡死. 泄露的`goroutine`并不会被自动回收, 因此应确保每一个不再需要的`goroutine`能正常退出.    
    - 当在循环中使用`goroutine`进行并发处理时, 常用`sync.WaitGroup`来等待从而保证所有的`goroutine`都已退出, 防止`goroutine`泄露.    
    - `Golang`中可以基于`select`实现多路复用:    
        - `select`语法类似于`switch`, 有多个`case`和一个可选的`default`    
        - `select`会等待`case`中有能够执行的`case`时去执行. 当条件满足时, `select`才会去通信并执行`case`之后的语句. 如果多个`case`同时就绪, `select`会随机的选择一个通信并执行, 这样来保证每一个`channel`都有平等的被`select`的机会.    
        - `select`可以有`default`语句, 此时行为相当于变成了非阻塞的`select`, 所有其他`case`条件都不满足时, 会进入`default`分支执行.   
        - `select`本身仅一次行为, 常配合`for`使用.    
    - 退出`goroutine`的一个常用用法: 利用`close()`一个特定`channel`来广播退出消息, 在`goroutine`中查询这个`channel`是否已经被关闭从而决定继续执行还是退出    
    - `Go`中并发的口头禅："不要使用共享数据来通信, 使用通信来共享数据"
    - `sync`包中几种常用的互斥锁/方法: `sync.Mutex`, `sync.RWMutex`, `sync.Once`, `sync.WaitGroup`
    - `Go`的`runtime`和工具链为我们装备了一个复杂但好用的动态分析工具, 竞争检查器(the race detector), 帮助我们记录和报告所有已经发生的同步事件/数据竞争. 完整的同步事件集合参考[The Go Memory Model](https://golang.org/ref/mem).
    - `Goroutine`与`OS Thread`的主要区别:    
        - `OS Thread`比`Goroutine`有更大的栈内存开销
            - `OS Thread`通常有固定大小的栈内存(`linux`上貌似可以动态增长), 初始值也会相对较大(e.g. `2MB`)   
            - `Goroutine`会以一个很小的栈开始其生命周期, 一般只需要`2KB`. 并会根据需要动态地伸缩, 最大值可以有`1GB`.    
        - `OS Thread`比`Goroutine`有更大的调度开销
            - `OS Thread`由`OS`进行调度, 一般会依赖于硬件计时器的中断调用一个叫`scheduler`的内核函数. 线程调度切换时需要完整的上下文切换, 也就是说, 保存一个用户线程的状态到内存, 恢复另一个线程的到寄存器, 然后更新调度器的数据结构. 这个上下文切换回很慢. 
            - `Goroutine`则由`Go`的调度器在程序内部进行`m:n`调度(在`n`个操作系统线程上多工调度`m`个`Goroutine`), 并不依赖硬件计时器, 也不需要内核层面的上下文切换, 调度代价低得多.    
            - `Go`调度器通过`GOMAXPROCS`可以决定有多少个(`n`)操作系统线程同时执行`Go`的代码. 其默认值是运行机器上的`CPU`核心数. 在休眠中的或者在通信中被阻塞的`goroutine`是不需要一个对应的线程来做调度的. 在`I/O`中或系统调用中或调用非`Go`语言函数时, 是需要一个对应的操作系统线程的. 但是`GOMAXPROCS`并不需要将这几种情况计数在内.     
        - `OSThread`有明确的身份标识(`thread id`), 而`Goroutine`没有. 
            - 这一点是设计上故意而为之, 从而鼓励更为简单的模式.    
- 封装    
    - `Go`语言只有一种控制可见性的手段: 大写首字母的标识符会从定义它们的包中被导出, 小写字母的则不会. 这种基于名字的手段使得在`Go`语言中最小的封装单元是`package`.     
- 错误处理    
    - panic是来自被调用函数的信号, 表示发生了某个已知的bug.     
        - 一个良好的程序永远不应该发生panic异常.   
        - 有些错误只能在运行时检查, 如数组访问越界、空指针引用等，这些错误会引起panic异常.    
        - 一般而言, panic异常发生时, 程序会中断执行, 并立即执行在该`goroutine`中被延迟的函数(`defer`机制). 随后, 程序崩溃并输出日志信息. 日志信息包括`panic value`和函数调用的堆栈信息. 通常这些日志信息已经提供了足够的诊断依据.    
        - `Go`的panic机制中, `defer`延迟函数的调用在堆栈释放之前.    
        - 直接调用内置的`panic`函数也可以引起panic异常. 而由于panic会引起程序的崩溃, 因此一般仅用于严重错误.(有点类似于`Release`支持断言, 不应滥用)     
        - 如果在`defer`函数中调用了内置函数`recover`, 并且定义该`defer`语句的函数发生了panic异常, `recover`会使程序从panic中恢复, 并返回`panic value`. 导致panic异常的函数不会继续运行, 但能正常返回. 未发送panic时调用`recover`, `recover`会返回`nil`.    
            - 通过`defer`和`recover`使程序从panic异常中恢复, 应当有选择的仅在必要时使用.    
    - 通常导致失败的原因不止一种. 因此, 一般函数额外的返回值不再是简单的`bool`类型, 而是`error`类型.    
        - 内置的`error`类型是接口类型.    
        - `error`类型的值可能是`nil`或`non-nil`.`nil`表示成功, `non-nil`表示失败, 并可获取字符串类型的错误信息.    
    - 在`Go`中, 函数运行失败时会返回错误信息, 这些错误信息被认为是一种预期的值而非异常(exception), 这使得`Go`有别于那些将函数运行失败看作是异常的语言.    
    - 常用的五种错误处理策略    
        - 传播错误    
        - 重试失败的操作    
            - 一般用于偶然性的错误, 或由不可预知的问题导致的错误.     
            - 在重试时, 我们需要限制重试的时间间隔或重试的次数, 防止无限制的重试.    
        - 输出错误信息并结束程序 .   
            - 需要注意: 这种策略通常只在`main`中执行   
            - 对于库函数而言, 应仅向上传播错误, 除非该错误意味着程序内部包含不一致性, 即遇到了bug, 才能在库函数中结束程序.    
        - 仅输出错误信息, 继续程序的运行    
        - 忽略错误    
    - 我们应该在每次函数调用后, 都养成考虑错误处理的习惯. 当你觉得忽略某个错误时, 应该清晰的记录下你的意图.    
    - `Go`中错误处理的编码风格(`C-style`)   
        - 检查某个子函数是否失败后, 我们通常将处理失败的逻辑代码放在处理成功的代码之前.     
        - 如果某个错误会导致函数返回, 那么成功时的逻辑代码不应放在`else`语句块中, 而应直接放在函数体中.    
        - `Go`中大部分函数的代码结构几乎相同    
            - 首先是一系列的初始检查, 防止错误发生, 之后是函数的实际逻辑.       
    - `runtime`包允许程序员输出堆栈信息, 以便于调试.    


## Reference Links 
- http://gopl.io
- http://github.com/golang-china/gopl-zh
- http://bitbucket.org/golang-china/gopl-zh