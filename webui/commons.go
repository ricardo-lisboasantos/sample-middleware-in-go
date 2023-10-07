package webui

var cssList = []string{
	"css/style.css",
}

var jsLibs = []string{
	"js/script.js",
}

var jsCustomList = []string{
	"js/custom.js",
}

func openHTML() string {
	return "<!DOCTYPE html><html><head><title>Example</title>" + importCSS() + importJSLibs() + "</head><body>"
}

func closeHTML() string {
	return importCustomJS() + "</body></html>"
}

func importCSS() string {
	var css string
	for _, v := range cssList {
		css += "<link rel=\"stylesheet\" href=\"" + v + "\">\n"
	}
	return css
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
