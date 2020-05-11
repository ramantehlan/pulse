<p align="center">
<img src="https://raw.githubusercontent.com/ramantehlan/pulse/master/resources/icons/logo.png?token=AG5RGAHB4UGKJMSKJ6TSAO26YEBQW" width="100">
</p>

<h1 align="center">Pulse</h1>
<h5 align="center"> Simple Heartbeat Monitor :heart: :bar_chart:</h5>

<p align="center">
 <a>
    <img src="https://goreportcard.com/badge/github.com/ramantehlan/pulse" align="center">
 </a>
 <a>
    <img src="https://img.shields.io/badge/godoc-reference-green" align="center">
 </a>
 <a>
    <img src="https://img.shields.io/badge/license-MIT-blue" align="center">
 </a>
  
</p>
  

**Pulse is a heartbeat monitor; it connects with your smart band and fetches your pulse in real-time to display it on a dashboard**. It currently supports MiBand 2 and 3, but support for more devices can be added.

>  This readme only contains the technical details about the project; the journal for this project can be found here with more information on design decisions, applications etc. 


# Index

- [About](#about)
- [Usage](#usage)
  - [Pre-Requisites](#pre-requisites)
  - [Development Environment](#development-environment)
  - [File Structure](#file-structure)
  - [Commands](#commands)
- [Contribution](#contribution)
- [FAQ](#faq)
- [Acknowledgment](#acknowledgment)
- [License](#license)

# About

There are 3 main services:

1. pulse
2. pulseExplore
3. mibandPulse

### pulse

This service act as a stub for other services and as a server for the users. **It serves the frontend which is compiled with its binary using [pkger](https://github.com/markbates/pkger)**. It starts with after the pulseExplore service and waits for 5 seconds before starting, to allow pulseExplore to fetch devices, and after that, it fetches the list of Bluetooth devices and stores them in a state to send it to frontend when requested. Once the user requests to connect their band, the request is sent to mibandPulse as that is the only device currently supported. Then mibandPulse connects to that device and **streams real-time heartbeats to pulse**.

> If you want to know why it is designed this way, you might want to check out this blog.


### pulseExplore

pulseExplore use [bettercap/gatt](https://github.com/bettercap/gatt)(Fork of paypal/gatt package), which is designed to be embedded and required Sudo access. It only runs for 8 seconds and scans all the devices in the area and stream them via gRPC server.

### mibandPulse
mibandPulse use [yogeshojha/MiBand3](https://github.com/yogeshojha/MiBand3) package to connect to MiBand and stream them to pulse using [gRPC](https://grpc.io/). It is converted into shared objects(so) file using [Cython](https://cython.org/) to optimise the speed. 


<p align="center">
<img src="./resources/pulse.png" width="800" >
</p>

# Usage
**Pulse can be deployed on a [Raspberry Pi](https://www.raspberrypi.org/), and some services under this project can also be embedded.** To run it on your system, you can use the [makefile](https://www.gnu.org/software/make/manual/html_node/Introduction.html). Assuming that you have already set up the development environment, you can use following commands to build and run this project.

```sh
$ make setup # Install dependencies.
$ make tools # Compile python to .so and install the command on the system.
$ make run # Install and run the project.
```

Here, `make tools` build and install the **mibandPulse** package in your pip, This is because the author wanted it to be easily exportable as a single `.whl` file and use it as a command. If you want to remove the package, you can directly remove it using `pip uninstall mibandPulse`, or you can use `make uninstall`.

> Pulse has not been tested on raspberry pi and any other system other than on Arch. However, it is expected to run on other platforms with minimum efforts.



## Pre-Requisities
This pre-requisites are not necessarily for running the project, but if you plan to use or contribute to this project or play with the source code, knowledge of following things is recommended.

- [Golang](https://golang.org/)
- [Cython](https://cython.org/)
- [ReactJS](https://reactjs.org/)/[NextJS](https://nextjs.org/)
- [Socket.io](https://socket.io/)
- [gRPC](https://grpc.io/)

## Development Environment

To develop or build this project, make sure you have the following environment setup:

- Install, and setup Go environment.
- Install python.
- Install NodeJS and yarn.
- Install Make.
- Clone this project in your workspace.

Once you have set up the above environment, we will use make to install dependencies. Go to the root of the project and run the following command; it installs all the node modules, Python and Go dependencies.

```sh
$ make setup
```

## File Structure

The file structure of this project follows the [conventional standard](https://github.com/golang-standards/project-layout), so it should be reasonably easy to understand. A description is given below:


 Folder/File Name | Description
------------------|------------
/api | Protocol Definition.
/cmd | Main applications for this project. Home for `pulse` and `pulseExplore` services.
/internal | Private application code, it doesn't exported by Go. See [release note](https://golang.org/doc/go1.4#internalpackages). Contains gRPC generated and other files.
/resources | Resources used for readme or public.
/tools | Supporting tools for this project. Holds `mibandPulse` service.
/web | Web application used by pulse.
/bin | Binary and build files.

## Commands


```sh
build                          Build pulse and pulseExplore command
clean                          Remove all the build files
dev                            Start the development environment to test
grpc                           Build files based on api proto files
help                           Display help screen
install                        Build and install pulse pulseExplore command
pkger                          Compile web files in a package using pkger
run                            Run the project
setup                          Setup dev environment
tools                          Compile and install the miband library
uninstall                      Uninstall the mibandPulse command
web                            Build web files
```

# Contribution

 Your contributions are always welcome and appreciated. Following are the things you can do to contribute to this project.

 1. Report a bug.
 2. Request a feature.
 3. Create a pull request.

**:sparkles: It takes time and efforts to think, design and develops open-source projects, so If you like this project, do star it so contributors can know you appreciate their efforts. :)**

 > If you are new to open-source, make sure to check read more about it [here](https://www.digitalocean.com/community/tutorial_series/an-introduction-to-open-source) and learn more about creating a pull request [here](https://www.digitalocean.com/community/tutorials/how-to-create-a-pull-request-on-github).

# FAQ

##### Q.) I want to use this project for a commercial purpose? 
Great! Yes, you can! :tada: However, it will be appreciated that you reach out to the author and let him know. It will be nice to see it being applied in the real world, and maybe the author can decide to get involved with your endeavours.

##### Q.) I want to know how and why this project was created?
You can refer to this blog.

##### Q.) Can I add support for more devices and how?
Yes! You will need to create/add a service to `/tools`  for it and in pulse identify the device using the [Manufacturer code](https://www.bluetooth.com/specifications/assigned-numbers/company-identifiers/) and send the request to the service for that device.

# Troubleshoot

If you are going to use or develop this project, the following are a few things which might help you fix some of the known troubles.

- When testing with `yarn dev` and pulse, you will need to enable CORS, as the dev frontend server run on 7001 and pulse command use 7000, the socketio connection won't be established otherwise.
- When testing mibuildPulse, you might have to remove the `milib.XXX` references manually from files and add it back when you are ready to build it or commit it. I know this is not how it should be, but I didn't get time to fix this.

# Acknowledgment

- Logo made by <a href="https://www.flaticon.com/authors/trinh-ho" title="Trinh Ho">Trinh Ho</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
- Diagram created using [draw.io](https://www.draw.io/)
- gRPC endpoints tested using [BloomRPC](https://github.com/uw-labs/bloomrpc)

# License

**MIT License**

Copyright (c) 2020 Raman Tehlan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


