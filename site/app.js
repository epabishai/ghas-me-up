(function () {
  'use strict';

  var greeting = document.getElementById('greeting');
  var nameInput = document.getElementById('name');

  function render() {
    var name = window.location.hash.slice(1);
    try {
      name = decodeURIComponent(name);
    } catch (error) {
      if (!(error instanceof URIError)) throw error;
      // Display malformed URL fragments as literal text.
    }
    greeting.textContent = 'Hello, ' + (name || 'friend') + '!';
  }

  document.getElementById('hello-form').addEventListener('submit', function (event) {
    event.preventDefault();
    window.location.hash = encodeURIComponent(nameInput.value);
  });

  window.addEventListener('hashchange', render);
  render();
})();
