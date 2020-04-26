<div class="container">
    <div class="h2 text-center">修改密码</div>
    <form action="reset_password" method="post">
        <div class="form-group">
            <label for="userName">用户名</label>
            <input class="form-control" type="text" value="{{.Name}}" id="userName" name="Name" aria-describedby="nameHelp" disabled />
        </div>
        <div class="form-group">
            <label for="userPassword">密码</label>
            <input type="password" class="form-control" name="Password" id="userPassword" />
        </div>
        <div class="form-group">
            <label for="userPasswordConfirmation">确认密码</label>
            <input type="password" class="form-control" name="PasswordConfirmation" id="userPasswordConfirmation" />
        </div>
        <button type="submit" class="btn btn-primary form-control font-weight-bold">提交</button>
    </form>
</div>
