<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- CSRF Token -->

    <title>无广告小说{{if .Name}} - 《{{.Name}}》{{end}}{{if .Title}} - {{.Title}}{{end}}</title>
    <meta content="无广告小说{{if .Name}} - 《{{.Name}}》{{end}}{{if .Title}} - {{.Title}}{{end}};无弹框小说;简单小说;免费小说" name="keywords">
    <meta content="无广告小说{{if .Name}} - 《{{.Name}}》{{end}}{{if .Title}} - {{.Title}}{{end}};提供小说在线阅读，小说TXT全文，网站没有弹窗广告页面简洁。" name="description">
    <meta name="xsrf-token" content="{{.xsrf_token}}" />
    <!-- Scripts -->
    <link rel="stylesheet" href="/static/css/app.css"/>
    <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/layer/2.3/layer.js"></script>
    <!-- Fonts -->
    <link href="/static/web/css/t.css?v=0.1" rel="stylesheet" type="text/css">
</head>
<body class="index">
<header> <!-- 头部 -->

    <div class="login">
        {{if .User}}
        <div style="margin-top: 8px;">
            <a href="#">{{str_limit .User.Username 6 "..."}}</a> |
            <a href="/user/login" onclick="event.preventDefault();document.getElementById('logout-form').submit();">
                退出登录
            </a>
            <form id="logout-form" action="/user/logout" method="POST" style="display: none;">
                {{.xsrfdata}}
            </form>
        </div>
        {{else}}
        <a class="common-btn" href="/user/login">登录</a>
        <a class="common-btn1" href="/user/register">注册</a>
        {{end}}
    </div>

    <nav> <!-- 导航 -->
        <a href="/" title="">首页</a>
    </nav>
</header>
    <div id="app">
        <main class="py-4">
            {{.LayoutContent}}
        </main>
    </div>
</body>
</html>
