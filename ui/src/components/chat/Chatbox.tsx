import { useState, useEffect } from 'react'
import "./chatbox.scss"
import { TfiMoreAlt } from "react-icons/tfi";
import { Image, OverlayTrigger,Tooltip,Button } from 'react-bootstrap';
import DefaultAvatar from "@/assets/imgs/default-avatar.png";
import { BsEmojiSmile } from "react-icons/bs";
import { HiOutlinePhotograph } from "react-icons/hi";
import { FaRegImage } from "react-icons/fa6";
import { TbPhoto } from "react-icons/tb";
const Chatbox = () => {
  const [contacts, setContacts] = useState<number[]>([])
  useEffect(() => {
    let arr = []
    for (let i = 0; i < 100; i++) {
      arr.push(i + 1)
    }
    setContacts(arr)
  }, [])
  return (
    <div className="com-chatbox">
      <div className="message-header text-truncate">
        <span>用户名</span>
        <TfiMoreAlt className="cp" />
      </div>

      <div className="message-box">
        {
          contacts.map(x => (
            <div className={`d-flex p-2 contact-item ${x % 2 == 0? 'keep-right' :''}`} key={x}>
              <Image src={DefaultAvatar} width={"40px"} height={"40px"} roundedCircle />
              <div className="ps-1 text">
                <div>
                  <span className="ft14 text-truncate ft12 username">用户名</span>
                  {/* <span className="ft12">2024/8/17</span> */}
                </div>
                <span className="ft12 last-message text-truncate">最后一条消息最后一条消息最后一条消息最后一条消息</span>
              </div>
            </div>
          ))
        }
      </div>

      <div className="message-input">
        <div className="toolbar">
          <BsEmojiSmile className="cp"></BsEmojiSmile>
          <TbPhoto className="photo-icon cp"></TbPhoto>
        </div>
        <div className="message-textarea">
          <textarea placeholder="请输入"></textarea>
        </div>
      </div>


    </div>
  )
}

export default Chatbox