# -*- coding: utf-8 -*-
import scrapy
import time
import pymysql
import codecs
import os

#指定id获取小说
class XiaoshuoSpider(scrapy.Spider):
    # handle_httpstatus_list = [404]
    name = 'xiaoshuo'
    allowed_domains = ['www.23us.so']
    header = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0'}

    def start_requests(self):
        url = 'http://www.23us.so/'
        #初始化小说地址
        self.id = getattr(self, 'id', None)
        if id is not None:
            url = url + 'xiaoshuo/' + self.id + '.html'
        else:
            return
        #初始化mysql
        self.connect = pymysql.connect(
            host='127.0.0.1',#数据库地址
            port=3306,# 数据库端口
            db='crawler', # 数据库名
            user='root', # 数据库用户名
            passwd='123456', # 数据库密码
            charset='utf8', # 编码方式
            use_unicode=True
        )
        self.cursor = self.connect.cursor()

        yield scrapy.Request(url, self.parse)

    def parse(self, response):
        #收集小说介绍页数据
        category = response.css("div.bdsub dl dd div.fl table td a::text").extract_first()
        self.novel_name = response.css("div.bdsub dl dd h1::text").extract_first().replace(' 全文阅读', '')
        content_validity = response.css("div.bdsub dl dd p").extract()[3]

        info = response.css("div.bdsub dl dd div.fl table td::text").extract()
        author = "".join(info[1].split())

        status = "".join(info[2].split())
        if status == '连载中':
            status = 1
        else:
            status = 2

        length_cont = "".join(info[4].split())
        length_cont = length_cont.replace('字', '')

        update_time = "".join(info[5].split())
        updated_at = int(time.mktime(time.strptime(update_time, '%Y-%m-%d%H:%M:%S')))
        #判断小说是否需要继续爬取
        self.novel_id = self.insert_novel(self.novel_name, author, category, content_validity, status, length_cont, updated_at)
        if self.novel_id == 0:
            return
        else:
            url = 'http://www.23us.so/'
            url = url + 'files/article/html/32/' + self.id + '/index.html'
            yield scrapy.Request(url, self.novel_parse)

    def novel_parse(self, response):
        data = self.detail_parse()
        #创建目录(已废弃)
        # file = '/data/novel/%s' % self.novel_id
        # if not os.path.exists(file):
        #     os.mkdir(file)
        time.sleep(1)
        #判断小说本地是否有章节
        if data:
            cont_list = response.css('table#at td a')
            list = dict()
            #获取网站上小说总章节名称
            for cont in cont_list:
                title = cont.css('a::text').extract_first()
                url = cont.css('a::attr(href)').extract_first()
                list.update({title: url})
            novel_list = self.novel_name_list()
            #对比小说已存在章节
            urls = []
            for key in list:
                if key not in novel_list:
                    urls.append(list[key])
        else:
            urls = response.css('table#at a::attr(href)').extract()

        #爬取所有没有的章节
        for url in urls:
            if url is not None:
                next_page = response.urljoin(url)
                yield scrapy.Request(next_page, callback=self.cont_parse, meta={'start_url': next_page})
            time.sleep(1)

    def cont_parse(self, response):
        title = response.css('div#amain dl dd h1::text').extract_first()
        cont = response.css('dd#contents').extract_first()
        sort = response.meta['start_url'].split('/')[-1].replace('.html', '')

        id = self.insert_data(title=title, sort=sort)
        if id != 0:
            self.insert_content(art_id=id, content=cont)
            # fileName = '/data/novel/%s/%s.txt' % (self.novel_id, id)
            # f = codecs.open(fileName, "w+", 'utf-8')
            # f.write(cont)
            # f.close()

    def insert_novel(self, name, author, category, content_validity, status, length_cont, updated_at):
        #查询小说是否存在
        self.cursor.execute("SELECT * FROM novel WHERE name = '" + name + "' and author = '" + author + "'")
        data = self.cursor.fetchone()
        #判断小说是否已存在
        if data:
            #判断小说数字是否增加
            if int(length_cont) == int(data[6]):
                return 0
            #返回小说ID
            return data[0]
        #新增小说
        sql = "INSERT INTO novel(name, author, category, content_validity, status, length_cont, crawl_id, created_at, updated_at) \
              VALUE ('%s', '%s', '%s', '%s', %s, %s, %s, %s, %s)" % \
              (name, author, category, content_validity, status, length_cont, self.id, int(time.time()), updated_at)
        try:
            self.cursor.execute(sql)
            id = self.connect.insert_id()
            self.connect.commit()
        except:
            self.connect.rollback()
            return 0
        return id

    #获取单个章节
    def detail_parse(self):
        self.cursor.execute("SELECT * FROM articles WHERE novel_id = %s Limit 1" % (self.novel_id))
        data = self.cursor.fetchone()
        return data

    #获取小说所有章节
    def novel_name_list(self):
        self.cursor.execute("SELECT * FROM articles WHERE novel_id = %s " % (self.novel_id))
        list_data = self.cursor.fetchall()
        name_list = []
        for val in list_data:
            name_list.append(val[2])
        return name_list

    def insert_data(self, title, sort):
        sql = "INSERT INTO articles(novel_id, title, sort, created_at) VALUE (%s, '%s', %s, %s)" % \
              (self.novel_id, title.strip(), sort, int(time.time()))
        try:
            self.cursor.execute(sql)
            id = self.connect.insert_id()
            self.connect.commit()
        except:
            self.connect.rollback()
            return 0
        return id

    def insert_content(self, art_id, content):
        sql = "INSERT INTO content(art_id, content) VALUE (%s, '%s')" % \
              (art_id, content)
        try:
            self.cursor.execute(sql)
            id = self.connect.insert_id()
            self.connect.commit()
        except:
            self.connect.rollback()
            return 0
        return id