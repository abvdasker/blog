'use strict';

var SUBMIT_PATH = "/api/users/login";
var LOGIN_PATH = "/static/html/cms/login.html";

// if user is already logged in, render cms
// else render login

window.onload = function init() {
  var container = document.getElementById("container");

  function render() {
    if (Auth.loggedIn()) {
      renderCMS(container);
    } else {
      renderLogin(container);
    }
  }

  function renderLogin(container) {
    Net.get(LOGIN_PATH, {}, function(response) {
      container.innerHTML = response;
      addSubmitListener();
    }, function (err) {
      console.error("failed to fetch login");
      console.error(err);
    })
  }

  function addSubmitListener() {
    var loginForm = document.getElementById("login");
    loginForm.addEventListener("submit", onSubmit);
  }

  function onSubmit(event) {
    event.preventDefault();
    var loginRequest = buildLoginRequest();
    Net.postJSON(SUBMIT_PATH, loginRequest, function(response) {
      Auth.setToken(response.token);
      render();
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

  function renderCMS(container) {
    console.log("logged in. Rendering CMS");
  }

  render();
  //addSubmitListener();
}
