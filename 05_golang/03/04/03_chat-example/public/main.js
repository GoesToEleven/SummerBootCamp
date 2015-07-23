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
      // check all 200 statuses
      // this part is for RECEIVING ajax response
    if (xhr.readyState === 4) {
      if (Math.floor(xhr.status/100) === 2) {
        if (callback) {
          callback(JSON.parse(xhr.responseText), null);
        }
          //
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

// when you first join the chat room, make a request to the channel
function joinChat(callback) {
    apiRequest("POST", "channels", {}, callback);
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
joinChat(function(msg, err){
    if (err) {
        alert(err);
        return
    }
    //alert(msg);
     var channel = new goog.appengine.Channel(msg);
     var socket = channel.open();
     socket.onmessage = function(msg){
         var data = JSON.parse(msg.data)
         onMessage(data);
     };
     socket.onerror = function() {
         alert("this is from joinChat method in the js file")
     };
    // or this:
    //var socket = channel.open({
    //    onopen: function() {
    //        console.log("OPEN", arguments);
    //    },
    //    onmessage: function(msg) {
    //        var data = JSON.parse(msg.data);
    //        alert("Message Received: " + data);
    //    },
    //    onerror: function() {
    //        console.log("ERROR", arguments);
    //    },
    //    onclose: function() {
    //        console.log("CLOSE", arguments);
    //    }
    //});
});

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

// receive token

function receiveToken() {
    apiRequest("GET", "channels", null, function(msg, err) {

    })
}
