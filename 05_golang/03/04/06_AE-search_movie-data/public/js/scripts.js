// make an api request
function apiRequest(method, endpoint, data, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, "/api/" + endpoint);
    xhr.send(JSON.stringify(data));
    xhr.onreadystatechange = function(evt) {
        if (xhr.readyState === 4) {
            if (Math.floor(xhr.status/100) === 2) {
                if (callback) {
                    callback(JSON.parse(xhr.responseText), null);
                }
                // any status other than 2XX, else ...
            } else {
                var msg;
                try {
                    msg = JSON.parse(xhr.responseText);
                } catch(e) {
                    msg = xhr.responseText;
                }
                if (callback) {
                    callback(null, msg);
                }
            }
        }
    };
}