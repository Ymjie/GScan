## 简介
使用Golang编写高性能网站爬虫、内容分析工具

**本项目只是本人个人学习开发并维护，本人不保证任何可用性，也不对使用本软件造成的任何后果负责。**
## 功能
1. 扫描网站上的外链，及时发现废弃域名被抢注指向非法网站，避免因为黑链而被网安部门通报。
2. 扫描网站的内容，及时发现敏感信息(如身份证)，避免因为信息泄露而被网安部门通报。
3. 扫描网站开放下载的文件连接，对内容进行排查及时发现敏感信息，避免因为信息泄露而被网安部门通报。

**Email:i@vshex.com**

因为本人水平有限，大家在使用过程中发现什么问题请和我联系，谢谢！

使用说明
--------------------------------------------------------------------------------
**配置文件说明:** 见`config.yml`注释

**使用方法：**
1. 编辑配置文件、调整`whitelist.txt`
2. 将目标网站保存到`url.txt`
3. 运行infoscan.exe (推荐在命令行中运行)
```
InfoScan

Usage:
   infoscan.exe
   infoscan.exe ls               #列出所有任务
   infoscan.exe export <JobID>   #导出任务结果
```

_如有需要可自行编译Linux版本_

## 文件说明

**目标文件：**同级目录下的`url.txt`
每行一个，类似`https://brey.cn`

**白名单：同级目录下的`whitelist.txt`**
每行一个，扫描到的外链中如果包含白名单中所列内容，将不处理，可以过滤安全的外链，例如：`edu.cn`。根据自己学校的情况，合理的设备白名单，可以大大减少人工核查的工作量。

## 处理结果说明

`data.db`：SQLit数据库文件
处理结果为`EXCEL`文件，目前有**外部链接**、**页面敏感信息**、**关键词检测**、**可下载文件**表

**对于外部链接的说明：**
可以针对找到的外链进行验证，看能否正常打开(针对无法打开的外链需要重点关注)，同时对可以正常打开的网站会提取标题，便于发现有问题的外链。
对于状态码200，标题为空的结果需要特别关注（手动访问）

**对于页面敏感信息的说明：**
目前只进行18位身份证的匹配，有一定误报率。

**对于关键词检测的说明：**
有一定误报率，需要特别关注匹配关键词数量多的结果。
欢迎贡献关键词。

**对于可下载文件的说明：**
推荐使用迅雷等专业的下载工具批量下载，然后使用`FileSearchy`进行文件扫描。


## Todo
- [ ] 限速功能
- [ ] 菜单功能 
   - [ ] ~~后台运行~~
   - [ ] ~~实时设置Crawler~~
   - [x] 查看任务列表
   - [x] 导出历史数据
   - [ ] 保存任务进度
- [ ] 完善内容分析模块
- [ ] 添加调用Chrome内核的Spider