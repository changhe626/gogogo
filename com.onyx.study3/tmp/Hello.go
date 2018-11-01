package tmp

//建议package 的名称和所所属于的文件名一致....
//一个文件夹内所有的文件的package必须是一样的...

/**

goroot  和  gopath的设置和注意点


获取扩展,下载到classpath中去
go  get  网址


go  build  编译
go  run  编译运行
go  get   动态获取远程代码包



package是最基本的分发单位和工程管理中依赖关系的提现

要生成go 的可执行程序,必须是main 的package包,且该包下有main()函数.

同一个路径下只能存在一个package,一个packgae可以拆成多个源文件组成


import  导入源代码文件所依赖的package包

不导入没用到,否则报错
2种形式.

如果导入的依赖其他包(B),会首先导入B包,然后初始化B包中常量和变量,最后如果B包中有init没回自动执行init()


导入完成后才对main中常量和变量进行初始化,再执行mian.


包只导入一次...不管导入几次



*/
