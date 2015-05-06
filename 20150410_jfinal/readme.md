# jfinal工程搭建

* 创建dynamic web工程，target runtime设置为none，default output folder设置为WebRoot\WEB-INF\classes，content directory设置为WebRoot，生成web.xml
* 拷贝jetty-server-xxx.jar和jfinal-xxx.jar到WebRoot/WEB-INF/lib下
* web.xml中增加jfinal的config作为filter
* 实现DemoConfig继承自JFinalConfig
* 实现HelloController，在DemoConfig中增加路由。
* 配置debug configration，增加意向主类执行com.jfinal.core.JFinal

# jfinal中c3p0和activerecord的使用方法
