import React from 'react'
import { Outlet } from "react-router-dom";
import { Container } from 'react-bootstrap';

import "./chatLayout.scss"

const ChatLayout = () => {

  // TODO: 一些想法和未来实现:
  // 1. 支持更换封面 

  return (
    <Container fluid className='ChatRootLayout'>
      <div className="w1024 m-auto chatbox">
       <Outlet />
      </div>
    </Container>
  )
}

export default ChatLayout




























