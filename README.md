# makeblog

Welcome to makeblog! These are the possible commands:

```
    makeblog serve
```

This command serves the website on port 3000.

```
    makeblog new <input.md> 
```

This command is for creating new blog posts. It takes the input markdown file,
parses it into html, places it inside the layout template, and writes into the
output file. If the file does not exist, it creates a new one. The output is
placed into the blog folder.

## To-do

- Write html for new blog posts in new folder inside the blog folder called blog and maybe change the name of the blog folder to source or something else.
- Add the new folder to the path for displaying the files in the controllers and the links used in the blog.html file.
