import redis
import os
import pymysql

#获取最新小说
rds = redis.Redis(host='127.0.0.1', port=6379, password='123456', db=1)
id = rds.blpop("novel_id_insert", 3)
if id:
    id = int(id[1])
    conn = pymysql.connect(
        host='127.0.0.1',
        port=3306,
        db='crawler',
        user='root',
        passwd='123456',
        charset='utf8',
        use_unicode=True
    )
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM novel WHERE crawl_id = %s" % id)
    data = cursor.fetchone()
    if not data:
        os.system("scrapy crawl xiaoshuo -a id=%s" % id)

