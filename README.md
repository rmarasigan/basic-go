# Basic GO Language

Please do install go x.x.x on your machine and configured. If you're not familiar with GO Language, please go over to their [page](https://golang.org/doc/install)

For Windows, you also need to install, GIT, [GCC Compiler](http://tdm-gcc.tdragon.net/download) and set the environment variables.


For Linux:

```
$ nano .profile
```

And add this inside the .profile

```
# set PATH so it includes user's private bin directories

PATH="$HOME/bin:$HOME/.local/bin:$PATH"
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Proper sequence:
* Hello World

* Constants
> Go supports constants of character, string, boolean, and numeric values.
> Constant is a variable in Go with a fixed value. Any attempt to change the value of a constant will cause a run-time panic. Constant must be declared with `const` keyword and an initial value.

* Functions
> A function is a block of code that takes some input(s), does some processing on the input(s) and produces some output(s).

* Encapsulation

* Startup

* HTTP
> A basic HTTP server has a few key jobs to take care of.
  - Process dynamic requests: Process incoming requests from users who browse the website, log into their accounts or post images.
  - Serve static assets: Serve JavaScript, CSS and images to browsers to create a dynamic experience for the user.
  - Accept connections: The HTTP Server must listen on a specific port to be able to accept connections from the internet.

* Environment

* Handling Signals

* TCP/IP

* UDP

Please go to [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read#leanpub-auto-introducing-go) for a detailed discussion.


# Reference
[Tour of Go](https://tour.golang.org/welcome/1)
[Golang Functions](https://www.callicoder.com/golang-functions/)
[Go Web Examples](https://gowebexamples.com/http-server/)
