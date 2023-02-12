## makeblog

3/2/2023

Makeblog is the tool I built to create this blog. Code can be found [here](https://github.com/jobutterlfy/makeblog). There are still some things that I'd like to imporve on but I think that it is in a position where I can already share the progress that I have made.

The reason why created I this tool is to make it easier to write new blog posts. I know there are other tools that can accomplish this same objective, like [Hugo](https://guhugo.io), but I wanted something simpler and I wanted to make it myself. It is easy to write new blog posts because I can write them in [Markdown](https://markdownguide.org) and the program then parses it to html. The parsing is done by an external library called [Goldmark](https://github.com/yuin/goldmark) (I wanted to make things myself but I didn't want to make my own parser, maybe something for the future).

There are two possible arguments for makeblog: new, serve. Calling `makeblog new <markdownfile>` creates a new blog post taking as input a markdown file. The html is then stored in the blog directory inside the project. Calling `makeblog serve` serves the blog in port 3000.

Most of the heavy lifting in this program is done by the New function. The first thing new does is it reads the input file and then parses it and writes it to a new html file. Before writing it to a new html file, the header and footer of my blog are put by making use of a template. After that, the blog.html file is updated. This file contains a list of all the blog posts, so when a new one is created, it must be put here too. The next thing it does is create a new controller for the new blog post that is then used in a new Handlefunc created in main.

This project has some similarities with my previous one, [mmblog](/mmblog). The main similarity is the need to open, read, and write files. Already having the experience of having done this in mmblog certainly made it easier to implement it here. All in all, the project is really straightforward. The only problem is that when a new blog post is created, a new binary needs to be built to implement the new endpoint for it. I tried to implement a feature where when a new blog post is created, the program is recompiled, but with no success. I will keep trying!
