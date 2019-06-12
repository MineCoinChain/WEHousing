function showSuccessMsg() {
    $('.popup_con').fadeIn('fast', function() {
        setTimeout(function(){
            $('.popup_con').fadeOut('fast',function(){}); 
        },1000) 
    });
}


function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

$(document).ready(function(){
    // 查询用户的实名认证信息
    $.get("/api/v1.0/user/auth", function(resp){
        // 4101代表用户未登录
        if ("4101" == resp.errno) {
            location.href = "/login.html";
        }
        else if ("0" == resp.errno) {
            // 如果返回的数据中real_name与id_card不为null，表示用户有填写实名信息
            if (resp.data.real_name && resp.data.id_card) {
                $("#real-name").val(resp.data.real_name);
                $("#id-card").val(resp.data.id_card);
                // 给input添加disabled属性，禁止用户修改
                $("#real-name").prop("disabled", true);
                $("#id-card").prop("disabled", true);
                // 隐藏提交保存按钮
                $("#form-auth>input[type=submit]").hide();
            }
        }
    }, "json");

    // 管理实名信息表单的提交行为
    $("#form-auth").submit(function(e){
        e.preventDefault();
        // 如果用户没有填写完整，展示错误信息
        if ($("#real-name").val()=="" || $("#id-card").val() == "") {
            $(".error-msg").show();
            return
        }

        // 将表单的数据转换为json字符串
        var data = {};
        $(this).serializeArray().map(function(x){data[x.name] = x.value;});
        var jsonData = JSON.stringify(data);

        // 向后端发送请求
        $.ajax({
            url:"/api/v1.0/user/auth",
            type:"post",
            data: jsonData,
            contentType: "application/json",
            dataType: "json",
            headers: {
                "X-CSRFTOKEN": getCookie("csrf_token")
            },
            success: function (resp) {
                if (0 == resp.errno) {
                    $(".error-msg").hide();
                    // 显示保存成功的提示信息
                    showSuccessMsg();
                    $("#real-name").prop("disabled", true);
                    $("#id-card").prop("disabled", true);
                    $("#form-auth>input[type=submit]").hide();
                }else {
                    // $(".error-msg").show();
                    $(".error-msg").html('<i class="fa fa-exclamation-circle"></i>信息不符,验证失败,请重新填写!</div>').show();
                }
            }
        });
    })

})