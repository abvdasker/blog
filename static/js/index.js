'use strict';

window.onload = function init() {
    function addArticles(articleListElem, articles) {
        for (var i = 0; i < 5; i++) {
            articles.forEach(function(article) {
                var articleElem = createArticleListItem(article);
                articleListElem.appendChild(articleElem);
            });
        }
    }

    function createArticleListItem(article) {
        var articleBase = article.base;
        var articleElem = document.createElement("div");
        articleElem.classList.add("article");

        var listItemContentElem = createListItemContentElem(article.base);
        articleElem.appendChild(listItemContentElem);

        var bottomBorder = document.createElement("hr");
        articleElem.appendChild(bottomBorder);

        var elemId = articleElemId(articleBase.id);
        articleElem.id = elemId;

        return articleElem;
    }

    function createListItemContentElem(articleBase) {
        var contentElem = document.createElement("div");
        contentElem.classList.add("left-pad-md");
        
        var titleElem = document.createElement("h2");
        var titleTextNode = document.createTextNode(articleBase.title);
        titleElem.appendChild(titleTextNode);
        contentElem.appendChild(titleElem);

        var dateElem = document.createElement("span");
        var createdAtFormatted = reformatDateString(articleBase.createdAt);
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
    
    Net.getJSON("/api/articles", function(articles) {
        var articleListElem = DOM.byID("article-list");
        addArticles(articleListElem, articles);
    }, function(err) {
      console.error(err);
    })
}
