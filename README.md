# grfromepub

This is a very simple tool that automatically opens the GoodReads page of a book
when given an epub file, if a valid ISBN is found in its metadata. 

## Installation

You need a [working Go installation](https://golang.org/doc/install).

    $ go get -u github.com/barsanuphe/grfromepub
    $ go install ...grfromepub

## Usage

The script only requires one or a list of existing epub files as arguments.

    $ grfromepub /home/user/epubs/firstepub.epub ../epubs/secondone.epub

It will try to find an ISBN number inside the files, then will call your default
browser to open the corresponding goodreads page.

## Finding ISBNs

The script will do nothing if it cannot find an ISBN, even though it will 
do its best to do so. 
However, since there is no standardized way of including an ISBN in epub 
metadata, some publishers can be creative, or forget it include it altogether.

If *grfromepub* cannot find an ISBN in one of your files:
- open it with an editor like Sigil
- if you can locate the ISBN number somewhere in the OPF file, please make an
issue and copy the relevant xml parts.
- if you cannot find an ISBN number, either the publisher is lazy and/or your
epub is a conversion from another format.

Note that public domain works may not contain ISBN numbers (like, for example,
the Gutenberg Project epubs).
