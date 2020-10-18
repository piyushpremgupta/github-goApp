$(document).ready(function () {
    $("li:not(:contains('Login'))").hide()
    $(".progress").hide()
    if (window.history && window.history.pushstate) {
        window.history.pushstate('forward', null, './#forward');
        $(window).on('popstate', function () {
            window.location.replace("/login");
        });
    }
    $("#toggle_icon").on('click', function (){
        var type = $("#password").attr("type")
        if (type == "password") {
            $("#password").attr("type", "text")
            $("#toggle_icon").removeClass().addClass("glyphicon glyphicon-eye-open")
        } else {
            $("#password").attr("type", "password")
            $("#toggle_icon").removeClass().addClass("glyphicon glyphicon-eye-close")
        }
    })
    
    $("[type='submit']").on("click", function(e) {
        e.preventDefault()
        frm = $(this).closest("form")
        $("body").fadeTo("fast", "0.6", function() {
            $(".progress").show(function() {
                var jsonData = {UserName:$("#username").val(),Password:$("#password").val()}
                console.log(jsonData)
                Url = "/auth"
                Method = "post"
                Payload = JSON.stringify(jsonData)
                User = jsonData.UserName
                Password = jsonData.Password
                Response = MakeCall(Url, Method, Payload, User, Password)
                if (Response.hasOwnProperty("Authenticate")) {
                    window.location.replace("/github");
                }
                else {
                    $("body").fadeTo(3000, "1")
                    $(".progress").hide()
                    $("").remove()
                    $("form").append("<p>" + Response.status + "</p>")
                }
            });
        });
    });
})
