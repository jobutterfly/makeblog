## mmblog

1/30/2023

mmblog stands for micro micro blog. It is a tool written in c to create and host your own blog. The source code is [here](https://github.com/jobutterfly/mmblog) The genesis of this idea came when I was browsing [nozx's](https://nozx.tech) website and I found [twtxt](https://github.com/buckket/twtxt). It seemed like a nice project and I decided to build something similar. I decided to use the C programming language because I wanted to see how the <sys/socket.h> library worked and how easy would it be to build a simple webserver.

mmblog can do three things: create, add, serve. When `mmblog create` is called, a new text file is create where the blog itself will be stored. When `mmblog add <some text here>` is called, it creates a new entry in the text file. Finally, when `mmblog serve` is called, it creates a webserver and it listens to calls according to the port set in the main.c file.

This was my first real project written in C and I'm happy with how it came out. It being my first project in C, I know that it is far from perfect but I will try to improve it as I keep learning. One of the things I learned about C while writing this project is that strings in C are really tricky. There is no string type in C so there are two ways of declaring a string in C, either:

```c
char *a_string;
```


or,

```c
char another_string[1024];
```


I preferred working with with character arrays rather than with character pointers. I felt that it was more intuitive and I didn't need to go through the hassle of having to allocate memory and free the memory when using it. The only problem is that is has a fixed size, but it is not that hard to work around that.

The other thing that I learned while making this project is the importance of the man pages. For many functions like fopen or strcat, by using the available man pages, I was able to grasp how the functions work and what they did with no hassle. I didn't have to open up a browser and search for an answer. This also encouraged me to use more the `go doc` tool for the go standard library.








