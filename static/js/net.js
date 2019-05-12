'use strict';

var Net = (function() {
  var DONE = 4;
  var JSON_CONTENT_TYPE = "application/json;charset=UTF-8";

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

  function get(url, headers, onSuccess, onErr) {
    var request = newRequest(headers, function(request) {
      request.open("GET", url, true);
    }, onSuccess, onErr);
    request.send();
  }

  function post(url, headers, body, onSuccess, onErr) {
    var request = newRequest(headers, function(request) {
      request.open("POST", url, true);
    }, onSuccess, onErr);
    request.send(body);
  }

  function newRequest(headers, openRequest, onSuccess, onErr) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = onResponse(onSuccess, onErr);
    openRequest(request);
    setHeaders(request, headers);
    maybeSetAuthHeader(request);
    return request;
  }

  function setHeaders(request, headers) {
    Object.keys(headers).forEach(function(key) {
      var value = headers[key];
      request.setRequestHeader(key, value);
    });
  }

  function maybeSetAuthHeader(request) {
    var token = Auth.getToken();
    if (token) {
      var headerValue = `Bearer ${token}`;
      request.setRequestHeader("Authorization", headerValue);
    }
  }

  function getJSON(url, onSuccess, onErr) {
    get(
      url,
      { "Content-Type": JSON_CONTENT_TYPE },
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
      { "Content-Type": JSON_CONTENT_TYPE },
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
    get: get,
    getJSON: getJSON,
    postJSON: postJSON,
  }
})()
