'use strict';

var SUBMIT_PATH = "/api/users/login";
var LOGIN_PATH = "/static/html/cms/login.html";
var EDIT_PATH = "/static/html/cms/edit.html";

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
      addSubmitLoginListener();
    }, function (err) {
      console.error("failed to fetch login");
      console.error(err);
    })
  }

  function renderCMS(container) {
    Net.get(EDIT_PATH, {}, function(response) {
      container.innerHTML = response;
      var logoutButton = document.getElementById("logout");
      var submitButton = document.getElementById("submit-edit");

      logoutButton.addEventListener("click", onLogout);
      submitButton.addEventListener("click", onSubmitEdit)
    }, function(err) {
      console.error(err);
    });
  }

  function addSubmitLoginListener() {
    var loginForm = document.getElementById("login");
    loginForm.addEventListener("submit", onSubmitLogin);
  }

  function onSubmitEdit(event) {
    console.log("edit submitted");
  }

  function onLogout(event) {
    Auth.logout();
    render();
  }

  function onSubmitLogin(event) {
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

  render();
}
