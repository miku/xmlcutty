README
======

> The game ain't in me no more. [None of it](https://www.youtube.com/watch?v=h7yf8Vp2KAI&feature=youtu.be&t=1m46s).

xmlcutty is a dead simple tool for carving out elements from *large* XML files and not much else.

Anecdata: Processes a 2G XML file with almost no memory in less than three minutes.

Usage
-----

```sh
$ cat fixtures/sample.xml
<a>
    <b>
        <c>
        </c>
    </b>
    <b>
        <c>
        </c>
    </b>
</a>
```

Options:

```sh
$ xmlcutty -h
Usage of xmlcutty:
  -delete
        delete wrapper element
  -path string
        select path (default "/")
  -rename string
        rename wrapper element to this name
  -root string
        synthetic root element
  -v    show version
```

It looks a bit like [XPath](https://en.wikipedia.org/wiki/XPath).

```sh
$ xmlcutty -path /a fixtures/sample.xml
<a>
    <b>
        <c>
        </c>
    </b>
    <b>
        <c>
        </c>
    </b>
</a>
```

But it can only interpret the target element, relative to the root. We not
even support extracting text. There are other tools for that.

```sh
$ xmlcutty -path /a/b fixtures/sample.xml
<b>
        <c>
        </c>
    </b><b>
        <c>
        </c>
    </b>
```

Make [xmllint](http://xmlsoft.org/xmllint.html) a bit happier, by adding a
synthetic root element:

```sh
$ xmlcutty -root hello -path /a/b fixtures/sample.xml | xmllint --format -
<?xml version="1.0"?>
<hello>
  <b>
    <c>
        </c>
  </b>
  <b>
    <c>
        </c>
  </b>
</hello>
```

Rename wrapper element on the fly:

```sh
$ xmlcutty -rename beee -path /a/b fixtures/sample.xml
<beee>
        <c>
        </c>
    </beee><beee>
        <c>
        </c>
    </beee>
```

All options:

```sh
$ xmlcutty -root hi -rename beee -path /a/b/c fixtures/sample.xml | xmllint --format -
<?xml version="1.0"?>
<hi>
  <beee>
        </beee>
  <beee>
        </beee>
</hi>
```

It is even possible to parse XML files without a root element:

```sh
$ head fixtures/oai.xml
<record>
<header>
 <identifier>oai:arXiv.org:0704.0004</identifier>
 <datestamp>2007-05-23</datestamp>
 <setSpec>math</setSpec>
</header>
<metadata>
 <oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/"... >
 <dc:title>A determinant of Stirling cycle numbers counts ...

$ cat fixtures/oai.xml | xmlcutty -root x -path /record/metadata/dc/identifier \
                       | xmllint --format -
<?xml version="1.0"?>
<x>
  <identifier>http://arxiv.org/abs/0704.0004</identifier>
  <identifier>http://arxiv.org/abs/0704.0010</identifier>
  <identifier>http://arxiv.org/abs/0704.0012</identifier>
</x>
```

Poor man's XPath `text()` extraction:

```sh
$ cat fixtures/oai.xml | ./xmlcutty -rename '\n' -path /record/metadata/dc/identifier \
                       | grep -v "^$"
http://arxiv.org/abs/0704.0004
http://arxiv.org/abs/0704.0010
http://arxiv.org/abs/0704.0012
```
