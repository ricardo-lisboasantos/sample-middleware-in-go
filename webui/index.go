package webui

// Example: We cloud define a struct here and use it in the template
// This template must be rendered by html/template package (The template is done in the routing functions)
// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Todos     []Todo
// }

func Index() string {
	body := `<h1>Hello World</h1>
	<p>My name is <strong>Ricardo</strong></p>
`
	return openHTML() + body + closeHTML()
}
