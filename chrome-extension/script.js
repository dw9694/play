const HOME = "localhost", PORT = "7531", URL = `http://${HOME}:${PORT}`;

const getVideo = async (title) => {
    const text = encodeURIComponent(title.selectionText);
    const query = `/?query=${text}&filters=long&longer=4000&count=200`;
    const request = `${URL}${query}`;

    const response = await fetch(request);
    const data = await response.json();
    openWin(data.response);
};

const openWin = (data) => {
    const win = window.open("", "_blank", "width=800,height=400");
    const html = `
      <html lang="en">
        <div>Count: ${data.count}</div>
        <div>${get_links(data)}</div>
      </html>`;
    win.document.write(html);
};

const get_links = (data) => {
    let links = "";
    for (let i of data.items) {
        links += `<div><a href="${i.player}">${i.title}</a></div>`;
    }
    return links;
};

chrome.contextMenus.create({
    "title": "vk.com",
    "contexts": ["selection"],
    "onclick": getVideo
});
