const getVideo = async (t) => {
	let url = "http://localhost:7531/video",
		selection = encodeURIComponent(t.selectionText),
		res = await fetch(`${url}?q=${selection}&filters=long&longer=4000&count=200`),
		data = await res.json()
	openWin(data.response)
}

const openWin = (data) => {
	let links = ""
	for (let i of data.items) {
		links += `<div><a href="${i.player}">${i.title}</a></div>`
	}
	let html = `
      <html lang="en">
        <div>Count: ${data.count}</div>
        ${links}
      </html>`
	let win = window.open("", "_blank", "width=800,height=400")
	win.document.write(html)
}

chrome.contextMenus.create(
	{"title": "vk.com", "contexts": ["selection"], "onclick": getVideo}
)
