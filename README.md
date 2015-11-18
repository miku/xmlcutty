README
======

> The game ain't in me no more. [None of it](https://www.youtube.com/watch?v=h7yf8Vp2KAI&feature=youtu.be&t=1m46s).

xmlcutty is a simple tool for carving out elements from *large* XML files,
*fast*. Since it works in a streaming fashion, it uses almost no memory and
can process around 1G of XML per minute.

Why? [Background](http://stackoverflow.com/q/33653844/89391).

Install
-------

Use a deb or rpm [release](https://github.com/miku/xmlcutty/releases).

Or install with the go tool:

    $ go get github.com/miku/xmlcutty/cmd/xmlcutty

Usage
-----

```sh
$ cat fixtures/sample.xml
<a>
    <b>
        <c></c>
    </b>
    <b>
        <c></c>
    </b>
</a>
```

Options:

```sh
$ xmlcutty -h
Usage of xmlcutty:
  -path string
        select path (default "/")
  -rename string
        rename wrapper element to this name
  -root string
        synthetic root element
  -v    show version
```

It *looks* a bit like [XPath](https://en.wikipedia.org/wiki/XPath), but it really
is only a simple matcher.

```sh
$ xmlcutty -path /a fixtures/sample.xml
<a>
    <b>
        <c></c>
    </b>
    <b>
        <c></c>
    </b>
</a>
```

You specify a path, e.g. `/a/b` and all elements matching this path are printed:

```sh
$ xmlcutty -path /a/b fixtures/sample.xml
<b>
    <c></c>
</b>
<b>
    <c></c>
</b>
```

You can end up with an XML document without a root. To make tools like
[xmllint](http://xmlsoft.org/xmllint.html) happy, you can add a
synthetic root element on the fly:

```sh
$ xmlcutty -root hello -path /a/b fixtures/sample.xml | xmllint --format -
<?xml version="1.0"?>
<hello>
    <b>
        <c></c>
    </b>
    <b>
        <c></c>
    </b>
</hello>
```

Rename wrapper element - that is the last element of the matching path:

```sh
$ xmlcutty -rename beee -path /a/b fixtures/sample.xml
<beee>
    <c></c>
</beee>
<beee>
    <c></c>
</beee>
```

All options, synthetic root element and a renamed path element:

```sh
$ xmlcutty -root hi -rename ceee -path /a/b/c fixtures/sample.xml | xmllint --format -
<?xml version="1.0"?>
<hi>
    <ceee/>
    <ceee/>
</hi>
```

It will parse XML files without a root element just fine.

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
...
```

This is an example XML response from a web service. We can slice out the
identifier elements. Note any namespace - here `oai_dc` - is completely
ignored for the sake of simplicity:

```sh
$ cat fixtures/oai.xml | xmlcutty -root x -path /record/metadata/dc/identifier \
                       | xmllint --format -
<?xml version="1.0"?>
<x>
    <identifier>http://arxiv.org/abs/0704.0004</identifier>
    <identifier>http://arxiv.org/abs/0704.0010</identifier>
    <identifier>http://arxiv.org/abs/0704.0012</identifier>
</x>
```

We can go a bit further and extract the text element, which is like a poor man
`text()` in XPath terms. By using the a newline as argument to rename, we
effectively get rid of the enclosing XML tag:

```sh
$ cat fixtures/oai.xml | xmlcutty -rename '\n' -path /record/metadata/dc/identifier \
                       | grep -v "^$"
http://arxiv.org/abs/0704.0004
http://arxiv.org/abs/0704.0010
http://arxiv.org/abs/0704.0012
```

This last feature is nice to quickly extract text from large XML files.
