import { Container, Card, Form, Button, Stack } from "react-bootstrap";
import "./login.scss";
import { useState } from "react";

const Login = () => {
  const [isLogin, setIsLogin] = useState(true)
  return (
    <Container
      className="position-fixed p-0 top-0 bottom-0 start-0 end-0 bg-body-tertiary" fluid>

      <div className="login-container">
        <h3 className="mb-4 text-center">欢迎{isLogin? '登陆' : '注册'}ChatBox</h3>
        <Card className="shadow rounded-3 login-card">
          <Card.Body>
            {
              isLogin ? (<Form className="p-4">
                <Form.Group className="mb-4" controlId="formLoginEmail">
                  <Form.Label>邮箱</Form.Label>
                  <Form.Control type="email" placeholder="请输入邮箱地址" />
                </Form.Group>
                <Form.Group className="mb-4" controlId="formLoginPassword">
                  <Form.Label>密码</Form.Label>
                  <Form.Control type="password" placeholder="请输入密码" />
                </Form.Group>

                <div className="d-grid mt-4 pt-2">
                  <Button variant="primary" type="submit" className="p-2">
                    登陆
                  </Button>
                </div>
              </Form>) : (
                <Form className="p-4">
                  <Form.Group className="mb-4" controlId="formSignUpEmail">
                    <Form.Label>邮箱</Form.Label>
                    <Form.Control type="email" placeholder="请输入邮箱地址" />
                  </Form.Group>
                  <Form.Group className="mb-4" controlId="formSignUpPassword">
                    <Form.Label>密码</Form.Label>
                    <Form.Control type="password" placeholder="请输入密码" />
                  </Form.Group>

                  <Form.Group className="mb-4" controlId="formSignUpCode">
                    <Form.Label>验证码</Form.Label>
                    <Stack gap={2} direction="horizontal">
                      <Form.Control type="text" placeholder="请输入验证码" />
                      <Button className="verfyCode">获取验证码</Button>
                    </Stack>
                  </Form.Group>

                  <div className="d-grid mt-4 pt-2">
                    <Button variant="primary" type="submit" className="p-2">
                      注册
                    </Button>
                  </div>
                </Form>
              )
            }

            {
              isLogin ? (
                <div className="trigger-form m-2">没有账号密码？
                  <button type="button" className="btn btn-link p-0" onClick={() => setIsLogin(false)}>去注册</button></div>
              ) : (
                <div className="trigger-form m-2">已有账号密码？
                  <button type="button" className="btn btn-link p-0" onClick={() => setIsLogin(true)}>去登陆</button></div>
              )
            }

          </Card.Body>
        </Card>
      </div>

    </Container>
  )
}

export default Login