import Chatbox from '@/components/chat/Chatbox'
import Contact from '@/components/chat/Contact'
import Sidebar from '@/components/chat/Sidebar'
import "./chat.scss"
const Chat = () => {
  let arr = []
  for (let i = 0; i < 100; i++) {
    arr.push(i)
  }
  return (
    <div className="chat-page">
      <aside className="aside-sidebar">
        <Sidebar />
      </aside>
      <nav className="nav-contact">
        <Contact />
      </nav>
      <article className="article-chatbox">
        <Chatbox />
      </article>
    </div>
  )
}

export default Chat