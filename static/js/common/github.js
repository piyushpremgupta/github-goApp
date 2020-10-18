$(document).ready(function () {
    $("li:contains('Login')").hide()
    $.get('/template/common/popup.html', function (data) {
        $("body").append(data)
    })


    baseUrl="http://localhost:5001"
    function GithubOperations(form_data, operation) {
        if (operation=="createbranch"){
            url=baseUrl+"/CreateBranch/"+form_data["organisation"]+"/"+form_data["repository"]+"/"+form_data["ftrbrnch"]
            method="post"
            res=MakeCall(url,method)
            tableparams=CreateDatatablesParameters(res,operation)
            ShowModal(tableparams,[res])
        }

        if (operation=="modifyfile"){
            url=baseUrl+"/ChangeFileContent/"+form_data["organisation"]+"/"+form_data["repository"]+"/"+form_data["filepath"]
            method="put"
            json_payload={FileContent:form_data["filecontent"],TargetBranch:form_data["ftrbrnch"]}
            json_payload=JSON.stringify(json_payload)
            console.log(url,method,json_payload)
            res=MakeCall(url,method,json_payload)
            tableparams=CreateDatatablesParameters(res,operation)
            console.log(tableparams)
            ShowModal(tableparams,[res])
        }

        if (operation=="createpullrequest"){
            url=baseUrl+"/CreatePullRequest/"+form_data["organisation"]+"/"+form_data["repository"]
            method="post"
            json_payload={Title:form_data["title"],Head:form_data["head"],Base:form_data["base"]}
            json_payload=JSON.stringify(json_payload)
            res=MakeCall(url,method,json_payload)
            tableparams=CreateDatatablesParameters(res,operation)
            ShowModal(tableparams,[res])
        }

    }

    $("[type='submit']").on("click", function (e) {
        e.preventDefault()
        frm = $(this).closest("form")
        operation = frm.prev()[0].outerText.toLowerCase()
        input_elements = frm.find("input")
        formData = {}
        for (var i = 0; i < input_elements.length; i++) {
            formData[input_elements[i].name] = input_elements[i].value

        }
        
        $.get("/template/common/progress.html", function (data) {
            $(frm).append(data);
        })
        $("body").fadeTo("fast", "0.6", function () {
            $(".progress").show(function () {
                console.log(formData, operation)
                GithubOperations(formData, operation)
            })
        })
        $("#loader").remove()

    })
})



