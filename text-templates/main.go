package main

import (
	"os"
	"text/template"
)

func main() {
	// Create a new template

	t1 := template.New("t1")

	// Parse some content and add it to the template
	// Templates are a mix of static text and actions enclosed in {{}}
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// We can use template.must to panic incase parse returns an error
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	// By executing a template we generate its text with specific values for actions
	// {{}} is replaced with the value passed to execute
	t1.Execute(os.Stdout, "Hello, World!")
	t1.Execute(os.Stdout, 42)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"Python",
	})

	// Helper function to create a template
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Create a new template using the helper function
	// We can pass a struct or a map to the template
	// The template will replace {{.Name}} with the value of Name
	t2 := Create("t2", "Name: {{.Name}}\n")

	// Execute the template with a struct
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Execute the template with a map
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// We can use if-else statements in templates
	// A value is considered false if it is the default value of a type
	// For example, 0 for integers, false for booleans, "" for strings
	// We can use the - to trim the whitespace
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// We can use range to iterate over a slice or a map
	t4 := Create("t4",
		"{{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "Python"})
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})

}
