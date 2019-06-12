function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

// 点击推出按钮时执行的函数
function logout() {
    $.ajax({
        url: "/api/v1.0/session",
        type: "delete",
        headers: {
            "X-CSRFToken": getCookie("csrf_token")
        },
        dataType: "json",
        success: function (resp) {
            if ("0" == resp.errno) {
                location.href = "/index.html";
            }
        }
    });
}

$(document).ready(function(){
    $.get("/api/v1.0/user", function(resp){
        // 用户未登录
        if ("4101" == resp.errno) {
            location.href = "/login.html";
        }
        // 查询到了用户的信息
        else if ("0" == resp.errno) {
            $("#user-name").html(resp.data.name);
            $("#user-mobile").html(resp.data.mobile);
            if (resp.data.avatar_url) {
                $("#user-avatar").attr("src", resp.data.avatar_url);
            }

        }
    }, "json");
});
