<div class="container">
    <div class="h2 text-center">登 录</div>
    <form action="/sign_in" method="post">
        <div class="form-group">
            <label for="userName">用户名</label>
            <input class="form-control" type="text" value="{{.Name}}" id="userName" name="Name" aria-describedby="nameHelp" />
        </div>
        <div class="form-group">
            <label for="userPassword">密码</label>
            <input type="password" class="form-control" name="Password" id="userPassword" />
        </div>
        <button type="submit" class="btn btn-primary form-control font-weight-bold">登 录</button>
    </form>
</div>
