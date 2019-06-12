$(document).ready(function(){
    // 对于发布房源，只有认证后的用户才可以，所以先判断用户的实名认证状态
    $.get("/api/v1.0/user/auth", function(resp){
        if ("4101" == resp.errno) {
            // 用户未登录
            location.href = "/login.html";
        } else if ("0" == resp.errno) {
            // 未认证的用户，在页面中展示 "去认证"的按钮
            if (!(resp.data.real_name && resp.data.id_card)) {
                $(".auth-warn").show();
                return;
            }
            // 已认证的用户，请求其之前发布的房源信息
            $.get("/api/v1.0/user/houses", function(resp){
                if ("0" == resp.errno) {
                    $("#houses-list").html(template("houses-list-tmpl", {houses:resp.data.houses}));
                } else {
                    $("#houses-list").html(template("houses-list-tmpl", {houses:[]}));
                }
            });
        }
    });
})