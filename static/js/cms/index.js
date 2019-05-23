'use strict';

var SUBMIT_LOGIN_PATH = "/api/users/login";
var SUBMIT_EDIT_PATH = "/api/articles";

var ARTICLES_PATH = "/api/articles";

var LOGIN_HTML_PATH = "/static/html/cms/login.html";
var EDIT_HTML_PATH = "/static/html/cms/edit.html";

// if user is already logged in, render cms
// else render login

window.onload = function init() {
  var container = document.getElementById("container");
  var articlesMap;

  function render() {
    if (Auth.loggedIn()) {
      renderCMS();
    } else {
      renderLogin();
    }
  }

  function renderLogin() {
    Net.get(LOGIN_HTML_PATH, {}, function(response) {
      container.innerHTML = response;
      addSubmitLoginListener();
    }, function (err) {
      console.error("failed to fetch login");
      console.error(err);
    })
  }

  function renderCMS() {
    renderCMSEdit();
  }

  function renderCMSEdit() {
    Net.get(EDIT_HTML_PATH, {}, function(editorHTML) {
      renderCMSHTML(editorHTML);
      renderArticleHistory();
    }, function(err) {
      console.error(err);
    });
  }

  function renderCMSHTML(html) {
    container.innerHTML = html;

    addLogoutListener();
    addSubmitUpdateListener();
    addSubmitCreateListener();
    addEditPreviewListener();
    addNewArticleListener();
  }

  function renderArticleHistory(onComplete) {
    Net.getJSON(ARTICLES_PATH, function(articles) {
      setArticlesMap(articles);
      renderArticlesSelector(articles);
      if (onComplete) {
        onComplete();
      }
    }, function(err) {
      console.error(err);
    });
  }

  function setArticlesMap(articles) {
    articlesMap = {};
    articles.forEach(function(article) {
      articlesMap[article.base.uuid] = article;
    });
  }

  function renderArticlesSelector(articles) {
    var articlesByDate = buildArticlesByDate(articles);
    var articleSelector = buildArticleSelector(articlesByDate);
    var articleMenu = document.getElementById("article-menu");
    articleMenu.appendChild(articleSelector);
  }

  function buildArticlesByDate(articles) {
    var yearMap = {};
    var yearList = [];
    articles.forEach(function(article, i) {
      var dateStr = article.base.createdAt;
      var date = new Date(dateStr);
      var year = date.getFullYear();
      var month = date.getMonth();

      if (!yearMap[year]) {
        yearMap[year] = {
          months: {
          },
          monthsList: [],
        };

        yearList.push({
          date: date,
          year: year,
        });
      }
      if (!yearMap[year].months[month]) {
        yearMap[year].months[month] = {
          articles: [],
        };
        yearMap[year].monthsList.push({
          month: month,
          date: date,
        });
      }
      yearMap[year].months[month].articles.push({
        article: article,
        date: date,
      });
    });

    // sort years descending
    yearList.sort(function(first, second) {
      return second.year - first.year;
    });

    Object.keys(yearMap).forEach(function(year) {
      var yearData = yearMap[year];
      // sort each year's months descending
      yearData.monthsList.sort(function(first, second) {
        return second.month - first.month;
      });

      // sort each month's articles descending
      Object.keys(yearData.months).forEach(function(month) {
        var monthData = yearData.months[month];
        monthData.articles.sort(function(article1, article2) {
          return article2.date - article1.date;
        });
      });
    });

    return {
      yearMap: yearMap,
      yearList: yearList,
    };
  }

  function buildArticleSelector(articlesByDate) {
    var yearListItems = articlesByDate.yearList.map(function(yearListData) {
      var year = yearListData.year; 

      var yearLink = createLink(year);

      var yearData = articlesByDate.yearMap[year];

      var monthsElements = yearData.monthsList.map(function(monthsListData) {
        var month = monthsListData.month;
        var monthData = articlesByDate.yearMap[year].months[month];

        var monthStr = monthString(monthsListData.date);
        var monthLink = createLink(monthStr);

        var articleLinks = monthData.articles.map(function(articleData) {
          var articleLink = createLink(articleData.article.base.title, onEditArticleClick);
          var articleUUID = articleData.article.base.uuid;
          articleLink.setAttribute("data-article-uuid", articleUUID);
          articleLink.id = getArticleLinkID(articleData.article);
          return buildLinkList(articleLink);
        });

        return buildLinkList(monthLink, articleLinks);
      });

      return buildLinkList(yearLink, monthsElements);
    });

    var listContainer = document.createElement("ul");
    yearListItems.forEach(function(listItem) {
      listContainer.appendChild(listItem);
    });
    return listContainer;
  }

  function onEditArticleClick(event) {
    var link = event.target;
    var uuid = link.getAttribute("data-article-uuid");
    editArticle(uuid);
  }

  function editArticle(uuid) {
    var article = articlesMap[uuid];
    var editElems = getEditArticleElems();
    editElems.title.value = article.base.title;
    editElems.html.value = article.html;
    editElems.container.setAttribute("data-current-article-uuid", uuid);
    resetPreview(editElems);
  }

  function createLink(text, onClick) {
    var linkElem = document.createElement("a");
    var textElem = document.createTextNode(text);
    linkElem.appendChild(textElem);
    if (onClick) {
      linkElem.href = "#";
      linkElem.addEventListener('click', onClick);
    }
    return linkElem;
  }

  function buildLinkList(headerLink, listItems) {
    var newListItem = document.createElement("li");
    newListItem.appendChild(headerLink);
    if (listItems) {
      var listElem = document.createElement("ul");
      listItems.forEach(function(listItem) {
        listElem.appendChild(listItem);
      });
      newListItem.appendChild(listElem);
    };
    return newListItem;
  }

  function monthString(date) {
    return date.toLocaleString('en-us', { month: 'long' });
  }

  function addLogoutListener() {
    var logoutButton = document.getElementById("logout");
    logoutButton.addEventListener("click", onLogout);
  }

  function addSubmitUpdateListener() {
    var submitButton = document.getElementById("submit-update");
    submitButton.addEventListener("click", onSubmitUpdate);
  }

  function addSubmitCreateListener() {
    var submitButton = document.getElementById("submit-create");
    submitButton.addEventListener("click", onSubmitCreate);
  }

  function addEditPreviewListener() {
    var editElems = getEditArticleElems();
    var previewArea = document.getElementById("preview");

    editElems.title.addEventListener("input", onEditChangedFn(editElems, previewArea));
    editElems.html.addEventListener("input", onEditChangedFn(editElems, previewArea));
  }

  function addNewArticleListener() {
    var newArticleBtn = document.getElementById("new-article");
    newArticleBtn.addEventListener("click", onNewArticle);
  }

  function onNewArticle(event) {
    var editElems = getEditArticleElems();
    editElems.container.removeAttribute("data-current-article-uuid");
    editElems.title.value = "";
    editElems.html.value = "";
    resetPreview(editElems);
  }

  function getEditArticleElems() {
    var editHTML = document.getElementById("edit-article-html");
    var editTitle = document.getElementById("edit-article-title");
    var editContainer = document.getElementById("edit");

    return {
      title: editTitle,
      html: editHTML,
      container: editContainer
    };
  }

  function resetPreview(editElems) {
    var previewArea = document.getElementById("preview");
    onEditChangedFn(editElems, previewArea)();
  }

  function onEditChangedFn(editElems, previewArea) {
    return function() {
      if (editElems.title.value === "" || editElems.html.value === "") {
        disableCreateButton();
      } else {
        enableCreateButton();
      }
      var article = buildPreviewArticle(editElems);
      previewArea.innerHTML = '';
      previewArea.appendChild(article.title);
      previewArea.appendChild(article.body);
    }
  }

  function disableCreateButton() {
    var submitCreateBtn = document.getElementById("submit-create");
    submitCreateBtn.disabled = true;
  }

  function enableCreateButton() {
    var submitCreateBtn = document.getElementById("submit-create");
    submitCreateBtn.disabled = false;
  }

  function buildPreviewArticle(editElems) {
    var title = document.createElement("h1");
    var titleText = document.createTextNode(editElems.title.value);
    title.appendChild(titleText);

    var articleBody = document.createElement("div");
    articleBody.innerHTML = editElems.html.value;

    return {
      title: title,
      body: articleBody,
    };
  }

  function addSubmitLoginListener() {
    var loginForm = document.getElementById("login");
    loginForm.addEventListener("submit", onSubmitLogin);
  }

  function onSubmitCreate(event) {
    var articleRequest = getArticleData();
    Net.postJSON(SUBMIT_EDIT_PATH, articleRequest, function(article) {
      console.log("ARTICLE RESPONSE");
      console.log(article);
      renderArticleHistory(function() {
        editArticle(article.base.uuid);
      });
    }, function(err) {
      console.error(err);
    });
    console.log("edit submitted");
  }

  function onSubmitUpdate(event) {
    var articleUUID = document.getElementById("edit").getAttribute("data-current-article-uuid");
    if (!articleUUID || articleUUID === "") {
      console.error("no current article uuid found");
    }
    var articleRequest = getArticleData();
    var path = getUpdateArticlePath(articleUUID);
    Net.putJSON(path, articleRequest, function(article) {
      articleUpdated(article);
    }, function(err) {
      console.error(err);
    });
  }

  function articleUpdated(article) {
    articlesMap[article.base.uuid] = article;
    var articleLinkID = getArticleLinkID(article);
    var articleLink = document.getElementById(articleLinkID);
    articleLink.innerHTML = article.base.title;
  }

  function getArticleLinkID(article) {
    return "article-link-" + article.base.uuid;
  }

  function onLogout(event) {
    Auth.logout();
    render();
  }

  function getArticleData() {
    var title = document.getElementById("edit-article-title").value;
    var html = document.getElementById("edit-article-html").value;
    return {
      title: title,
      html: html
    };
  }

  function onSubmitLogin(event) {
    event.preventDefault();
    var loginRequest = buildLoginRequest();
    Net.postJSON(SUBMIT_LOGIN_PATH, loginRequest, function(response) {
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

  function getUpdateArticlePath(articleUUID) {
    return "api/articles/" + articleUUID;
  }

  render();
}
