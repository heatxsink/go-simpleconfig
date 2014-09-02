go-simpleconfig
===============

A simple config file abstraction. Why? I kept having to write something like this over and over again. Sadly, I'm tired of having N versions/variations in various projects.

Setup / Dependencies
--------------------

	$ go get github.com/heatxsink/go-simpleconfig

Some Details
------------

There's a function called Load(filename, object) and a Save(filename, object). These functions do precisely what they are named.