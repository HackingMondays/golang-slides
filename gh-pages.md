Update online slides on GitHub
==============================

Checkout master:

~~~
git checkout master
~~~

Uncomment this line in `pandoc.option`:

~~~
--variable revealjs-url=https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.1.0
~~~

Generate slides:

~~~
mkdoc
git checkout gh-pages
mv golang-slides.html index.html
git add index.html
git commit -m "updated slides"
git push
git checkout master
~~~

Check on: http://hackingmondays.github.io/golang-slides/

Comment line in `pandoc.option`
