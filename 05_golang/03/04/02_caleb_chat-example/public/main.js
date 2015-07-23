// make an api request
function apiRequest(method, endpoint, data, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, "/api/" + endpoint);
    if (method === "POST") {
        // if posting data, send it
        // this part is for SENDING ajax request
        xhr.send(JSON.stringify(data));
    } else {
        xhr.send(null);
    }
    xhr.onreadystatechange = function(evt) {
        // check AJAX response (readyState ===4)
        // this part is for RECEIVING ajax response
        if (xhr.readyState === 4) {
            if (Math.floor(xhr.status/100) === 2) {
                // all 200 statuses
                if (callback) {
                    callback(JSON.parse(xhr.responseText), null);
                }
                //
            } else {
                // all other statuses
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

// send a message
function sendMessage(text, callback) {
  apiRequest("POST", "messages", {
    "Text": text
  }, callback);
}

// when a message is received
function onMessage(message) {
  var el = document.getElementById("messages");
  var p = document.createElement("p");
  p.textContent = message.Text;
  el.appendChild(p);
}

(function() {

// hook up text input
var controls = document.getElementById("controls");
controls.addEventListener("submit", function(evt) {
  evt.preventDefault();
  var textInput = document.getElementById("text-input");
  var text = textInput.value;
  sendMessage(text, function(msg, err) {
    if (err) {
      alert(err);
    }
  });
  textInput.value = "";
}, false);

})();
