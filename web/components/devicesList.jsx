import {Component} from 'react'
import styled from "styled-components"
import io from "socket.io-client"

import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faBluetoothB} from '@fortawesome/free-brands-svg-icons'
import {faMobile} from '@fortawesome/free-solid-svg-icons'

export const Heading = styled.div`
  text-align: center;
  font-weight: bold;
  color: rgb(26,143,227,0.4)
`


export const UL = styled.div`
  margin-top: 20px;
  text-transform: capitalize;
  color: rgba(57, 158, 90, 1);

  li {
    list-style:none;
    cursor:pointer;
    height:25px;
    line-height:25px;
    padding-left:10px;
    margin-top:5px;
    border-radius:3px;

    &:hover{
      background: rgba(57,158,90,1);
      color:rgba(255,255,255,1);
      }
  }
`

export default class DeviceList extends Component {
  constructor(props) {
    super(props)
    this.state = {
      socket: io(":7000"),
      listSocket: io(":7004"),
      selectedDevice: "",
      devicesList: {},
      connected: false,
    }
  }

  componentDidMount() {
   this.interval = setInterval(() => {
      this.state.listSocket.emit("get_devices", true)
      console.log("Sending devices list request")

      this.state.listSocket.on('disconnect', function () {
        console.log("socket disconnected");
      });   
   }, 3000)

   setTimeout( () => { 
     clearInterval(this.interval) 
     console.log("Stopping devices list fetch") 
   }, 16000)

   this.state.listSocket.on("devices_list", (data) => {
      data = JSON.parse(data)
      console.log(data)
      this.setState({devicesList: data})
    })
  }

  handleClick = (key) => {
    this.state.socket.emit("select_device", key, (data) => {
      console.log("selection data sent, data that got return is: " + data)
      this.setState({selectedDevice: data})
    })
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <>
      <Heading>
        Online Devices</Heading>
      <UL>
        {
          Object.entries(this.state.devicesList).sort().map(([key, value]) =>
          <li onClick={() => this.handleClick(key)} key={key}>
             <FontAwesomeIcon icon={faBluetoothB}/> {value.Name == ""? key: value.Name}
          </li>)
        }
      </UL>
      </>
    )
  }


}
