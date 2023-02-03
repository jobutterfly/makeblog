package new

import (
    "os"
    "fmt"
    "time"
    "bytes"
    "strings"
    "errors"
    "path/filepath"
    "html/template"

    "github.com/yuin/goldmark"
)

func fileContains(s string, fp string) (bool, error) {
	fullPath, err := getPath(fp)
	if err != nil {
	    return false, err
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
	    return false, err
	}

	contains := strings.Contains(string(content), s)

	return contains, nil
}

func getPath(p string) (string, error) {
	dirPath := "/repos/makeblog"
	hPath, err := os.UserHomeDir()
	if err != nil {
	    return "", err
	}
	path := filepath.Join(hPath, dirPath, p)
	return path, nil
}

func getLast(s string) string {
	sPath := strings.Split(s, "/")
	if len(sPath) < 2 {
	    return strings.ReplaceAll(sPath[0], ".md", "")
	} else {
	    return strings.ReplaceAll(sPath[len(sPath) - 1], ".md", "")
	}
}

func getName(s string) string {
	s = getLast(s)

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Split(s, " ")
	for i, w := range words {
	    fChar := w[:1]
	    rest := w[1:]
	    fChar = strings.ToUpper(fChar)
	    fString :=  []string {fChar, rest}
	    w = strings.Join(fString, "")
	    words[i] = w
	}

	s = strings.Join(words, " ")
	return s
}

func updateBlog(name string) error {
	path, err := getPath("/blog/blog.html")
	if err != nil {
	    return err
	}

	file, err := os.ReadFile(path)
	if err != nil {
	    return err
	}

	i := strings.Index(string(file), `<ul id="posts-list">`)
	if i < 0 {
	    return errors.New("Could not find posts list in blog.html")
	}

	fileName := getLast(name)
	finalName := getName(name)

	t := time.Now()
	y, m, d := t.Date()
	date := fmt.Sprintf("%d/%d/%d", d, m, y)

	newLi := fmt.Sprintf("\n<li><a href=\"%s\">%s</a> - %s</li>", "/" + fileName, finalName, date)

	filePart := strings.SplitAfter(string(file), `<ul id="posts-list">`)
	if len(filePart) < 2  || len(filePart) > 2 {
	    return errors.New("Could not find posts list in blog.html")
	}

	newFile := strings.Join(filePart, newLi)

	if err := os.WriteFile(path, []byte(newFile), 0666); err != nil {
	    return err
	}

	return nil
}

func updateControllers(name string) error {
	path, err:= getPath("/controllers/controllers.go")
	if err != nil {
	    return err
	}

	finalName := getLast(name)

	funcName := getName(name)
	funcName = strings.ReplaceAll(funcName, " ", "")

	output := fmt.Sprintf(`
func Serve%s(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/%s.html");
}`, funcName, finalName)

	file, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
	    return err
	}

	defer file.Close()

	if _, err := file.WriteString(output); err != nil {
	    return err
	}

	return nil
}

func updateMain(name string) error {
	path, err := getPath("/main.go")
	if err != nil {
	    return err
	}

	reqPath := getLast(name)
	funcName := getName(name)
	funcName = strings.ReplaceAll(funcName, " ", "")
	newLi := fmt.Sprintf(`
	http.HandleFunc("/%s", controllers.Serve%s);`,reqPath, funcName)

	file, err := os.ReadFile(path);
	if err != nil {
	    return err
	}

	filePart := strings.SplitAfter(string(file), `http.HandleFunc("/about", controllers.ServeAbout);`)
	if len(filePart) < 2  || len(filePart) > 2 {
	    return errors.New("Could not find handle funcs in main.go")
	}

	newFile := strings.Join(filePart, newLi)

	if err := os.WriteFile(path, []byte(newFile), 0666); err != nil {
	    return err
	}

	return nil
}

func New(input string) error {
	var buf bytes.Buffer
	var final string
	var wrt bytes.Buffer

	file, err := os.ReadFile(input);
	if err != nil {
	    return err
	}

	if err := goldmark.Convert(file, &buf); err != nil {
	    return err;
	}

	final = buf.String();

	lp, err:= getPath("/new/layout.html")
	if err != nil {
	    return err
	}

	tmpl := template.New("layout")
	tmpl, err = tmpl.ParseFiles(lp)
	if err != nil {
	    return err
	}

	if err := tmpl.Execute(&wrt, template.HTML(final)); err != nil {
	    return err
	}

	outName := getLast(input)
	out, err := getPath("/blog/" + outName + ".html");
	if err != nil {
	    return err
	}

	if err := os.WriteFile(out, []byte(wrt.String()), 0666); err != nil {
	    return err
	}

	funcName := getName(input)
	funcName = strings.ReplaceAll(funcName, " ", "")


	inBlog, err := fileContains(outName, "/blog/blog.html")
	if err != nil {
	    return err
	}

	inControllers, err := fileContains(funcName, "/controllers/controllers.go")
	if err != nil {
	    return err
	}

	inMain, err := fileContains(funcName, "/main.go")
	if err != nil {
	    return err
	}

	if !inBlog {
	    if err := updateBlog(input); err != nil {
		return err
	    }
	}

	if !inControllers {
	    if err := updateControllers(input); err != nil {
		return err
	    }
	}

	if !inMain {
	    if err := updateMain(input); err != nil {
		return err
	    }
	}

	return nil
}













