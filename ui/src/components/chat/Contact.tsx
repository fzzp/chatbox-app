import React, { useEffect, useState } from 'react'
import "./contact.scss"
import DefaultAvatar from "@/assets/imgs/default-avatar.png";
import { Image } from "react-bootstrap"
import { IoIosAdd } from "react-icons/io";
import { IoSearchOutline } from "react-icons/io5";
import AddModal from './AddModal';

const Contact = () => {
  const [contacts, setContacts] = useState<number[]>([])
  const [showModal, setShowModal] = useState(false)

  useEffect(() => {
    let arr = []
    for (let i = 0; i < 100; i++) {
      arr.push(i + 1)
    }
    setContacts(arr)
  }, [])

  return (
    <div className="com-contact">
      <div className="search-box d-flex">
        <div className="search-input">
          <input type="text" placeholder="搜索" />
          <IoSearchOutline className="search-icon" />
        </div>
        <div className="add-contact">
          <IoIosAdd />
        </div>
      </div>
      <div className="contact-box">
        {
          contacts.map(x => (
            <div className="d-flex p-2 contact-item" key={x}>
              <Image src={DefaultAvatar} width={"40px"} height={"40px"} roundedCircle />
              <div className="ps-1 text">
                <div className="d-flex justify-content-between">
                  <span className="ft14 text-truncate username">用户明用户明用户明用户明用户明</span>
                  <span className="ft12">2024/8/17</span>
                </div>
                <span className="ft12 last-message text-truncate">最后一条消息最后一条消息最后一条消息最后一条消息</span>
              </div>
            </div>
          ))
        }
      </div>
      <AddModal showModel={showModal} setShowModal={setShowModal}/>
    </div>
  )
}

export default Contact