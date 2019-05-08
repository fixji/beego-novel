# 爬虫Python3  爬取www.23us.so

安装
```
需要安装scrapy 参考
http://www.scrapyd.cn/doc/123.html
额外安装redis和mysql扩展

```

scrapy下面文件放到scrapy里面运行,每个文件对应不同命令
爬取指定小说命令
	scrapy crawl xiaoshuo -a id=14019 

写的有点简略,因为我也忘了要下载哪些扩展