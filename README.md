dbcustom_plus

#### 介绍
主要为go的gorm方法重新封装，涉及增删查改，含有自定义map工具，字符串工具(如：蛇形，驼峰，字符串空值)，
    反射工具等

#### 说明
gorm封装提供方法： 字段，条件，聚合函数，条件函数，预加载，自定义，事务等方法
gorm封装使用步骤： 1.调用 NewDB(...) 设置gorm指针
                2.调用 NewDbmp() 返回包装结构体指针
                3.可以使用InitParams() 封装查询条件，返回包装结构体指针
                4.使用第2步获取的指针进行函数操作，最好调用curd方法...
                5.使用第3步返回含有Result结构体，其提供影响行数，错误，总数(针对分页)
gorm封装使用事务：调用 Transaction(func(...)...)，实现返回函数 - 调用增删改方法

#### 快速安装
go get -u github.com/by-zxy/dbcustom_plus

#### go mod 安装
go mod tidy
