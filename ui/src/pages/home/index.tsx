import React from 'react'
import { Stack, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import "./home.scss";
const Home = () => {

  const socket = new WebSocket("ws://localhost:4321/ws")
  socket.onopen = function () {
    console.log("WebSocket 链接成功～")
  }

  socket.onclose = function () {
    console.log("WebSocket 关闭了～")
  }

  socket.onerror = function (err) {
    console.log("WebSocket error: ", err)
  }

  socket.onmessage = function (msg) {
    console.log("->>> 接受到消息：", msg)
  }

  function handleSend () {
    socket.send(JSON.stringify({type: "send_message", payload: "你好 Hello"}))
  }

  return (
    <>
      <h1>Home</h1>
      <Stack direction="horizontal" gap={2}>
        <Link to={`/chat`}>Chat</Link>
        <Link to={`/login`}>Login</Link>
        <Link to={`/error`}>Error</Link>
        <Link to={`/404`}>404</Link>
      </Stack>

      <div>
        <Button variant="primary" onClick={handleSend}>
          发送消息
        </Button>
      </div>
    </>
  )
}

export default Home