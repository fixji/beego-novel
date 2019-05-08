import redis
import os
import pymysql

#每7天跑一次 获取推荐小说
rds = redis.Redis(host='127.0.0.1', port=6379, password='123456', db=1)
ids = []
conn = pymysql.connect(
    host='127.0.0.1',
    port=3306,
    db='crawler',
    user='root',
    passwd='123456',
    charset='utf8',
    use_unicode=True
)
for id in rds.lrange("novel_id_recommend", 0, -1):
    id = int(id)
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM novel WHERE crawl_id = %s" % id)
    data = cursor.fetchone()
    if not data:
        os.system("scrapy crawl xiaoshuo -a id=%s" % id)





