<div class="container micropost">
    <div class="h2 text-center">{{.json.Title}}</div>
    <div class="desc">
        <ul>
            <li>
                作者：{{.userName}}
            </li>
            <li>
                日期：{{.json.UpdatedAt | timeFormat}}
            </li>
        </ul>
    </div>
    <div class="content">
        {{.json.Content}}
    </div>
</div>