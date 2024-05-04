# goProjectCommu
海量用户通讯系统

该系统分为两个部分，一个是客户端，一个是服务端。可以部署在不同实例上，这时候需要修改一下tcp请求

环境要求：Golang 1.21  //  Redis  //   Mysql(暂时不需要) 

Golang 连接redis环境        go get  github.com/garyburd/redigo/redis

目前 项目框架算是搭完了，支持一下功能：

1.用户注册

2.用户登录

3.用户退出

4.给在线用户群发消息

5.跟在线用户私聊


Todo:

1.支持离线留言

2.能互加好友