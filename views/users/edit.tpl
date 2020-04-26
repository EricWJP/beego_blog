<div class="container">
    <div class="h2 text-center">设 置</div>
    <form action="/setup" method="post">
        <div class="form-group">
            <label for="userName">用户名</label>
            <input class="form-control" type="text" value="{{.json.Name}}" id="userName" name="Name" aria-describedby="nameHelp" disabled/>
            <small id="nameHelp" class="form-text text-muted">注册后不可更改！</small>
        </div>
        <div class="form-group">
            <label for="userEmail">邮件</label>
            <input class="form-control" type="text" value="{{.json.Email}}" id="userEmail" name="Email" aria-describedby="emailHelp" />
            <small id="emailHelp" class="form-text text-muted">我们一定会保密！</small>
        </div>
        <div class="form-group">
            <label for="userPhone">手机</label>
            <input class="form-control" type="text" value="{{.json.Phone}}" id="userPhone" name="Phone" aria-describedby="phoneHelp" />
            <small id="phoneHelp" class="form-text text-muted">我们一定保密！</small>
        </div>
        <div class="form-group">
            <label>性别： </label>
            <div class="form-check form-check-inline">
                {{if eq (or .json.Gender "n") "man" }}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender1" value="man" checked />
                {{else}}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender1" value="man" />
                {{end}}
                <label class="form-check-label" for="userGender1">男</label>
            </div>
            <div class="form-check form-check-inline">
                {{if eq (or .json.Gender "n") "female" }}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender2" value="female" checked />
                {{else}}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender2" value="female" />
                {{end}}
                <label class="form-check-label" for="userGender2">女</label>
            </div>
        </div>
        <div class="form-group">
            <label for="userComment">个人描述</label>
            <textarea class="form-control" name="Comment" id="userComment" rows=5>{{.json.Comment}}</textarea>
        </div>

        <button type="submit" class="btn btn-primary form-control font-weight-bold">提 交</button>
    </form>
</div>
