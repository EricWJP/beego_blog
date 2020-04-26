<div class="container">
    <div class="h2 text-center" style="margin: 0 40px;padding-bottom:10px;border-bottom: 1px solid #d3d3d3;">博客列表</div>
    <ul class="microposts">
        {{range $ind, $el := .microposts }}
            {{template "microposts/micropost.tpl" $el}}
        {{end}}
    </ul>
</div>