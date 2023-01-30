package new

import (
    "fmt"
    "os"
    "bytes"
    "path/filepath"
    "html/template"
    "github.com/yuin/goldmark"
)

func New(name string, output string) error {
    var buf bytes.Buffer
    var final string
    var wrt bytes.Buffer

    file, err := os.ReadFile(name);
    if err != nil {
	return err
    }

    if err := goldmark.Convert(file, &buf); err != nil {
	return err;
    }

    final = buf.String();

    lp := filepath.Join("new", "layout.html")
    tmpl := template.New("layout")
    tmpl, err = tmpl.ParseFiles(lp)
    if err != nil {
	return err
    }

    if err := tmpl.Execute(&wrt, template.HTML(final)); err != nil {
	return err
    }

    os.WriteFile(output, []byte(wrt.String()), 0666);

    fmt.Printf(wrt.String())

    return nil
}













