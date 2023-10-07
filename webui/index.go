package webui

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Index() string {
	body := `<h1>Hello World</h1>
	<p>My name is <strong>Ricardo</strong></p>
`
	return openHTML() + body + closeHTML()
}
