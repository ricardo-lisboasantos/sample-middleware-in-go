package webui

func openHTML() string {
	return "<!DOCTYPE html><html><head><title>Example</title>" + importCSS() + importJSLibs() + "</head><body>"
}

func closeHTML() string {
	return importCustomJS() + "</body></html>"
}
