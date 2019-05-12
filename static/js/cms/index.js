'use strict';

var SUBMIT_PATH = "/api/users/login";

window.onload = function init() {
  function addSubmitListener() {
    var loginForm = document.getElementById("login");
    loginForm.addEventListener("submit", onSubmit);
  }

  function onSubmit(event) {
    event.preventDefault();
    var loginRequest = buildLoginRequest();
    Net.postJSON(SUBMIT_PATH, loginRequest, function(response) {
      console.log("login response");
      console.log(response);
    }, function(err) {
      console.error(err);
    })
    return false;
  }

  function buildLoginRequest() {
    var usernameElem = document.getElementById("username");
    var passwordElem = document.getElementById("password");

    return {
      username: usernameElem.value,
      password: passwordElem.value
    };
  }

  addSubmitListener();
}
