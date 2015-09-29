# golang-slides [![Build Status](https://travis-ci.org/HackingMondays/golang-slides.svg)](https://travis-ci.org/HackingMondays/golang-slides)

Presentation slides for [GO tutorial](http://hackingmondays.github.io/golang-slides/) based on
[monday-slides](https://github.com/HackingMondays/monday-slides).

### Dependencies

* [Pandoc](https://github.com/jgm/pandoc/releases) 1.15.0.6
* [reveal.js](https://github.com/hakimel/reveal.js) 3.1.0
* [GO](https://golang.org/dl/) 1.5.1
* Some extra tooling:

~~~
go get github.com/ffel/pandocfilter/cmd/includes
go get github.com/tischda/mkdoc
~~~

### Compile slides

To compile slides, run `mkdoc`

~~~
$ mkdoc
Running pandoc with options: [--from=markdown+yaml_metadata_block --template template/reveal.template
    --no-highlight --variable transition=slide --variable css=css/hacking.css --filter includes -t revealjs
    -o golang-slides.html 01-introduction.md 02-exemple.md 03-installation.md 04-langage.md
    05-exercices-1.md 06-exercices-2.md 07-conclusion.md metadata.yaml]
Total time: 87.202ms
~~~

### Exercise sources

The source code for the exercises is in the `src` directory. Add the project folder to your `GOPATH`:

~~~
export GOPATH=$HOME/src/gopath:$HOME/src/hacking/golang-slides
~~~

Test:

~~~
cd src/01-fail
go test
~~~

You have failing test, this is a good start ;-)


