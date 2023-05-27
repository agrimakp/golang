
const socket = new WebSocket("ws://localhost:4000/ws");

socket.addEventListener("open", (event) => {
  socket.send("Hello Server!");
});

// Listen for messages
socket.addEventListener("message", (event) => {
  console.log("Message from server ", event.data);
});


socket.send("aaaa")



<!-- Run server
 -->

 go run main.go