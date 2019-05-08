<div role="navigation" class="page">
    {{range $elem := .Page.PageUrls}}
        {{str2html $elem}}
    {{end}}
</div>

