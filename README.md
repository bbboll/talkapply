# talkapply
talkapply is designed to be a simple tool that helps you dealing with a typical situation at university: 
Students form groups by themself to work on several projects they can choose from. 
You don't want 20 groups working on the same project, while others are left over. 
talkapply provides a minimal web app for signing up a group to a project. 
Once signed up, the project is not available for others anymore.

![Sample image of the login window](http://i.imgur.com/CDJ7vWp.png)

## How can I use it?
*I don't host any public talkapply server at the moment. Also I would not recommend use it in production at the moment.*

### Single file
I hated it to upload 175684359 lines of dependencies, way too many images and then maintain the database when I worked with PHP/Laravel. 
Because of that **you can just put a single executable that includes everything it needs** on your server. It won't download anything.

### Storage driver
You don't need any database system installed on your system. talkapply holds all data in the memory and will create several files for a permanent storage.

### Building
``` 
$ resources/build.py 
$ go build
```

*build.py* creates a golang file resources.go containing all resources as hex strings to be embed in the binary.

Dependencies:
* Go >= 1.7.1
* Python >= 3.4
  * css-js-html-minify (https://github.com/juancarlospaco/css-html-js-minify)
