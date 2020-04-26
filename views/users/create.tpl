<div class="container">
    <div class="h2 text-center">注 册</div>
    <form action="/sign_up" method="post">
        <div class="form-group">
            <label for="userName">用户名</label>
            <input class="form-control" type="text" value="{{.Name}}" id="userName" name="Name" aria-describedby="nameHelp" />
            <small id="nameHelp" class="form-text text-muted">注册后不可更改！</small>
        </div>
        <div class="form-group">
            <label for="userEmail">邮件</label>
            <input class="form-control" type="text" value="{{.Email}}" id="userEmail" name="Email" aria-describedby="emailHelp" />
            <small id="emailHelp" class="form-text text-muted">我们一定会保密！</small>
        </div>
        <div class="form-group">
            <label for="userPhone">手机</label>
            <input class="form-control" type="text" value="{{.Phone}}" id="userPhone" name="Phone" aria-describedby="phoneHelp" />
            <small id="phoneHelp" class="form-text text-muted">我们一定保密！</small>
        </div>
        <div class="form-group">
            <label>性别： </label>
            <div class="form-check form-check-inline">
                {{if eq (or .Gender "n") "man" }}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender1" value="man" checked />
                {{else}}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender1" value="man" />
                {{end}}
                <label class="form-check-label" for="userGender1">男</label>
            </div>
            <div class="form-check form-check-inline">
                {{if eq (or .Gender "n") "female" }}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender2" value="female" checked />
                {{else}}
                    <input class="form-check-input" type="radio" name="Gender" id="userGender2" value="female" />
                {{end}}
                <label class="form-check-label" for="userGender2">女</label>
            </div>
        </div>
        <div class="form-group">
            <label for="userPassword">密码</label>
            <input type="password" class="form-control" name="Password" id="userPassword" />
        </div>
        <div class="form-group">
            <label for="userPasswordConfirmation">确认密码</label>
            <input type="password" class="form-control" name="PasswordConfirmation" id="userPasswordConfirmation" />
        </div>
        <div class="form-group">
            <label for="userComment">个人描述</label>
            <textarea class="form-control" name="Comment" id="userComment" rows=5>{{.Comment}}</textarea>
        </div>

        <button type="submit" class="btn btn-primary form-control font-weight-bold">注 册</button>
    </form>
</div>
