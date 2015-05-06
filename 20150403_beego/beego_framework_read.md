go get github.com/astaxie/beego
go get github.com/beego/bee

bee new mybeegoproject
bee run
http://localhost:8080/

bee pack

bee generate 如何生成mvc？？

配置：模式、端口、db参数等

model：没有明确的mysql实例；mongo应该很容易
    TODO http://beego.me/docs/mvc/model/overview.md

controller：固定路由，正则路由（通过this.Ctx.Input.Param(":id")获取），自定义方法和restful规则，反射匹配，注解路由，命名空间
    TODO session控制：http://beego.me/docs/mvc/controller/session.md
    返回json、xml等 http://beego.me/docs/mvc/controller/jsonxml.md

view：自动模板
