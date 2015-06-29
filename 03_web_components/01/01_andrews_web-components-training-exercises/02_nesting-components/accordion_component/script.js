var accordions = Array.prototype.slice.call(document.querySelectorAll('.accordion-component'));
accordions.forEach(function(accordion) {
  var allContent = Array.prototype.slice.call(accordion.querySelectorAll('.content'));
  var sections = Array.prototype.slice.call(accordion.querySelectorAll('.section'));
  sections.forEach(function(section) {
    var content = section.querySelector('.content');
      console.count();
    section.querySelector('h2').addEventListener('click', function() {
      var closed = content.classList.contains('closed');
      allContent.forEach(function(content) {
        content.classList.add('closed');
      });
      closed ? content.classList.remove('closed') : null;
    });
  });
});