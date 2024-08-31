# chatbox-app

基于Go+React+WebSocket的聊天应用程序栗子。


## 基础需求

- 单聊/群聊
- 加群看不到之前信息，移除群看不到之后信息
- 第一版仅支持文本信息


### ui 前端

```bash
cd ui
npm install 
npm run dev
```


### 数据设计

- 用户表（users）
- 联系人表 (contacts)
- 用户与联系人关系是多对多（user_contact）
- 消息表（messages）
  由谁发出，发给谁？单人/群聊；不管是单人还是群聊都可以理解为对话，
  可以理解为一对一多话，一对多对话，由此引申出对话表。
- 对话表（conversation）
  - 有对话就有对话成员
- 对话成员表（conversation_member）也可以理解为群。私聊群就只有两个人。


