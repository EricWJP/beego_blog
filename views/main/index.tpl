<div class="container home">
    <div class="left-panel">
        <div class="left-panel-top">
            <div class="title">
                <h4>这周发布</h4>
            </div>
            <div class="container">
                <ul class="microposts">
                    {{range $ind, $el := .microposts }}
                        {{template "microposts/micropost.tpl" $el}}
                    {{end}}
                </ul>
            </div>
        </div>
        <div class="left-panel-bottom">
            <div class="title">
                <h4>最热</h4>
            </div>
            <div class="container">
                <ul class="microposts">
                    {{range $ind, $el := .microposts }}
                        {{template "microposts/micropost.tpl" $el}}
                    {{end}}
                </ul>
            </div>
        </div>
    </div>
    <div class="right-panel">
        <div class="right-panel-top">
            <div class="h5 header">最新留言</div>
            <div>
                <ul>
                    {{range $ind, $el := .users}}
                        <li>
                            <a href="#">{{$el.Name}}</a>
                        </li>
                    {{end}}
                </ul>
            </div>
        </div>
        {{if .signed }}
            <div class="right-panel-bottom">
                <div class="h5 header">关于</div>
                <div>
                    <img src="https://www.kaixin00.com//uploads/image/rand8.jpeg" width="80" height="80"/>
                    <ul>
                        <li>
                            <a href="#">{{.currentUser.Name}}</a>
                        </li>
                    </ul>
                </div>
            </div>
        {{end}}
    </div>
</div>