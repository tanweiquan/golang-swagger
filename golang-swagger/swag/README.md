## Swagger

### 快速入门
#### 安装
```shell
go get -u github.com/swaggo/swag/cmd/swag
go get "github.com/swaggo/gin-swagger" // gin-swagger middleware
go get "github.com/swaggo/files" // swagger embed files
```

#### 初始化
```shell
swag init -g swag.go -o ../../api/swagger
```
如果不使用`-g`指定入口，则默认使用寻找`main.go`文件，可通过`swag init -h`查看所有命令选项

#### 示例
```go
package main

import (
    // 导入init生成的文件	
    _ "l-sample/docs/swag"
    
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    route := gin.Default()
    route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    _ = route.Run()
}
```
通过浏览器访问地址 `http://host:port/swagger/index.html` 即可。 

### 常用字段

#### 应用描述
|字段|说明|示例|
|---|---|---|
|title|必填 应用程序的名称|@title Swagger Example API|
|version|必填 提供应用程序API的版本|@version 1.0|
|description|应用程序的简短描述|@description This is a sample|
|host|运行API的主机（主机名或IP地址）|@host localhost:8080|
|BasePath|运行API的基本路径|@BasePath /api/v1|

#### 接口描述
|字段|说明|示例|
|---|---|---|
|summary|简短摘要||
|description|行为的详细说明||
|tags|标签列表，以逗号分隔||
|accept|使用的MIME类型||
|produce|生成的MIME类型||
|router|路径定义，以空格分隔： [path] [httpMethod]||
|param|请求参数，用空格分隔：[param_name] [param_type] [data_type] [is_mandatory?] [comment] [attribute?]||
|success|成功响应，以空格分隔：[return_code] [param_type?] [data_type] [comment]||
|header|头字段，以空格分隔，格式同success||
|failure|故障响应，以空格分隔，格式同success||

常用MIME类型为`json`，更多详情可访问[swagger中文教程](https://github.com/swaggo/swag/blob/master/README_zh-CN.md) 

常用的请求参数类型(param_type)如下：
- query
- path
- header
- body
- formData

常用的数据类型（data_type）如下：
- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- 自定义类型

常用的属性（attribute）如下：
- enums：枚举值
- default：设置默认值
- maximum：设置最大数值
- minimum：设置最小数值
- maxLength：设置最大长度
- minLength：设置最小长度

注意：属性值不仅可以在swagger参数中使用，还可以在自定义的结构体中使用
