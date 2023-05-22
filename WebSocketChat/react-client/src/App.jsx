import React, { useEffect, useState } from 'react'
import logo from './logo.svg'
import './App.css'

let socket = new WebSocket("ws://localhost:4000/ws")

const ChatItem = (props) => {
  return (
    <div className={`chat-item ${props.isActive ? 'active' : ''}`}>
      <div className='avatar'></div>
      <div className='profile'>
        <h3>{props.name}</h3>
        <p>{props.message}</p>
      </div>
      <div className='time'>{props.time}</div>
    </div>
  )
}

function App() {
  const [messages, setMessages] = useState(["Kunja", "Chippy"])
  const [draft, setDraft] = useState("")

  useEffect(() => {
    console.log("starting")
    socket.onmessage = (event) => {
      console.log("received from the server: ", event.data)
      setMessages([...messages, event.data])
    }
  })

  console.log(messages)

  const typeMessage = (e) => {
    e.preventDefault()
    setDraft(e.target.value)
  }

  const sendMessage = () => {
    socket.send(draft)
    setDraft("")
  }

  return (
    <div className="App">
      <div className='container'>

        {/* // sidebar */}
        <div className='sidebar'>
          <div className='avatar'>
            <h3><b>A</b></h3>
          </div>

          <img src='./icon.png' className='icon' />
          <img src='./icon.png' className='icon' />
          <img src='./icon.png' className='icon' />

        </div>

        {/* // messages list */}
        <div className='chatList'>
          <div className='header'>
            <h3>Messages</h3>
            <span className='count'>12</span>
            <div className='newChat'>+</div>
          </div>
          <input type='search' name='search' className='search' placeholder='Search messages' />


          <ChatItem isActive name='chinchu' message='1Lorem ipsum dolor sit amet' time='10.00' />
          <ChatItem name='manchu' message='2Lorem ipsum dolor sit amet' time='10.00' />

        </div>

        {/* chat */}
        <div className='chat'>

          {
            messages.map((m, i) => (
              <p key={i}>{m}</p>
            ))
          }

          <input type="text" name="message" value={draft} onChange={typeMessage} />
          <input type="submit" value="Send" onClick={sendMessage}></input>
        </div>










        <div className='chatDirectory'>ddfsdfds</div>
      </div>
    </div>
  )
}

export default App