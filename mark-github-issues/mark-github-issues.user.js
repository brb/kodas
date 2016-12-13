// ==UserScript==
// @name mark-github-issues
// @description Mark GH issues in the milestones view.
// @version 0.0.1
//
// @match https://github.com/*/*/milestone/*
// ==/UserScript==

issues = document.querySelector("ul.js-draggable-issues-container");

for (i = 0; i < issues.children.length; i++) {
  link = issues.children[i].querySelector("a");
  if (link.href.match("/issues/")) {
    issues.children[i].style.backgroundColor = "#fcfcd9";
  }
}
