function SplitCookie(cookie_arr) {
    var auth_key = "";
    for (i = 0; i < cookie_arr.length; i++) {
        var result = cookie_arr[i].match("Token") || [];
        if (result.length != 0) {
            auth_key = cookie_arr[i].split("=")[1];
            return auth_key
        }
        else {
            return null
        }
    }
}


function ConvertFormToJSON(frm) {
    function TraverseParent(element) {
        if (element) {
            var pattern = new RegExp("<body>|<html>")
            if (pattern.test(element.outerHTML)) {

                pattern = new RegExp("^<label")

                if (pattern.test(element.children[0].outerHTML)) {

                    pattern = new RegExp(">.*<")

                    key = key + pattern.exec(element.children(0).outerHTML)[0].replace("<", "").replace(">", "").tolowerCase() + ","
                }
            }
            TraverseParent(element.parentElement)
        }
        console.log(key)
        return key
    }


    function createNestedobject(base, names, value) {
        var lastName = arguments.length === 3 ? names.pop() : false;
        for (var i = 0; i < names.length; i++) {
            base = base[names[i]] = base[names[i]] || {};
        }
        if (lastName) base = base[lastName] = value;
    }

    input_elements = frm.find("input")
    console.log(input_elements)
    form_data = {}
    operation = frm.prev()[0].outerText.toLowerCase()
    for (var j = 0; j < input_elements.length; j++) {
        key = ""
        key = key + TraverseParent(input_elements[j].parentElement)
        key_array = key.split(",")
        input_element_data = {}
        createNestedobject(form_data, key_array.reverse(), input_elements[j].value)
    }
    return [form_data[""], operation];
};





function CreateDatatablesParameters(JsonResponse, Tablename) {
    Tablename = Tablename || "Example"
    var HeaderArray = Object.keys(JsonResponse)
    var Thead = ""
    var cols = []
    for (i = 0; i < HeaderArray.length; i++) {
        var ele = new Object();
        ele["data"] = HeaderArray[i]
        cols = cols.concat(ele)
    }
    for (var i = 0; i < HeaderArray.length; i++) {
        Thead = Thead + "<th>" + HeaderArray[i].charAt(0).toUpperCase() + HeaderArray[i].slice(1) + "</th>"
    }
    var TableStructure = '<table id="' + Tablename + '" class="table table-striped table-bordered table-hover table-condensed" width="100%"><thead><tr>' + Thead + '</tr></thead><tbody></tbody></table>'
    return [TableStructure, cols, Tablename]
};


function MakeCall(Url, Method, Payload, User, Password) {
    User = User || null
    Password = Password || null
    Payload = Payload || null
    var cookie_arr = document.cookie.split(";");
    var jqXHR = $.ajax({
        url: Url,
        type: Method,
        data: Payload,
        headers: {
            "Authorization": 'token ' + SplitCookie(cookie_arr) || btoa(User + '.' + Password), "Content-Type": "application/json",
            'Origin': 'http://localhost:5001'
        },
        async: false,
        success: function (response) { },
        error: function (xhr, status, errortext) {
            console.log(errortext)
        }
    });


    var ct = jqXHR.getResponseHeader("content-type") || "";
    if (ct.indexOf('html') > -1) {
        return jqXHR.responseText;
    }
    if (ct.indexOf('application/json') > -1) {
        if (jqXHR.responseJSON.length > 0 || Object.keys(jqXHR.responseJSON).length > 0) {
            console.log(jqXHR.responseJSON)
            return jqXHR.responseJSON;
        }
        else {
            return [{ Output: "No Rows Returned" }]
        }
    }
}



function ShowModal(tableparams, tbldata) {
    $(".modal-body").append("<h4>" + tableparams[2] + "</h4>" + tableparams[0])
    $("#" + tableparams[2]).DataTable({
        "pageLength": 50,
        "buttons": ['csv'],
        dom: 'Bfrtip',
        "data": tbldata,
        "columns": tableparams[1],
        "bDestroy": true,
        "scrollx": true,
        "bAutoWidth": false,
        "bPaginate": false,
        "ÞInfo": false,
    });



    $("body").fadeTo("fast", "1")
    $(".progress").hide()
    $('#popup').modal({ "show": true });
    $('#popup').on('shown.bs.modal', function () {
        $("th:first-child").click();
    });
    $("#popup").on('hidden.bs.modal', function () {
        $(".modal-body").empty();
    });
}


