# Minifies all resources and creates a resources.go file containing
# all of them in binary format. You can access them in your go code
# like that: Resource('file_path') 

# TODO: Add SASS support

import os
# import sass
from css_html_js_minify import html_minify, js_minify, css_minify
import tempfile
import binascii
import html

def isJs(filename):
	return filename.endswith(".js")

def isCss(filename):
	return filename.endswith(".css")

def isSass(filename):
	return filename.endswith((".scss", ".sass"))

def isTpl(filename):
	return filename.endswith(".tpl.js")

def isHtml(filename):
	return filename.endswith(".html")

def isSvg(filename):
	return filename.endswith(".svg")

# convert a string to a hex representation (\xNN\xNN\xNN...)
def decodeString(s):
	return binascii.hexlify(s.encode("utf-8")).decode('utf-8')

# merci http://stackoverflow.com/a/19053800
def toCamelCase(snake_str):
    components = snake_str.split('_')
    # We capitalize the first letter of each component except the first one
    # with the 'title' method and join them together.
    return components[0] + "".join(x.title() for x in components[1:])

# def minifyJsFile(filename):


if __name__ == "__main__":

	# Dict of output files and the files to put in
	job = {
		
		"index.html": ["templates/index.html"],
		"scripts.js": ["scripts/jquery.js", "scripts/backbone.js"],
		"templates.tpl.js" : ["templates/overview.html"],
		"styles.css": ["styles/normalize.css", "styles/spinkit.css", "styles/milligram.css", "styles/custom.css"],
		"brand.svg" : ["images/brand.svg"]
	}

	# go code to produce
	goCode  = "// DO NOT CHANGE ANYTHING IN THIS FILE MANUALLY! See: resources/build.py\n"
	goCode += "package main\n\n"
	goCode += "import (\n    \"encoding/hex\"\n    \"errors\"\n)\n\n"
	goCode += "func ResourceFile(name string) ([]byte, error) {\n"

	# absolute path to the resources directory
	resDir = os.path.dirname(__file__)

	print("\nStarting resource building...\n")

	# create temporary directory
	# tmpDir = tempfile.mkdtemp(prefix='talkapply_resources_')
	#print("Creating temporary directory...\n     ", tmpDir)

	count = 0 # how many files already have been proceeded

	for outputFileName, contentFiles in job.items():

		goCode += "    if name == \"" + outputFileName + "\" {\n"
		output = "" # file content BEFORE converting to hex representation
		
		# generate html output
		if isHtml(outputFileName):
			print("Building html file '", outputFileName ,"'")

			for file in contentFiles:
				print("   ", file)
				fh = open(os.path.join(resDir, file), "r")

				if isHtml(file):
					output += html_minify(fh.read())
				else:
					print("  \033[91m\033[1mERROR: NO VALID FILE EXTENSION. SUPPORTED: .html\033[0m")

				fh.close()

		# generate template output (this one is a litte bit weird and needs a clean up... :| )
		# TODO: support more than one level in the file system
		elif isTpl(outputFileName):
			print("Building template js file '", outputFileName ,"'")
			output += "window._tpls = {"

			for file in contentFiles:
				print("   ", file)
				fh = open(os.path.join(resDir, file), "r")
				name = os.path.splitext(file)[0] # filename without extension to be used as object index

				if isHtml(file):
					output += toCamelCase(name) + ": "
					output += "\"" + html_minify(fh.read().replace('"','\\"')) + "\",\n"
				else:
					print("  \033[91m\033[1mERROR: NO VALID FILE EXTENSION. SUPPORTED: .html\033[0m")

				fh.close()

			output += "};"


		# generate javascript output
		elif isJs(outputFileName):
			print("Building javascript file '", outputFileName ,"'")

			for file in contentFiles:
				print("   ", file)
				fh = open(os.path.join(resDir, file), "r")

				if isJs(file):
					output += js_minify(fh.read())
				else:
					print("  \033[91m\033[1mERROR: NO VALID FILE EXTENSION. SUPPORTED: .js\033[0m")

				fh.close()

		# generate css output (Input can be css and sass)
		elif isCss(outputFileName):
			print("Building css file '", outputFileName ,"'")

			for file in contentFiles:
				print("   ", file)
				fh = open(os.path.join(resDir, file), "r")
				
				if isCss(file):
					output += css_minify(fh.read())

				#elif isSass(file):
				#	output += css_minify(sass.compile(string=fh.read()))

				else:
					print("  \033[91m\033[1mERROR: NO VALID FILE EXTENSION. SUPPORTED: .css\033[0m")
				
				fh.close()

		# generate css output (Input can be css and sass)
		elif isSvg(outputFileName):
			print("Building svg file '", outputFileName ,"'")

			for file in contentFiles:
				print("   ", file)
				fh = open(os.path.join(resDir, file), "r")
				
				if isSvg(file):
					# Minifying is a little bit buggy
					# output += html_minify(fh.read()) 
					output += fh.read()

				else:
					print("  \033[91m\033[1mERROR: NO VALID FILE EXTENSION. SUPPORTED: .svg\033[0m")
				
				fh.close()

		goCode += "        s,err := hex.DecodeString(\"" + decodeString(output) + "\")\n"
		goCode += "        if err != nil {\n" 
		goCode += "            return nil, err\n"
		goCode += "        }\n"
		goCode += "        return s,nil\n" 
		goCode += "    }\n"
		count  += 1


	# finally done
	goCode += "\n    return nil, errors.New(\"Asset could not be found.\")\n"
	goCode += "}"		

	# print(goCode)
	goFile = open("./resources.go", "w+")
	goFile.truncate()
	goFile.write(goCode)
	goFile.close()
