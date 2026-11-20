(function () {
  'use strict';

  var greeting = document.getElementById('greeting');
  var nameInput = document.getElementById('name');

  if (greeting && nameInput) {
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
  }

  document.querySelectorAll('.security-demo').forEach(function (demo) {
    demo.querySelectorAll('[data-show]').forEach(function (button) {
      button.addEventListener('click', function () {
        var target = document.getElementById(button.dataset.show);
        var expanded = button.getAttribute('aria-expanded') === 'true';
        target.hidden = expanded;
        button.setAttribute('aria-expanded', String(!expanded));
      });
    });

    var status = demo.querySelector('.autofix-status');
    demo.querySelector('[data-autofix]').addEventListener('click', function () {
      var confirmed = window.confirm(
        'This static demo has no protected agent integration. No changes will be made. Continue to see the integration status?'
      );
      if (confirmed) {
        status.textContent =
          'Not started: agentic autofix is unsupported here. No request was sent and no files changed. ' +
          'Review current alerts and use an authorized GitHub workflow; verify each finding independently.';
      }
    });
  });
})();
