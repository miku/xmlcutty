README
======

> The game ain't in me no more. None of it.

xmlcutty is a dead simple tool for carving out elements from large XML files and not much else.

Processes a 2G XML file with almost no memory in less than three minutes.

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
