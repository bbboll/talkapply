# talkapply
talkapply is a simple, self-contained tool to solve an organizational problem at university:
Projects need to be assigned to students in a first-come-first-serve fashion.
talkapply provides a minimal web app for students to pick a project. Once picked, the project will no longer be available.

![Sample image of the login window](http://i.imgur.com/CDJ7vWp.png)

## How can I use it?
*I don't recommend hosting any public talkapply servers at the moment.*

## Features
- **Single binary, no dependencies.**
    I always hated uploading 175684359 lines of dependencies and maintaining a database whenever I wanted to do anything with PHP/Laravel. 
    Instead, with talkapply **you can just put a single executable that includes everything it needs** on your server. It won't download anything.
- **No external storage driver**
    You don't need any database. talkapply caches temporary data in memory and will regularly flush it to files for permanent storage.

### Building
``` 
$ resources/build.py 
$ go build
```

*build.py* creates a golang file resources.go that contains all resources (markup, js,...) as hex strings. Building the project will embed them in the binary.

Dependencies:
* Go >= 1.7.1
* Python >= 3.4
  * css-js-html-minify (https://github.com/juancarlospaco/css-html-js-minify)
