var ref = new Firebase("https://scorching-heat-8600.firebaseio.com/");
// on is like addEventListener in firebase
ref.on('value', function (snapshot) {
    var data = snapshot.val();
    console.log(data);
    for (var i in data.speakers) {
        console.log(data.speakers[i]);
    }
    var template = document.querySelector('#displayInfoTemplate').innerHTML;
    var html = Mustache.to_html(template, data);
    document.querySelector('#putDataHere').innerHTML = html;
});