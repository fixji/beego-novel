<!DOCTYPE HTML>
<html>
<head>
    <title>无广告小说</title>
    <meta content="无广告小说" name="keywords">
    <meta content="无广告小说" name="description">
    <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" name="viewport"/>
    <meta charset="utf-8">
    <link href="/static/web/css/t.css" rel="stylesheet" type="text/css">
</head>
<body class="index">
<!-- 搜索与阅读记录 -->
<div class="searchbox">
    <form action="/novel/search" method="get">
        <input class="search-input" name="name" type="text" value=""/>
        <input class="bton-search" type="submit" value=""/>
    </form>
</div>
<h1 style="text-align: center;margin-top: 20px;"> 404,{{if .ErrorMessage}}{{.ErrorMessage}}{{else}}未找到该页面!{{end}}</h1>
<a href="/"><h3 style="margin-top: 20px;text-align: center;">返回首页</h3></a>
</body>
</html>