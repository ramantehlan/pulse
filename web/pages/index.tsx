import { Component } from 'react'
import io from "socket.io-client"

type props = {}
type state = {
  socket: any,
  selectedDevice: string
}

class App extends Component<props, state> {
  constructor(props: any){
    super(props)
    this.state = {
          socket: io(":7000"),
          selectedDevice: ""
    }
  }

  componentDidMount() {
    this.state.socket.on("devices_list", (data:any) => {
      console.log("devices list: " + data)
    })
  }

  handleClick = (e:any) => {
     e.preventDefault()
     this.state.socket.emit("select_device", "12:E3:TC:UW", (data:any) => {
      console.log("selection data sent, data that got return is: " + data)
      this.setState({selectedDevice: data})
    })
  }

  render(){
    return (
        <div>
              Welcome to page! Test
              <button onClick={this.handleClick}> select device </button>
              {this.state.selectedDevice}
        </div>)
  }
}

export default App
