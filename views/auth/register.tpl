<div>
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card">
                <div class="card-header">注册账号</div>

                <div class="card-body">
                    <form method="POST" action="/user/register" id="sub_register">
                        <div class="form-group row">
                            <label for="reg_username" class="col-md-4 col-form-label text-md-right">用户名</label>

                            <div class="col-md-6">
                                <input id="reg_username" type="text" class="form-control" name="reg_username" value="" required autofocus autocomplete="off">
                            </div>
                        </div>

                        <div class="form-group row">
                            <label for="reg_password" class="col-md-4 col-form-label text-md-right">密码</label>

                            <div class="col-md-6">
                                <input id="reg_password" type="password" class="form-control" name="reg_password" required autocomplete="off">
                            </div>
                        </div>

                        <div class="form-group row">
                            <label for="password-confirm" class="col-md-4 col-form-label text-md-right">确认密码</label>
                            <div class="col-md-6">
                                <input id="password-confirm" type="password" class="form-control" name="password_confirmation" required>
                            </div>
                        </div>

                        <div class="form-group row mb-0">
                            <div class="col-md-6 offset-md-4">
                                <button type="submit" class="btn btn-primary">
                                    注册
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
                url:"/user/register",
                type:"post",
                headers: {
                    'X-Xsrftoken': $('meta[name="xsrf-token"]').attr('content')
                },
                data: $('#sub_register').serialize(),
                success:function(data){
                    btnClick = 0;
                    if (data.code != 200) {
                        alert(data.message);
                        return;
                    }
                    window.location.href = "/user/login";
                },
                error:function(e){
                    btnClick = 0;
                    alert('服务器错误!');
                }
            });
        });
    });
</script>