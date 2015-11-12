README
======

> The game ain't in me no more. None of it.

xmlcutty is a dead simple tool for carving out elements from *large* XML files and not much else.

Anecdata: It processes a 2G XML file with almost no memory in less than three minutes.

Usage
-----

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

    $ xmlcutty -h
    Usage of xmlcutty:
      -path string
            select path (default "/")
      -rename string
            rename wrapper element to this name
      -root string
            synthetic root element
      -v    show version

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

    $ xmlcutty -path /a/b fixtures/sample.xml
    <b>
            <c>
            </c>
        </b><b>
            <c>
            </c>
        </b>

Make xmllint a bit happier:

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

Rename wrapper element on the fly:

    $ xmlcutty -rename beee -path /a/b fixtures/sample.xml
    <beee>
            <c>
            </c>
        </beee><beee>
            <c>
            </c>
        </beee>

All options:

    $ xmlcutty -root hi -rename beee -path /a/b/c fixtures/sample.xml | xmllint --format -
    <?xml version="1.0"?>
    <hi>
      <beee>
            </beee>
      <beee>
            </beee>
    </hi>
