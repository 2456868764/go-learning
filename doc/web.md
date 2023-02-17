# web实战项目需求和目标
学习和工作中用过比较多Web框架，对Web框架的原理需要进一步了解比如怎么注册路由，如何路由匹配，如何灵活运用 Web 框架提供的middleware方案解决登录校验、鉴权、日志、tracing、logging,限流等问题
项目实践练习设计一个Web Server，该 Web Server 框架这个教程的很多设计，包括源码，参考了Gin,可以看到很多Gin的影子。

* 上下文设计(Context)
* 前缀路由树Trie树路由(Router)，支持路由通配符匹配、路径参数
* 分组控制(Group)
* 中间件(Middleware)
* HTML模板(Template)
* 错误恢复(Panic Recover)
* 优雅关闭
* 统一错误处理