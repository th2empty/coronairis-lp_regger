function register_user() {
    let access_token = document.getElementById("access_token_field").value;
    if (access_token.length === 0) {
        alert("Прежде чем нажимать на кнопочки токен вставь...")
        return
    }

    document.getElementById("msg").style.display = "none"

    let httpReq = new XMLHttpRequest();
    httpReq.onreadystatechange = function() {
        if (httpReq.readyState !== 4) return;

        if (httpReq.status !== 200) {
            document.getElementById("msg").style.display = "flex"
            let err_msg = JSON.parse(httpReq.responseText)
            document.getElementById("msg").style.color = "red"
            document.getElementById("msg").innerHTML = `Error: ${err_msg.message}`
        } else {
            let res = JSON.parse(httpReq.responseText)
            document.getElementById("msg").style.display = "flex"
            document.getElementById("msg").style.color = "green"
            document.getElementById("msg").innerHTML = `${res.message}`
        }
    }
    httpReq.open("POST", "http://192.168.1.147:9000/api/v1/register", true); // change url before deploy
    httpReq.setRequestHeader('Content-type', 'application/json');
    httpReq.send(JSON.stringify({
        "access_token": access_token
    }));
}

function update_user() {
    let access_token = document.getElementById("access_token_field").value;
    if (access_token.length === 0) {
        alert("Прежде чем нажимать на кнопочки токен вставь...")
        return
    }

    document.getElementById("msg").style.display = "none"

    let httpReq = new XMLHttpRequest();
    httpReq.onreadystatechange = function() {
        if (httpReq.readyState !== 4) return;

        if (httpReq.status !== 200) {
            document.getElementById("msg").style.display = "flex"
            let err_msg = JSON.parse(httpReq.responseText)
            document.getElementById("msg").style.color = "red"
            document.getElementById("msg").innerHTML = `Error: ${err_msg.message}`
        } else {
            let res = JSON.parse(httpReq.responseText)
            document.getElementById("msg").style.display = "flex"
            document.getElementById("msg").style.color = "green"
            document.getElementById("msg").innerHTML = `${res.message}`
        }
    }
    httpReq.open("POST", "http://192.168.1.147:9000/api/v1/update", true); // change url before deploy
    httpReq.setRequestHeader('Content-type', 'application/json');
    httpReq.send(JSON.stringify({
        "access_token": access_token
    }));
}