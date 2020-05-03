<p align="center">
  <img src="https://user-images.githubusercontent.com/29037312/79879805-cc488480-840c-11ea-8d79-737f9b22167d.png">
  <h1 align="center">Pulse</h1>
</p>
Check Pulse In Real Time

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


## Failed attempts

- use cython to build c file and use it in go | not working
- use cythong generated.so file in go using rainycape/dl | not working
- use go-python and go-python3 to import python file and use it | working but slow and buggy | not a good solution
- Tried compile main.py and command, not working | not working

## Trying now

- Will try to complie to .whl and then use it as a different service.
- Now, using gRPC to communicate with python script and then using the

- Best solution is to figure out the gatt and do it in go itself, but for now just doing it like this.

## To test with the front end

- In root, run `make start`, this will start the server.
- Make changes to client and run npm dev server `npm dev`, it will run on 7001, but will communicate with 7000.

> You might have to allow cross origin in your browser to test it, it will work fine in production as the server and template both will be on 7000 port

# Acknowledgment
- <div>Icons made by <a href="https://www.flaticon.com/authors/kiranshastry" title="Kiranshastry">Kiranshastry</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>

# License

GPL V3
