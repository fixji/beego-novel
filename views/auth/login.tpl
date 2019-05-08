<div>
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card">
                <div class="card-header">登录</div>

                <div class="card-body">
                    <form method="POST" action="/user/login" id="sub_login">
                        <div class="form-group row">
                            <label for="name" class="col-md-4 col-form-label text-md-right">用户名</label>

                            <div class="col-md-6">
                                <input id="name" type="text" class="form-control" name="username" value="" required autofocus >

                            </div>
                        </div>

                        <div class="form-group row">
                            <label for="password" class="col-md-4 col-form-label text-md-right">密码</label>

                            <div class="col-md-6">
                                <input id="password" type="password" class="form-control" name="password" required>
                            </div>
                        </div>

                        <div class="form-group row mb-0">
                            <div class="col-md-8 offset-md-4">
                                <button type="submit" class="btn btn-primary">
                                    登录
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    $(function(){
        var btnClick = 0;
        $('.btn-primary').on('click', function () {
            if (btnClick == 1) return;
            btnClick = 1;
            event.preventDefault();
            $.ajax({
                url:"/user/login",
                type:"post",
                headers: {
                    'X-Xsrftoken': $('meta[name="xsrf-token"]').attr('content')
                },
                data: $('#sub_login').serialize(),
                success:function(data){
                    btnClick = 0;
                    if (data.code != 200) {
                        alert(data.message);
                        return;
                    }
                    window.location.href = "/";
                },
                error:function(e){
                    btnClick = 0;
                    alert('服务器错误!');
                }
            });
        });
    });
</script>