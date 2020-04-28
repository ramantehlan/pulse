import {Component} from 'react'
import styled from "styled-components"
import {Line} from 'react-chartjs-2'

// We need empty labels list
function createLabels(num){
  var labels = []
  for (let i = 0; i < num; i++) {
    labels.push("")
  }
  return labels;
}

export default class HeartBeat extends Component {
  constructor(props) {
    super(props)
    this.state = {
      pulse: [ 60, 120, 105, 92, 70, 80, 65, 85, 100, 114, 18,  120, 105, 92, 70, 80, 65, 85, 100, 114, 80]
    }
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
        <Line data={data} width={150} options={options} height={60}/>
      </>
    )
  }


}
