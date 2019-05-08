# -*- coding: utf-8 -*-
import scrapy
import time
import pymysql

#更新线上小说
class UpdateXiaoshuoSpider(scrapy.Spider):
    # handle_httpstatus_list = [404]
    name = 'update_xiaoshuo'
    allowed_domains = ['www.23us.so']
    header = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0'}

    def start_requests(self):
        # url = 'http://www.23us.so/'
        # #初始化小说地址
        # self.id = getattr(self, 'id', None)
        # if id is not None:
        #     url = url + 'xiaoshuo/' + self.id + '.html'
        # else:
        #     return
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
        novel_list = self.novel_list()
        for id in novel_list:
            url = 'http://www.23us.so/'
            url = url + 'files/article/html/32/' + str(id[1]) + '/index.html'
            yield scrapy.Request(url, self.novel_parse, meta={'novel_id': id[0]})

    def novel_parse(self, response):
        novel_id = response.meta['novel_id']
        data = self.detail_parse(novel_id)
        #判断小说本地是否有章节
        if data:
            cont_list = response.css('table#at td a')
            list = dict()
            #获取网站上小说总章节名称
            for cont in cont_list:
                title = cont.css('a::text').extract_first()
                url = cont.css('a::attr(href)').extract_first()
                list.update({title: url})
            novel_list = self.novel_name_list(novel_id)
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
                yield scrapy.Request(next_page, callback=self.cont_parse, meta={'start_url': next_page, 'novel_id': novel_id})
                time.sleep(1)

    def cont_parse(self, response):
        title = response.css('div#amain dl dd h1::text').extract_first()
        cont = response.css('dd#contents').extract_first()
        sort = response.meta['start_url'].split('/')[-1].replace('.html', '')
        novel_id = response.meta['novel_id']

        id = self.insert_data(novel_id=novel_id, title=title, sort=sort)
        if id != 0:
            self.insert_content(art_id=id, content=cont)
            # fileName = '/data/novel/%s/%s.txt' % (self.novel_id, id)
            # f = codecs.open(fileName, "w+", 'utf-8')
            # f.write(cont)
            # f.close()

    #获取所有小说Url ID
    def novel_list(self):
        self.cursor.execute("SELECT id, crawl_id FROM novel WHERE status = 1")
        list_data = self.cursor.fetchall()
        return list_data

    #获取单个章节
    def detail_parse(self, novel_id):
        self.cursor.execute("SELECT * FROM articles WHERE novel_id = %s Limit 1" % (novel_id))
        data = self.cursor.fetchone()
        return data

    #获取小说所有章节
    def novel_name_list(self, novel_id):
        self.cursor.execute("SELECT * FROM articles WHERE novel_id = %s " % (novel_id))
        list_data = self.cursor.fetchall()
        name_list = []
        for val in list_data:
            name_list.append(val[2])
        return name_list

    def insert_data(self, novel_id, title, sort):
        sql = "INSERT INTO articles(novel_id, title, sort, created_at) VALUE (%s, '%s', %s, %s)" % \
              (novel_id, title.strip(), sort, int(time.time()))
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