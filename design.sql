CREATE TABLE `XGBlog`.`mb_articles` (
  `articleID` int(10) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `cateID` int(10) NOT NULL DEFAULT '0' COMMENT '所属分类ID',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '状态：1-公开、2-私密' ,
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '简介、描叙',
  `keywords` varchar(256) NOT NULL DEFAULT '' COMMENT '关键词，用英文逗号隔开',
  `imgPath` varchar(256) NOT NULL DEFAULT '' COMMENT '图片',
  `opID` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `opUser` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifyTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`articleID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='文章表';

CREATE TABLE `XGBlog`.`mb_articles_contents` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `articleID` int(10) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `showType` tinyint(2) NOT NULL DEFAULT '0' COMMENT '内容展示类型：1-html、2-markdown',
  `contents` text NOT NULL COMMENT '文章内容',
  PRIMARY KEY (`id`),
  KEY `idx_articleID` (`articleID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='文章内容表';

CREATE TABLE `XGBlog`.`mb_articles_cate` (
  `cateID` int(10) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar (64) NOT NULL DEFAULT '' COMMENT '分类名',
  `cName` varchar(64) NOT NULL DEFAULT '' COMMENT '分类中文名',
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '简介、描叙',
  `parentID` int(10) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `orderby` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `opID` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `opUser` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifyTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`cateID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章类型表';

CREATE TABLE `XGBlog`.`mb_admins` (
  `adminID` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：1-正常，2-禁用',
  `opID` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `opUser` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifyTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`adminID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';

insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (1,'php','PHP');
insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (2,'mysql','Mysql');
insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (3,'go','Go');
insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (4,'linux','Linux');
insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (5,'qianduan','前端');
insert into `XGBlog`.`mb_articles_cate`(cateID,`name`, cName) value (6,'other','其他');


insert into `XGBlog`.`mb_articles`(cateID, title, description) value (1, '测试文章01','测试文章01');
insert into `XGBlog`.`mb_articles`(cateID, title, description) value (1, '测试文章02','测试文章02');
insert into `XGBlog`.`mb_articles`(cateID, title, description) value (1, '测试文章03','测试文章03');


insert into `XGBlog`.`mb_articles_contents`(articleID, contents) value (1,'测试文章01');
insert into `XGBlog`.`mb_articles_contents`(articleID, contents) value (2,'测试文章02');
insert into `XGBlog`.`mb_articles_contents`(articleID, contents) value (3,'测试文章03');

insert into `XGBlog`.`mb_admins` (username,password,status) value ('admin', 'admin',1);

