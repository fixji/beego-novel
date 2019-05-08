<!-- 搜索与阅读记录 -->
<div class="searchbox">
    <form action="/novel/search" method="get">
        <input class="search-input" name="name" autocomplete="off" type="text" value=""/>
        <input class="bton-search" type="submit" value=""/>
    </form>
</div>
<div class="hot">
    {{range .Hot}}
    <div class="item">
        <div class="p10">
            <dl>
                <dt>
                    <span>{{.Author}}</span>
                    <a href="/novel/{{.Id}}">{{.Name}}</a>
                </dt>
                <dd><a href="/novel/{{.Id}}" style="color: #AAA;-webkit-tap-highlight-color:transparent;">{{.ContentValidity}}</a></dd>
            </dl>
        </div>
    </div>
    {{end}}
</div>
<div class="module">
    <h2 class="tit-news"><span>武侠仙侠</span></h2>
    <ul class="list">
        {{range .XianXia}}
            <li class="" style="border-bottom: 1px solid #ECECEC;">
                <a href="/novel/{{.Id}}" title="{{.Name}}">
                    <span style="color: #999999;">[{{.Category}}]</span>
                    {{.Name}}
                </a>
            </li>
        {{end}}
    </ul>
</div>
<div class="module">
    <h2 class="tit-news"><span>玄幻奇幻</span></h2>
    <ul class="list">
        {{range .XuanHuan}}
            <li class="" style="border-bottom: 1px solid #ECECEC;">
                <a href="/novel/{{.Id}}" title="{{.Name}}">
                    <span style="color: #999999;">[{{.Category}}]</span>
                    {{.Name}}
                </a>
            </li>
        {{end}}
    </ul>
</div>
<div class="module">
    <h2 class="tit-news"><span>都市言情</span></h2>
    <ul class="list">
        {{range .DuShi}}
            <li class="" style="border-bottom: 1px solid #ECECEC;">
                <a href="/novel/{{.Id}}" title="{{.Name}}">
                    <span style="color: #999999;">[{{.Category}}]</span>
                    {{.Name}}
                </a>
            </li>
        {{end}}
    </ul>
</div>
