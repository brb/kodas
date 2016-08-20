// ==UserScript==
// @name widen-gh-issues
// @description Widens code sections at Github issues pages
// @version 0.0.1
//
// @match https://github.com/*/*/issues/*
// @match https://github.com/*/*/pull/*
// ==/UserScript==

WIDEN = "WIDEN";
UNWIDEN = "UNWIDEN";

panel = document.querySelector("div#partial-discussion-sidebar");
timeline = document.querySelector("div.discussion-timeline");

sep = document.createElement("span");
sep.className = "path-divider";
sep.innerText = "#";

link = document.createElement("a");
link.innerText = WIDEN;
link.onclick = function() {
  if (this.innerText == WIDEN) {
    this.innerText = UNWIDEN;
    panel.style.display = "none";
    timeline.style.width = "980px";
    code = document.getElementsByTagName("pre");
    for (i = 0; i < code.length; i++) {
      code[i].style.width = "97vw";
      code[i].style.marginLeft = "calc(-1 * ((100vw - 100%) / 2))";
    }
  } else {
    this.innerText = WIDEN;
    panel.style.display = "";
    timeline.style.width = "760px";
    code = document.getElementsByTagName("pre");
    for (i = 0; i < code.length; i++) {
      code[i].style.width = "";
      code[i].style.marginLeft = "";
    }
  }
};

title = document.querySelector("h1.public");
if (title == null) {
    title = document.querySelector("h1.private");
}
title.appendChild(sep);
title.appendChild(link);
