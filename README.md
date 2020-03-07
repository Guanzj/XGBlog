<!--
 * @Author: moonmist.guan
 * @Date: 2020-03-06 20:36:00
 * @LastEditTime: 2020-03-08 01:18:10
 * @FilePath: /XGBlog/README.md
 * @Description: 
 -->

# XGBlog

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

使用基于go语言的Gin框架开发的Blog系统后台API接口

首先感谢我的女朋友@xi

本项目为React+Gin+Gorm+MySQL构建的前后端分离博客项目

其中后端采用基于go语言的Gin框架开发API接口

使用MySQL数据库存储相关信息

API接口设计原则遵循restfull设计原则，API文档采用postman文档，链接如下

https://documenter.getpostman.com/view/4578839/SzRw2Wpw?version=latest

仓库包含以下内容

1.博客项目后端开发代码

2.gormt mysql 转换工具

3.项目go mod配置文件

4.项目启动及数据库等配置文件


## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [License](#license)

## Background

作为一名程序员，有记笔记的习惯但是很惭愧一直没有自己的博客，正好借此机会开发一套适合自己需求的blog框架



## Install

得益于go优秀的部署便捷性，只需要clone这份代码后执行以下代码即可

```sh
$ go run main.go
```

## Usage

项目初始接口为localhost:8080，可参考API文档进行使用

## 更新记录

>2020.03.07
>
>- 使用SFTP实时同步本地与服务器代码，下一步考虑使用k8s进行部署，单纯使用build会出现配置文件无法读取的问题，所以暂不考虑

## 待完成事项

- [x] 完成本地和服务器端文件同步
- [ ] 使用k8s进行服务器部署

## Maintainers

[@Guanzj](https://github.com/Guanzj).


## License

[MIT](LICENSE) © Monnmist
