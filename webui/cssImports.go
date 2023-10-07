package webui

var cssList = []string{
	"css/style.css",
}

func importCSS() string {
	var css string
	for _, v := range cssList {
		css += "<link rel=\"stylesheet\" href=\"" + v + "\">\n"
	}
	return css
}
