# summon
Personalized man pages for things you want to remember

Using things like the [man pages](https://en.wikipedia.org/wiki/Man_page) or [tldr](https://tldr.sh) are great and very helpful, however, there are things I regularly find myself looking up that aren't on either the man pages or tldr. That's where `summon` comes in. It doesn't try to replace the man pages or tldr or even do what they do, but rather it's to call up pages containing things that are personalized and specfic for your usage.

## Installation
- Clone it
- build it (`go build`)
- `cp summon /usr/local/bin/`
- `summon init` - create directory to store pages

## Usage
```
summon [page/command]
```
### Commands
```
list - list all available pages
help - print this list and quit
version - print current version
init - create pages directory at `$HOME/.config/summon/`
```
### Calling a page
```
summon page-name
```

## Creating a page
Pages are stored in `$HOME/.config/summon`.

Create a file in that directory with the name you want to use to call that page.

Follow this template for writing the page otherwise the syntax highlighting might be messed up. It follows the same format as `tldr`.

```
note-title/header

- description of something to remember
    the command or line of code (start this line with a tab(\t))
```

**Example:**  
```
pypi

- build package
    python setup.py sdist bdist_wheel 

- test package
    python3 -m twine check dist/* 

- upload package to testpypi
    python3 -m twine upload --repository-url https://test.pypi.org/legacy/ dist/* 

- upload to pypi
    python3 -m twine upload dist/* 
```
Then that page can be called with `summon pypi`!  

Pypi is a good example because it's a page that's not on the man pages or `tldr` but is something I need to know how to do occasionally.

## Customizing syntax highlighting colors
There are default colors built in, however, they can be customized through the use of environment variables.  

Define any of these variables with a hex value.

`SUMMON_HEADER_COLOR` - Colors the page name and headers.  
`SUMMON_DESC_COLOR` - Colors the description line (line that starts with '-').  
`SUMMON_CMD_COLOR` - Colors the line containing the command (the tabbed line).  
`SUMMON_FLAG_COLOR` - colors anything in the line containing the command that starts wth '-' or '<'.  

`summon` usse the [termenv](https://github.com/muesli/termenv) for coloring.

## Use `summon` with `fzf`
These days I like to incorporate `fzf` into everything if I can.  

You can pipe the list of pages into `fzf` with: `summon list | fzf`  

To really make use of it, this script lets you call pages with `fzf`:
```sh
#!/bin/bash

d=$(summon list | fzf)

if [[ "$d" == '' ]]; then
  exit 0
else
  summon "$d"
fi
```


## Contributing
Feel free to make a pull request or open an issue if you find a bug.