var myEls = document.querySelectorAll('.email-item');
//console.log(myEls);
//console.log(myEls.length);
var myEmailCount = document.querySelector('.email-count');
myEmailCount.innerHTML = '(' + myEls.length + ')';