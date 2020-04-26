import Head from "next/head"
import {Component} from 'react'
import io from "socket.io-client"
import styled from "styled-components"
import {Line} from 'react-chartjs-2'

import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faBluetoothB} from '@fortawesome/free-brands-svg-icons'
import {faMobile} from '@fortawesome/free-solid-svg-icons'

const Layout = styled.div`
  width:1200px;
  margin-left: calc( (100% - 1200px)/2 );
  min-height: calc(80vh);
  margin-top:calc(20vh);
  display: grid;
  grid-template-rows: auto 1fr;
  grid-template-columns: 250px auto;
  grid-row-gap: 0px;
  grid-column-gap: 0px;
  grid-template-areas:
      "side body"
      "side body"
      "side body";
  font-size: 14px;
`;

const Body = styled.div`
  grid-area: body;
  padding: 10px;
  overflow:auto;
`

const Side = styled.div`
 padding:5px;
 grid-area: side;
 color: #000;
`

const Heading = styled.div`
  text-align: center;
  font-weight: bold;
  color: rgb(26,143,227,0.4)
`

const UL = styled.div`
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

// We need empty labels list
function createLabels(num){
  var labels = []
  for (let i = 0; i < num; i++) {
    labels.push("")
  }
  return labels;
}

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      socket: io(":7000"),
      selectedDevice: "",
      devicesList: {},
      connected: false,
      pulse: [ 60, 120, 105, 92, 70, 80, 65, 85, 100, 114, 18,  120, 105, 92, 70, 80, 65, 85, 100, 114, 80]
    }
  }

  componentDidMount() {
    this.state.socket.on("devices_list", (data) => {
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

  render() {
    let data = {
      labels: createLabels(30),
      datasets: [
        {
          label: "Heartbeat",
          fill: false,
          backgroundColor: 'rgb(255,100,100)',
          borderColor: 'rgb(158, 42, 43)',
          lineTension: 0.2,
          borderWidth:0.8,
          pointRadius: 1.3,
          pointBorderWidth: 0.2,
          data: this.state.pulse
        }
      ]
    }
    let options = {
        animation: {
           duration: 600,
           easing: "easeOutSine",
         },
        title: {
          display: true,
          text: 'Heartbeats',
          fontStyle: "bold",
        },
        legend: {
          display: false,
        },
        layout: {
          padding: {
            top: 0,
            left: 20,
            right: 0
          }
        },
      scales: {
        xAxes: [
          {
            display: false,
            gridLines: {
              display: false
            }
          }
        ],
        yAxes: [
          {
            ticks: {
              beginAtZero: false,
              padding: 10,
              fontSize: 11,
              stepSize: 6,
              fontColor: 'rgba(100,100,100,0.6)'
            },
            gridLines: {
              display: false
            }
          }
        ]
      }
    }
    return (
<>
  < Head >
        <title>Pulse</title>
  </Head>
  <Layout>
    <Side>

      <Heading>
        Online Devices</Heading>
      <UL>
        <li>
          <FontAwesomeIcon icon={faBluetoothB}/>  MI Band
        </li>
        {
          Object.entries(this.state.devicesList).sort().map(([key, value]) => 
          <li onClick={() => this.handleClick(key)} key={key}>  
             <FontAwesomeIcon icon={faBluetoothB}/> {value.Name == ""? key: value.Name}
          </li>)
        }
      </UL>
    </Side>
    <Body>

      <Line data={data} width={150} options={options} height={60}/>

    </Body>
  </Layout>
</>)
  }
}

export default App
