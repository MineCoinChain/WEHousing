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

$(document).ready(function () {
    // 在页面加载是向后端查询用户的信息
    $.get("/api/v1.0/user", function(resp){
        // 用户未登录
        if ("4101" == resp.errno) {
            location.href = "/login.html";
        }
        // 查询到了用户的信息
        else if ("0" == resp.errno) {
            $("#user-name").val(resp.data.name);
            if (resp.data.avatar) {
                $("#user-avatar").attr("src", resp.data.avatar);
            }
        }
    }, "json");

    // 管理上传用户头像表单的行为
    $("#form-avatar").submit(function (e) {
        // 禁止浏览器对于表单的默认行为
        e.preventDefault();
        $(this).ajaxSubmit({
            url: "/api/v1.0/user/avatar",
            type: "post",
            headers: {
                "X-CSRFToken": getCookie("csrf_token"),
            },
            dataType: "json",
            success: function (resp) {
                if (resp.errno == "0") {
                    // 表示上传成功， 将头像图片的src属性设置为图片的url
                    $("#user-avatar").attr("src", resp.data.avatar_url);
                } else if (resp.errno == "4101") {
                    // 表示用户未登录，跳转到登录页面
                    location.href = "/login.html";
                } else {
                    alert(resp.errmsg);
                }
            }
        });

    });
    $("#form-name").submit(function(e){
        e.preventDefault();
        // 获取参数
        var name = $("#user-name").val();

        if (!name) {
            alert("请填写用户名！");
            return;
        }
        $.ajax({
            url:"/api/v1.0/user/name",
            type:"PUT",
            data: JSON.stringify({name: name}),
            contentType: "application/json",
            dataType: "json",
            headers:{
                "X-CSRFTOKEN":getCookie("csrf_token")
            },
            success: function (data) {
                if ("0" == data.errno) {
                    $(".error-msg").hide();
                    showSuccessMsg();
                } else if ("4001" == data.errno) {
                    $(".error-msg").show();
                } else if ("4101" == data.errno) {
                    location.href = "/login.html";
                }
            }
        });
    })
})

