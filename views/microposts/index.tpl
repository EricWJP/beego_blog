<div class="container">
    <div class="h2 text-center" style="padding-bottom:1rem;border-bottom: 1px solid #d3d3d3;">博客列表</div>
    <div class="microposts">
        <ul>
            {{range $ind, $el := .microposts }}
                {{template "microposts/micropost.tpl" $el}}
            {{end}}
        </ul>
    </div>
</div>

<nav aria-label="Page navigation example" class="d-flex justify-content-center mt-4">
    <ul class="pagination" style="width:min-content">
        {{if gt .curPage 1}}
            <li class="page-item">
                <a class="page-link" href="{{.href}}search={{.search}}&page={{add .curPage -1}}" aria-label="Previous">
                    <span aria-hidden="true">&laquo;</span>
                    <span class="sr-only">Previous</span>
                </a>
            </li>
        {{end}}
        {{range $i, $e := .pages}}
            <li class="page-item {{if eq $e $.curPage}}active{{end}}"><a class="page-link" href="{{$.href}}search={{$.search}}&page={{$e}}">{{$e}}</a></li>
        {{end}}

        {{if lt .curPage .maxPage}}
            <li class="page-item">
                <a class="page-link" href="{{.href}}search={{.search}}&page={{add .curPage 1}}" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                    <span class="sr-only">Next</span>
                </a>
            </li>
        {{end}}
    </ul>
</nav>