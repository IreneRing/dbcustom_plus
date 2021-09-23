#### dbcustom_plus

[github](https://github.com/by-zxy/dbcustom_plus) | [gitee]()

#### 介绍
主要为go的gorm方法重新封装,涉及增删查改,含有自定义map工具,字符串工具(如: 蛇形,驼峰,字符串空值),
    反射工具等

#### 说明
gorm封装提供方法: 字段,条件,聚合函数,条件函数,预加载,自定义,事务等方法

gorm结构体(接收传入参数)条件封装方法(使用InitParams): 
    1.时间,数组,自定义等类型 建议使用String
    2.参数tag定义:`json:"sppUuidInclude" form:"sppUuidInclude" query:"in,uuid" join:"TTestSpp"`
        1.json/form: 接收传参
        2.query: 查询条件, ("eq_or,test_id","eq,test_id")
                eq_or/eq 查询方式,默认是and,"_or"是or
                test_id  映射字段
        3.join: 左连接表，存在join其query是为连接表查询条件，而非主表
        4.排序: 使用结构体OrderBy
                type PageInfo struct {
                    Page    int64 `json:"page" form:"page"`
                    Limit 	int64 `json:"limit" form:"limit"`
                }
        5.分页: 使用结构体PageInfo
                type OrderBy struct {
                    Desc 	string	`json:"desc" form:"desc"`
                    Asc		string	`json:"asc" form:"asc"`
                }

gorm结构体(返回)封装参考:(references:关联主表字段名, foreignKey:关联字段名)
    TTestBp 	model.TTestBp 	`json:"tTestBp,omitempty" gorm:"foreignKey:Uuid;references:TestId"`
    TTestSpp 	[]model.TTestSpp `json:"tTestSpp,omitempty" gorm:"foreignKey:Uuid;references:TestUuid"`
    
gorm封装使用步骤:  
    1.调用 NewDbmp(defDb *gorm.DB) 设置gorm指针,返回包装结构体指针
    2.可以使用InitParams() 封装查询条件,返回包装结构体指针
    3.使用第2步获取的指针进行函数操作,最好调用curd方法...
    4.使用第3步返回含有Result结构体,其提供影响行数,错误,总数(针对分页)

gorm封装使用事务: 调用 Transaction(func(...)...),实现返回函数 - 调用增删改方法

#### 快速安装
go get -u github.com/by-zxy/dbcustom_plus

#### go mod 安装
go mod tidy
