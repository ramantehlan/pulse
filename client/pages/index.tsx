import { Component } from 'react'
import io from "socket.io-client"

type props = {}
type state = {
  list: string
}

class App extends Component<props, state> {
  constructor(props: any){
    super(props)
    this.state = {
          list: ""
    }
  }

  componentDidMount() {
    let socket = io()
    socket.on('/', (data:any) => {
      this.setState({
        list: data.message
      })
    })
  }

  render(){
    return (
        <div>
              Welcome to page!
              {this.state.list}
        </div>)
  }
}

export default App
