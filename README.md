<p align="center">
  <img src="https://raw.githubusercontent.com/ramantehlan/pulse/master/resources/icons/xxhdpi.png?token=AG5RGAHLG2IXVO66VVVZOQ26XEODA">
  <h1 align="center">Pulse</h1>
</p>
Check Pulse In Real Time

- Add go report
- Add release
- Add godoc reference


# Development

> This project is under heavy development

- ReactJS
- Golang
- Socket.io
- NextJS
- Snap Packaging

## Commands

`make dev` to setup dev environment
`make build` to build the project

I had to take so many design decisions, not sure if all of them are right, the point here is to see what we can do to achieve something

## Failed attempts

- use cython to build c file and use it in go | not working | Because cython writes alot of G
- use cythong generated.so file in go using rainycape/dl | not working | You can link .so file in c, and then use it in go, but that will be just too much for a simple program
- in so | Gettin `undefined symbol: PyUnicode_FromFormat` error

- use go-python and go-python3 to import python file and use it | working but slow and buggy | not a good solution
- Tried compile main.py and command, not working | not working

## Trying now

- Will try to complie to .whl and then use it as a different service.
- Now, using gRPC to communicate with python script and then using the

- Best solution is to figure out the gatt and do it in go itself, but for now just doing it like this.

## Install files

- Fetch the go build file
- Fetch the .whl file to install to pip
- Fetch the desktop entry files

## remove script

## To test with the front end

- In root, run `make start`, this will start the server.
- Make changes to client and run npm dev server `npm dev`, it will run on 7001, but will communicate with 7000.

> You might have to allow cross origin in your browser to test it, it will work fine in production as the server and template both will be on 7000 port

# Acknowledgment
- <div>Icons made by <a href="https://www.flaticon.com/authors/kiranshastry" title="Kiranshastry">Kiranshastry</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>

<div>Icons made by <a href="https://www.flaticon.com/authors/monkik" title="monkik">monkik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>
- https://github.com/golang-standards/project-layout
- https://github.com/markbates/pkger

- BloomRPC

## Reason of having a go as a middleware is to able to identify device using the manufacturing data and then fetch the appropriate client from the tools | Add client in any language

## Developers note

- Add requirement.txt file
- Setup grpc for go and python

This project is just to test out

- You can add support for more devices by writing them in the tools folder with the device name and script to dcrypt it

In tools
- Lib to add the lib to talk to the device
- client to talk to the grpc server
- deviceName and pulse in the end to launch

## Possilbe development piviots 

- You can use it for other devices too
- Add support for other kind of data points

If you ever add support for more devices: here are somepoints to take care of:
  - Identify the devices in the pulse cmd itself and then pick the tool for it
  - Add build instructions to the Makefile

- We need different protoc for go and python since the output is differnet 

- I can't run pulse with out sudo since it's designed for embidding, but I can't run miband pulse with sudo since it doesn't install files to sudo paht
Had to symlink the local site-packages to sudo site-packages | add the python path didn't wor

- It was stupid of me to write two programs to access one bluetooths device.


- I don't even know why I am writting this blog, actually this is not a blog, it's a journal for the project I created.

- Sometimes you won't be able to see if this will lang long or not, but you should ask yourself that.

- I am not an expert when I started this project.

I guess on a upper level, microservices work great!

# conclusion

I even tried to compile python and create so files to make it fast and import it, then I create a wheel and install it in system, but that didn't work either.

I think, there can be ways to do multilever compilations, but that will be really raise the questions of stability, and even if you are able to run something, it won't be stable
You will 

Before you jump into any language, make sure you look at the library support and define the goals before you pick a library

No problem is big, we just haven't found ways to break it. 

There are many if, or else in this project, but it's the concolusoins I got are. 

# Desclamer

This is a very experimental project created a college submission in a 2 weeks time, so manythings could be done better, but things are done to fit the time and the decisions are based on it. If you still have any suggestion, you can add them to the suggestions issue and Thank you for understanding.

# Troubleshoot

If you regenerate the grpc files for python, you might have to manually enter the import part to add src init

## Applications

- Hospitals
- Monitoring
- Heart analysis

## DELETE - TEMP-REF
- https://www.delftstack.com/howto/python/how-to-install-a-python-package-.whl-file/
- https://medium.com/swlh/distributing-python-packages-protected-with-cython-40fc29d84caf
- https://dzone.com/articles/executable-package-pip-install
- https://www.cs.swarthmore.edu/~newhall/unixhelp/howto_C_libraries.html
- https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf

# License

GPL V3
