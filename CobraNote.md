>cobra是一个基于go的命令行工具框架

### 程序框架
通常基于 Cobra 的应用程序将遵循以下组织结构
```
  ▾ appName/
    ▾ cmd/
        add.go
        your.go
        commands.go
        here.go
      main.go
```

### rootCmd
函数启动顺序
```go
// The *Run functions are executed in the following order://   * PersistentPreRun()  
//   * PreRun()  
//   * Run()  
//   * PostRun()  
//   * PersistentPostRun()
```

rootCmd是不加参数时的默认调用命令
```go
var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}
```

### 创建命令
```go
//利用AddCommand成员函数追加
func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
```

创建的命令会继承rootCmd的persistent方法


### flag
由于标志是在不同的位置定义和使用的，因此我们需要在具有正确范围的外部定义一个变量来分配标志以进行操作。

```go
var Verbose bool
```
标志可以是**持久**’的，这意味着此标志将对它被分配到的命令以及该命令下的所有命令可用。
