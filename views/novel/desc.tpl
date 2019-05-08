<div class="partlist-info">
    <dl>
        <dt>《{{.Info.Name}}》</dt>
        <dd class="info">[作者]：<a href="#">{{.Info.Author}}</a><br>
            [点击]：<span>{{.Click}}</span><br>
            [字数]：{{.Info.LengthCont}}字<br>
            [分类]：<a href="/category/{{.Info.Category}}">{{.Info.Category}}</a><br>
        </dd>
        <dd class="zTag01"></dd>
    </dl>
    <div class="newUpdate">
        <a href="/novel/content/{{.New.Id}}" title="{{.New.Title}}"><span>最新章节：{{.New.Title}}</span>({{date_time .New.CreatedAt "Y-m-d H:i:s"}})</a>
    </div>
</div>
<div class="bottonTools">
    <a class="btn-green" href="/novel/content/{{.First.Id}}" title="{{.First.Title}}">开始阅读</a>
</div>
<div class="bookMenu">
    <h2>目录
        <div style="float: right;margin-right: 20px;font-size: 14px;">
            <a href="?sort=desc&page={{.Page.CurrentPage}}">倒序</a> |
            <a href="?sort=asc&page={{.Page.CurrentPage}}">正序</a>
        </div>
    </h2>
    <ul class="list">
        {{range .Page.Data}}
            <li class=""><a href="/novel/content/{{.Id}}" title="{{.Title}}"><span>{{.Title}}</span></a></li>
        {{end}}
    </ul>
    {{template "novel/page.tpl" .}}
    <div class="page">
        <form action="/novel/{{.Info.Id}}" method="get">
            <input type="hidden" name="sort" value="{{.Sort}}"/>
            跳至<input type="text" name="page" size="3"/>页<input type="submit" value="GO">
        </form>
    </div>
</div>
<script>

</script>
