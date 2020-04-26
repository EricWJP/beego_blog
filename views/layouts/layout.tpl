<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>{{.siteName}}</title>
    <link rel="stylesheet" href="/static/css/bootstrap-4.4.1.min.css" />
    <link rel="stylesheet" href="/static/css/app.css" type="text/css" />
{{/*    <script src="https://cdn.jsdelivr.net/npm/jquery@3.4.1/dist/jquery.slim.min.js"*/}}
{{/*            integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n"*/}}
{{/*            crossorigin="anonymous"></script>*/}}
    <script src="/static/js/jquery-3.4.1.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js"
            integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo"
            crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js"
            integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6"
            crossorigin="anonymous"></script>
</head>
<body style="background-color: #f5f5d5;color: #111;font-size: 12px;">
    <div class="main">
        <div class="container" style="border-bottom: 1px solid #d3d3d3;max-width: 90%;font-size: 1.25rem;">
            <nav class="navbar navbar-expand-lg navbar-light" >
                <a class="navbar-brand" href="/">Eric的博客</a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
                        aira-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item">
{{/*                            {{if eq .data.class_id  0}}active{{end}}*/}}
                            <a class="nav-link" href="/">首页 <span class="sr-only">(current)</span></a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/microposts">博客列表</a>
                        </li>
                    </ul>
                    <form class="form-inline my-2 my-lg-0 mr-auto" action="/search" method="post">
                        <input class="form-control mr-sm-2" type="search" name="search" value="{{.search}}"  placeholder="查找博客..." aria-label="Search"/>
                        <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
                    </form>
                    <ul class="navbar-nav">
                        {{if .signed }}
                            <li class="nav-item">
                                <a class="nav-link" href="/microposts/create">写博客</a>
                            </li>
                            <li class="nav-item dropdown">
                                <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
                                   aria-haspopup="true" aria-expanded="false">
                                    我的
                                </a>
                                <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                                    <a class="dropdown-item" href="/my_microposts">我的博客</a>
                                    <div class="dropdown-divider"></div>
                                    <a class="dropdown-item" href="#">我收藏的</a>
                                    <a class="dropdown-item" href="#">我关注的</a>
                                    <div class="dropdown-divider"></div>
                                    <a class="dropdown-item" href="#">关注我的</a>
                                    <div class="dropdown-divider"></div>
                                    <a class="dropdown-item" href="/setup">设置</a>
                                    <a class="dropdown-item" href="/reset_password">修改密码</a>
                                    <a class="dropdown-item" href="/sign_out" >登出</a>
                                </div>
                            </li>
                        {{ else }}
                            <li class="nav-item">
                                <a class="nav-link" href="/sign_in">登录</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="/sign_up">注册</a>
                            </li>
                        {{ end }}
                    </ul>
                </div>
            </nav>
        </div>
        <div class="container" style="margin-bottom: 40px;">
            {{range $a,$b := .flash}}
                <div class="alert alert-{{$a}}">{{$b}}</div>
            {{end}}
        </div>

        {{.LayoutContent}}
    </div>
</body>
</html>