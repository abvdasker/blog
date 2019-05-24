'use strict';

var DOM = {
  byID: function(id) {
    return document.getElementById(id);
  },
  empty: function(elem) {
    var removed = 0;
    while (elem.firstChild) {
      elem.removeChild(elem.firstChild);
      removed++;
    }
    return removed;
  }
};

var I18n = {
  getLang: function() {
    if (navigator.languages != undefined) 
      return navigator.languages[0]; 
    else 
      return navigator.language;
  },
  formatDate: function(date) {
    return date.toLocaleDateString(I18n.getLang());
  },
  formatDateVerbose: function(date) {
    return date.toLocaleDateString(I18n.getLang(), {
      weekday: "short",
      day: "numeric",
      month: "long",
      year: "numeric"
    });
  }
};
