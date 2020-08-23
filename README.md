# safari-bookmarks-parser

一个用Go语言开发的解析Safari浏览器的Bookmarks.plist文件（后文简称书签文件）并以HTTP方式对外提供查询功能的工具。

本项目依赖于`howett.net/plist`进行Plist文件解析（可通过`go get howett.net/plist`命令安装）。

## 默认设置

默认端口为8080，默认解析当前目录下的书签文件。

## 可选参数

-bookmarksFilePath：书签文件路径

-serverPort：侦听端口号

## 数据格式

### 请求

http://host:port?kw=xyz

> kw参数的值为搜索关键词（不区分大小写），置空或无此参数表示不执行筛选，将返回所有书签条目

### 响应

[{"title": "foo", "url": "bar"}, ...]

## 异常情况及应对方案

1. 服务器错误，端口（XXX）可能被占用：找到占用端口的进程并结束/指定另一端口启动本应用
1. 发出查询时无结果返回：已知当前用户对书签文件无读取权限会造成这个问题
