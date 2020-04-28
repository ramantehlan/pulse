import Head from "next/head"
import {Component} from 'react'
import styled from "styled-components"

import DevicesList from "../components/devicesList.jsx"
import HeartBeat from "../components/HeartBeat.jsx"

export const Layout = styled.div`
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

export const Body = styled.div`
  grid-area: body;
  padding: 10px;
  overflow:auto;
`

export const Side = styled.div`
 padding:5px;
 grid-area: side;
 color: #000;
`


class App extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
<>
  < Head >
        <title>Pulse</title>
  </Head>
  <Layout>
    <Side>
          <DevicesList />
    </Side>
    <Body>
          <HeartBeat />
    </Body>
  </Layout>
</>)
  }
}

export default App
