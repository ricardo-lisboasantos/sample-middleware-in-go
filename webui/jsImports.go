package webui

var jsLibs = []string{
	"js/script.js",
}

var jsCustomList = []string{
	"js/custom.js",
}

func importJSLibs() string {
	var js string
	for _, v := range jsLibs {
		js += "<script src=\"" + v + "\"></script>\n"
	}
	return js
}

func importCustomJS() string {
	var js string
	for _, v := range jsCustomList {
		js += "<script src=\"" + v + "\"></script>\n"
	}
	return js
}
