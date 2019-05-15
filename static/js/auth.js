'use strict';

var Auth = (function() {
  var LOCAL_STORAGE_TOKEN_KEY = "token"

  function setToken(token) {
    localStorage.setItem(LOCAL_STORAGE_TOKEN_KEY, token);
  }

  function getToken() {
    return localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY);
  }

  function loggedIn() {
    return getToken() != null;
  }

  function logout() {
    localStorage.removeItem(LOCAL_STORAGE_TOKEN_KEY);
  }
  
  return {
    setToken: setToken,
    getToken: getToken,
    loggedIn: loggedIn,
    logout: logout,
  };
})()
