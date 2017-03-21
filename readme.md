Description
===========

Project Euler A Golang hacking mission

Website: [Go Project Euler Website](http://goprojecteuler.appspot.com/)

** Who's It For? **

* Developers of all experience levels

** Why Should you play this? **

1. To test your Go skills and solve programming puzzles
2. It's a fun, interactive way to experience Go through the context of a story
3. Compete with other Go developers to find the most creative solutions

** Ok, what else ? **

1. This project will test your skills
2. The questions are based on the Project Euler, a series of challenging mathematical/computer programming problems
that will require more than just mathematical insights to solve.
3. Have fun.


Requirements
============

To use this, you must have the following items installed and working:

* Go 1.4 or newer (Tested with versions: 1.4; 1.5; 1.6; 1.7)
* Go present tool
* Google App Engine Golang https://cloud.google.com/appengine/docs/go/download

Usage
=====

** Installing **

Download required [*Present](https://godoc.org/golang.org/x/tools/cmd/present), or install it via go get.
`go get golang.org/x/tools/cmd/present`
* There are no external dependencies, no VENDOR needed.

** Running **

There are two ways to run this project:

# 1. Using Golang Present directly:
`present -base=present -orighost=localhost -play=true`

Visit `http://127.0.0.1:3999` in your browser


# 2. Using go appengine:
`$ goapp serve -port=8081`
Visit `http://127.0.0.1:8081` in your browser


Enjoy hacking o/

APP ENGINE DEV
===============

* Run locally with appengine
`$ goapp serve -port=8081`

* Deploy to prod
`$ appcfg.py update -A goprojecteuler -V 1 /opt/gocode/src/github.com/projecteuler`

* Rollback deployment
`$ appcfg.py rollback -A goprojecteuler -V 1 /opt/gocode/src/github.com/projecteuler`


Levels:
Right now, there are 6 levels of this game done, But we are looking for other Gophers who are looking to contribute and add new levels.
Please feel free to submit PRs, they will be very appreciated! :)


Credits
=====
The app idea was inspired by the Operation Go: A Routine Mission game


License
=======

```
The MIT License (MIT)

Copyright (c) 2016 Thomas Modeneis

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```