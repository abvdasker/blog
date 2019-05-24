'use strict';

window.onload = function init() {
  var currentState = 0;

  var ARTICLE_PATH_RGX = /^\/articles\/(.*)$/i;
  var articlesByURLSlug;

  function addArticles(articleListElem, articles) {
    articles.forEach(function(article) {
      var articleElem = createArticleListItem(article);
      articleListElem.appendChild(articleElem);
    });
  }

  function createArticleListItem(article) {
    var articleElem = document.createElement("div");
    articleElem.classList.add("article");

    var listItemContentElem = createListItemContentElem(article);
    articleElem.appendChild(listItemContentElem);

    var bottomBorder = document.createElement("hr");
    articleElem.appendChild(bottomBorder);

    var elemId = articleElemId(article.base.id);
    articleElem.id = elemId;

    return articleElem;
  }

  function createListItemContentElem(article) {
    var contentElem = document.createElement("div");
    contentElem.classList.add("left-pad-md");
    
    var titleElem = document.createElement("h2");
    var titleTextNode = document.createTextNode(article.base.title);
    titleElem.appendChild(titleTextNode);

    var titleLink = document.createElement("a");
    titleLink.href = getArticleURL(article);
    titleLink.addEventListener("click", onArticleClickFn(article));

    titleLink.appendChild(titleElem);

    contentElem.appendChild(titleLink);

    var dateElem = document.createElement("span");
    var createdAtFormatted = reformatDateString(article.base.createdAt);
    var createdAtTextNode = document.createTextNode(createdAtFormatted);
    dateElem.appendChild(createdAtTextNode);
    dateElem.classList.add("right");
    contentElem.appendChild(dateElem);

    return contentElem;
  }

  function articleElemId(id) {
    return `article-${id}`;
  }

  function reformatDateString(dateString) {
    var date = new Date(dateString);
    return date.toLocaleDateString();
  }

  function onArticleClickFn(article) {
    return function(event) {
      renderArticle(article);
      var articleURL = getArticleURL(article);
      var state = {
        state: ++currentState,
        articleURLSlug: article.base.urlSlug,
      }
      history.pushState(state, article.base.title, articleURL);
      event.preventDefault();
      return false;
    }
  }

  function onArticleBack() {
    if (!articlesByURLSlug) {
      loadArticles(function(articles) {
        articlesByURLSlug = mapByURLSlug(articles);
        renderArticles(articles);
      });
    } else {
      var articles = articlesFromMappedByURLSlug();
      renderArticles(articles);
    }
  }

  function articlesFromMappedByURLSlug() {
    var articles = [];
    Object.keys(articlesByURLSlug).forEach(function(urlSlug) {
      var article = articlesByURLSlug[urlSlug];
      articles.push(article);
    });
    return articles;
  }

  function onArticleListForward() {
    if (!articlesByURLSlug) {
      loadArticles(function(articles) {
        articlesByURLSlug = mapByURLSlug(articles);
        var article = articlesByURLSlug[history.state.articleURLSlug];
        renderArticle(article);
      });
    } else {
      var article = articlesByURLSlug[history.state.articleURLSlug];
      renderArticle(article);
    }
  }

  function renderArticle(article) {
    var titleElem = document.createElement("h1");
    var titleText = document.createTextNode(article.base.title);
    titleElem.appendChild(titleText);

    var articleBody = document.createElement("div");
    articleBody.innerHTML = article.html;

    var dateElem = document.createElement("span");
    var dateText = document.createTextNode(article.base.createdAt);
    dateElem.appendChild(dateText);

    var articleContainer = DOM.byID("article-container");
    DOM.empty(articleContainer);

    articleContainer.appendChild(titleElem);
    articleContainer.appendChild(dateElem);
    articleContainer.appendChild(articleBody);

    showArticleContainer();
  }

  function mapByURLSlug(articlesJSON) {
    var mappedArticles = {};
    articlesJSON.forEach(function(article) {
      mappedArticles[article.base.urlSlug] = article;
    });
    return mappedArticles;
  }

  function getArticleURL(article) {
    return `/articles/${article.base.urlSlug}`
  }

  function showArticleContainer(articleContainer) {
    articleContainer = articleContainer || DOM.byID("article-container")
    DOM.byID("article-list").classList.add("hidden");
    articleContainer.classList.remove("hidden");
  }

  function showArticleList(articleListElem) {
    articleListElem = articleListElem || DOM.byID("article-list")
    DOM.byID("article-container").classList.add("hidden");
    articleListElem.classList.remove("hidden");
  }

  function renderArticles(articles) {
    var articleListElem = DOM.byID("article-list");
    DOM.empty(articleListElem);
    addArticles(articleListElem, articles);
    showArticleList(articleListElem);
  }

  function loadArticles(onSuccess) {
    Net.getJSON("/api/articles", onSuccess, function(err) {
      console.error(err);
    })
  }

  function loadArticle(path) {
    var urlSlug = getURLSlug(path);
    if (!urlSlug) {
      console.error("missing URL slug");
      return;
    }
    var articlePath = `/api/articles/${urlSlug}`;
    Net.getJSON(articlePath, function(article) {
      renderArticle(article);
    }, function(err) {
      console.error(err);
    })
  }

  function isArticlePage(path) {
    return !!ARTICLE_PATH_RGX.exec(path)
  }

  function getURLSlug(path) {
    var matches = ARTICLE_PATH_RGX.exec(path);
    if (!matches || matches.length < 2) {
      return null;
    }
    return matches[1];
  }

  function onStateChanged() {
    console.log("state changed");

    var pageState = history.state;
    console.log(pageState);
    if (!pageState || pageState.state < currentState) {
      console.log("article back");
      onArticleBack();
    } else if (pageState.state > currentState)  {
      console.log("article forward");
      onArticleListForward();
    } else {
      console.log("same state");
    }
    currentState = (pageState || {state: 0}).state;
  }

  function load() {
    window.onpopstate = onStateChanged;
    var path = window.location.pathname;
    if (isArticlePage(path)) {
      return loadArticle(path);
    }

    loadArticles(function(articlesJSON) {
      articlesByURLSlug = mapByURLSlug(articlesJSON);
      renderArticles(articlesJSON);
    });
  }

  function runNow(cb) {
  }

  load();
}
