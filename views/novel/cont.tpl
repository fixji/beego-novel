<div class="novel">
    <div class="bottonTools1">
        {{if .PrevCont}}
        <a href="/novel/content/{{.PrevCont.Id}}" title="{{.PrevCont.Title}}">上一页</a>
        {{end}}
        <a href="/novel/{{.Novel.Id}}">目录</a>
        {{if .NextCont}}
        <a href="/novel/content/{{.NextCont.Id}}" title="{{.NextCont.Title}}">下一页</a>
        {{end}}
    </div>
    <div class="ContentBody">
        <h2>{{.Article.Title}}</h2>
        <div id="article">
            {{str2html .Cont}}
        </div>
    </div>
    <div class="bottonTools1" style="margin-bottom: 30px;">
        {{if .PrevCont}}
        <a href="/novel/content/{{.PrevCont.Id}}" title="{{.PrevCont.Title}}">上一页</a>
        {{end}}
        <a href="/novel/{{.Novel.Id}}">目录</a>
        {{if .NextCont}}
        <a href="/novel/content/{{.NextCont.Id}}" title="{{.NextCont.Title}}">下一页</a>
        {{end}}
    </div>
</div>
