# [Go] golang_test
我的`Golang`学习笔记及示例实验代码.     
`Golang`不愧是现代的`C`语言, 系统的学习下来, 真是足够简洁, 而又灵活且强大, 为服务器应用而生. 迫不及待想用它开发点什么了!!    
- 罗列几项我认为的关键特性如下:     
    - 编译型语言, 仅支持静态链接于是可以run everywhere; 编译过程要么`error`要么`pass`, 于是终于不用再纠结`warning`到底要不要解决的问题.    
    - 没有`class`与继承等概念, 面向对象特性都通过`method`和`interface`来实现(封装通过命名的大小写来实现). `method`绑定到类型仅仅是简单的定义一个属于类型的函数即可, 非常方便扩展.
        - 这一点感觉就是在怼`C++/Java`中的过度抽象, 最终目的是为了更好更方便的解决问题, 一切皆对象未必是最好的solution.     
    - 简洁又强大的`goroutine/channel`! 
    - 完善的`profile`工具, 包括对`cpu`, `memory`, `block`的`profile`, 以及`race detector`工具. 性能分析不再是问题!    
    - 极简的`test`框架    
- 个人不太喜欢的特性:    
    - 固定的`workspace`结构: 一台机器上同时还不能有2个`workspace`了? 
        - 补充说明, 从`GO 1.11`开始已经支持了[Go Modules](https://github.com/golang/go/wiki/Modules)以用于管理依赖. 于是虽然`workspace`的固定结构依然存在, 但至少不再需要强制所有的代码固定在`workspace`目录结构中, 方便了依赖管理的同时, 也增加了易用性.    
        - **Go Modules** 的一些参考资料:    
            - [Go Modules](https://github.com/golang/go/wiki/Modules)
            - [Using Go Modules](https://blog.golang.org/using-go-modules)
            - [How do I find the Go module source cache?](https://stackoverflow.com/questions/52082783/how-do-i-find-the-go-module-source-cache)

### my_gopl_test  
Sample/exercise codes and also my learning notes from learning GOPL《The Go Programming Language》. 
- [我的Golang学习笔记及GOPL示例实验代码](./my_gopl_test/README.md)


### Study Materials 
- https://golang.org
- [How to Write Go Code](https://golang.org/doc/code.html)
- [SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
- [The Go Programming Language](http://gopl.io)
- [Effective Go](https://golang.org/doc/effective_go.html)

