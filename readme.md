Description
===========

Project Euler A Golang hacking mission

 [Go Project Euler Website](http://goprojecteuler.appspot.com/)

** Who's It For? **
Developers of all experience levels

** Why Should you play this? **

1. To test their Go skills and solve programming puzzles
2. It's a fun, interactive way to experience Go through the context of a story
3. Compete with other Go developers to find the most creative solutions

1. This project will test your skills
2. The questions are based on the Project Euler, a series of challenging mathematical/computer programming problems
that will require more than just mathematical insights to solve.
3. Have fun.


Requirements
============

To use this, you must have the following items installed and working:

* Go 1.4
* Go present tool

Usage
=====

** Install **

`git clone https://github.com/gophergala/operation-go.git`

Download Go's present tool:

`go get code.google.com/p/go.tools/cmd/present`


** Run **

cd into project directory

`$ present -base=. -orighost=localhost -play=true`

Visit `http://127.0.0.1:3999` in your browser

Enjoy hacking o/

APP ENGINE DEV
===============

* Run locally with appengine
`$ goapp serve -port=8081`

* Deploy to prod
`$ appcfg.py update -A goprojecteuler -V 1 /opt/gocode/src/github.com/gibraltargolang/projecteuler`

* Rollback deployment
`$ appcfg.py rollback -A goprojecteuler -V 1 /opt/gocode/src/github.com/gibraltargolang/projecteuler`


Credits
=====
The app idea was inspired by the Operation Go: A Routine Mission game


License
=======
The MIT License (MIT) Copyright (c) 2015 Thomas Modeneis