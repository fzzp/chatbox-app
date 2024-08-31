import React from 'react'
import { Image } from 'react-bootstrap'
import DefaultAvatar from "@/assets/imgs/default-avatar.png";
import { AiOutlineMenu, AiOutlineLogout, AiOutlineRight } from "react-icons/ai";
import { GoDotFill } from "react-icons/go";
import "./sidebar.scss"
const Sidebar = () => {
  return (
    <div className="com-siderbar">
      <div className="d-flex justify-content-center avatar">
        <Image src={DefaultAvatar} width={"40px"} height={"40px"} roundedCircle />
        <div className="dot online offline">
          <GoDotFill />
        </div>
      </div>
      <div className="d-flex justify-content-center position-relative settings">
        <AiOutlineMenu style={{fontSize: "20px"}} />
        <ul className="settings-box text-secondary">
          <li>个人资料 <AiOutlineRight /></li>
          <li>退出登陆 <AiOutlineLogout /> </li>
        </ul>
      </div>
    </div>
  )
}

export default Sidebar