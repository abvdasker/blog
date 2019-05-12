'use strict';

var Net = (function() {
  var DONE = 4;
  var ARTICLES = "/api/articles";

  function onResponse(onSuccess, onErr) {
    return function() {
      if (this.readyState !== DONE) {
        return;
      }

      if (this.status === 200 && onSuccess) {
        onSuccess(this.responseText);
        return;
      }

      if (onErr) {
        onErr(this.responseText);
      }
    }
  }

  function get(url, onSuccess, onErr) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = onResponse(onSuccess, onErr);
    request.open("GET", url, true);
    request.send();
  }

  function post(url, body, onSuccess, onErr) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = onResponse(onSuccess, onErr);
    request.open("POST", url, true);
    request.send(body);
  }

  function getJSON(url, onSuccess, onErr) {
    get(
      url,
      function(responseText) {
        if (onSuccess) {
          var response = JSON.parse(responseText);
          onSuccess(response);
        }
      },
      function(responseText) {
        if (onErr) {
          onErr(responseText);
        }
      }
    );
  }

  function postJSON(url, body, onSuccess, onErr) {
    post(
      url,
      JSON.stringify(body),
      function(responseText) {
        if (onSuccess) {
          var response = JSON.parse(responseText);
          onSuccess(response);
        }
      },
      function(responseText) {
        if (onErr) {
          onErr(responseText);
        }
      }
    );
  }

  return {
    getJSON: getJSON,
    postJSON: postJSON,
  }
})()
