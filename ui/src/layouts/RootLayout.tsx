import React from 'react'
import { Link, Outlet } from "react-router-dom";
import { Button, Container, Nav, Navbar } from "react-bootstrap"

const RootLayout = () => {
  return (
    <>
      <Navbar className="bg-body-tertiary">
        <Container>
          <Navbar.Brand href="#home" className="me-4">ChatBox</Navbar.Brand>
          <Nav className="me-auto d-flex">
            <Nav.Link href="#home" className="me-4">首页</Nav.Link>
            <Nav.Link href="#link">关于</Nav.Link>
          </Nav>
          <Link className="ms-auto" to={`/login`}>
            <Button size={`sm`}>登陆</Button>
          </Link>
        </Container>
      </Navbar>
      <Outlet />
    </>
  )
}

export default RootLayout