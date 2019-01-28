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

    return {
        getJSON: getJSON,
    }
})()
