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
