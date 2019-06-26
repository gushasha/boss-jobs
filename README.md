# 爬取BOSS直聘的招聘信息

看完golang的文档，写个爬虫练手，记录下学习过程。

## 目标
目标站点是[Boss直聘](https://www.zhipin.com/)，以成都为例，爬取招聘信息。
字段9个，如下所示：
![招聘信息字段示例](https://github.com/gushasha/boss-jobs/raw/master/docs/img/job.jpeg)

## 分析目标站点

**url地址分析**

进入boss直聘首页，城市切换为“成都”（可自行选择），选择热门职位“Java”（可自行填写或选择）。则跳转到新页面，地址为: `https://www.zhipin.com/c101270100/?query=Java&industry=&position=`，其中：

* `c101270100` 为城市代码（可审查元素查看其它城市代码）
* `query=java`参数为语言类型（可审查元素查看其他语言类型）

**分页分析**

页面最底部的分页栏，点击页码"2"可以看到页面跳转到`https://www.zhipin.com/c101270100/?query=Java&page=2&ka=page-2` ，多了一个参数`page=2`。页码部分审查元素，如下图：
![页码审查元素](https://github.com/gushasha/boss-jobs/raw/master/docs/img/page.png)
页码规律为，参数page递增1。如何得知最后一页呢？

点击下页...下页...， 直到点不动（也可手动修改url参数的页码快速定位），审查元素，“下一页的图标” class 为disabled。因此，可以通过page内的最后一个a标签的class判断是否为最后一页。

```
<div class="page">
        <a href="/c101270100/?query=Java&amp;page=9" ka="page-prev" class="prev"></a>
        <a href="/c101270100/?query=Java&amp;page=1" ka="page-1">1</a>
        <span>...</span>
        <a href="/c101270100/?query=Java&amp;page=8" ka="page-8">8</a>
        <a href="/c101270100/?query=Java&amp;page=9" ka="page-9">9</a>
        <a href="javascript:;" class="cur" ka="page-cur">10</a>
        <!-- 注意：最后一页的页面，page-next 的class为 disabled -->
        <a href="javascript:;" ka="page-next" class="next disabled"></a>
</div>
```

## 安装

```
// 拉取项目代码
go get -u -v github.com/gushasha/boss-jobs

// govendor安装包和依赖
cd $GOPATH/src/github.com/gushasha/boss-jobs
govendor sync

// 修改配置文件
修改 `boss_jobs/conf/init.go`中的配置

// 创建数据库`spiders`，数据表`spiders_boss_jobs`
// 表结构在 `boss-jobs/docs`中

// 启动
go run main.go

```

#### 项目目录

```
$GOPATH/src/github.com/gushasha/boss-jobs/
|____docs
| |____spiders_boss_jobs.sql
|____main.go
|____models
| |____jobs.go
| |____model.go
|____parse
| |____jobs.go
|____vendor
| |____vendor.json
```

#### 项目用到的包
* `github.com/PuerkitoBio/goquery` 解析Dom
* `github.com/jinzhu/gorm` 数据库操作 


#### 初始化项目数据库
新建`spiders`数据库，编码为`utf8_general_ci `，在`spiders`数据库下新建表`spiders_boss_jobs`：

```
CREATE TABLE `spiders_boss_jobs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `jid` varchar(255) NOT NULL DEFAULT '' COMMENT '第三方平台上的ID：jobID',
  `job_type` varchar(30) NOT NULL COMMENT '工作类型',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `salary_range` varchar(255) NOT NULL DEFAULT '' COMMENT '薪资范围',
  `work_years` varchar(255) DEFAULT '' COMMENT '工作年限',
  `education` varchar(255) DEFAULT '' COMMENT '学历要求',
  `company_name` varchar(255) DEFAULT '' COMMENT '公司名称',
  `company_address` varchar(255) NOT NULL DEFAULT '' COMMENT '公司地址',
  `company_label` varchar(255) DEFAULT '' COMMENT '公司类型',
  `financing_stage` varchar(255) DEFAULT '' COMMENT '融资阶段',
  `company_employees_count` varchar(255) DEFAULT '' COMMENT '公司规模-员工人数',
  `create_time` int(10) unsigned DEFAULT '0',
  `update_time` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_jid` (`jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT "boss招聘信息表";
```

