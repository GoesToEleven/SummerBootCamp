// sending data to server
var modal = document.querySelector('#openModal');
modal.addEventListener('click', function(e){
    if ((e.target.id === 'modal-submit') && (e.target.textContent !== '')) {
        var msg = e.target.textContent;
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/post");
        var json = JSON.stringify({Message: msg});
        console.log(e);
        console.log(e.target);
        console.log(e.target.id);
        console.log(e.target.textContent);
        console.log(msg);
        console.log(json);
        xhr.send(json);
        // clear entry
        e.target.textContent = '';
    }
}, false);


var opensModal = document.querySelector('#opens-modal');
opensModal.addEventListener('click', function(e){
    window.setTimeout(function(){
        var tweet = document.querySelector('#modal-tweet');
        tweet.textContent = '';
        console.log(tweet)
    }, 100);
}, false);