const URL = "http://localhost:7531/video";

async function getVideo(text) {
  let s = encodeURIComponent(text.selectionText);
  let r = await fetch(`${URL}?q=${s}&filters=long&longer=4000&count=200`);
  let d = await r.json();
  openWin(d.response);
}

function openWin(data) {
  let win = window.open("", "_blank", "width=800,height=400");
  win.document.write(prepareHtml(data));
}

function prepareHtml(data) {
  let html = document.createElement('html');
  let count = document.createElement('div');
  count.innerHTML = `Count: ${data.count}`;
  html.appendChild(count);

  for (i of data.items) {
    let div = document.createElement('div');
    if (i.description != "") {
      div.innerHTML = `<abbr title="${i.description}"><a href="${i.player}">${i.title}▾<a></abbr>`
      html.appendChild(div);
    } else {
      div.innerHTML = `<a href="${i.player}">${i.title}</a>`
      html.appendChild(div);
    }
  }
  return html.outerHTML;
}

chrome.contextMenus.create(
  { "title": "Search video", "contexts": ["selection"], "onclick": getVideo }
)
