// make an api request
function apiRequest(method, endpoint, data, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, "/" + endpoint);
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

// send data to server
function sendMessage(obj, callback) {
    apiRequest("POST", "api", obj, callback);
}

// display data received from server
function onMovies(movies) {
    var el = document.getElementById("messages");
    var p = document.createElement("p");
    p.textContent = message.Text;
    el.appendChild(p);
}

// IIFE
(function() {
// hook up text input
    var postForm = document.querySelector("post-form");
    postForm.addEventListener("submit", function(evt) {
        evt.preventDefault();
        var movieTitle = document.querySelector("movie-title");
        var movieDescrip = document.querySelector("movie-description");
        var movie = {
            "Title":movieTitle,
            "Description":movieDescrip
        };
        sendMessage(movie, function(movies, err) {
                if (err) {
                    alert(err);
                }
            // call function that handles incoming obj of movies
            // put on page
            var title = document.createElement('h1');
            var title = document.createElement('p');
            for (var i = 0; i < length(movies); i++) {

            }
            title.innerText =
        });
        movieTitle.value = "";
        movieDescrip.value = "";
    }, false);

    apiRequest("GET", "api", null, function(movies, err) {
        if (err) {
            alert(err);
            return;
        }
        // call function that handles incoming obj of movies
        // put on page
    });
})();