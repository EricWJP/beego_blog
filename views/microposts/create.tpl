<div class="container">
    <div class="h1 text-center">新博客</div>
    <form action="create" method="post">
        <div class="form-group">
            <label for="micropostsTitle">标题：</label>
            <input class="form-control" value="{{.Title}}" type="text" name="Title" id="micropostsTitle"  />
        </div>
        <div class="form-group">
            <label for="micropostsContent">内容：</label>
            <textarea class="form-control" name="Content" id="micropostsContent" rows=5>{{.Content}}</textarea>
        </div>
        <button type="submit" class="btn btn-primary form-control font-weight-bold">提交</button>
    </form>
</div>
