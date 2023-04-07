import React, { useEffect, useState } from 'react'
import logo from './logo.svg'
import './App.css'

let socket = new WebSocket("ws://localhost:4000/ws")
const list = [1]
const ChatList = (props) => {
  return list.map(item => {
    return (
      <div key={item} className='chat-item'>
        <div className='Avatar'>pic</div>
        <div className='profile'>
          <h3>{props.name}</h3>
          <p>{props.message}</p>
        </div>
        <div className='time'>10:00</div>
      </div>
    )
  })
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
  }

  return (
    <div className="App">
      
    </div>
  )
}

export default App


   {/* <ul>
        {
          messages.map((m, i) => (
            <li key={i}>{m}</li>
          ))
        }
      </ul> */}
 {/* <input type="text" name="message" onChange={typeMessage} />
        <input type="submit" value="Send" onClick={sendMessage}></input> */}