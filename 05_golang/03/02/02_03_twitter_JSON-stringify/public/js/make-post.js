// sending data to server
var modal = document.querySelector('#openModal');
modal.addEventListener('click', function(e){
    if ((e.target.id === 'modal-submit') && (e.target.textContent !== '')) {
        var modalTweet = document.querySelector('#modal-tweet')
        var msg = modalTweet.value;
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/tweet");
        var json = JSON.stringify({Message: msg});
        console.log('Got to JSON.stringify')
        console.log('e: ', e);
        console.log('target: ', e.target);
        console.log('id: ',e.target.id);
        console.log('value: ',modalTweet.value);
        console.log('msg: ',msg);
        console.log('json: ',json);
        xhr.send(json);
        // clear entry
        modalTweet.value = '';
    }
}, false);

var opensModal = document.querySelector('#opens-modal');
opensModal.addEventListener('click', function(e){
    window.setTimeout(function(){
        var tweet = document.querySelector('#modal-tweet');
        tweet.value = '';
        tweet.focus();
    }, 100);
}, false);