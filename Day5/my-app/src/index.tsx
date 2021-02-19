import React from 'react';
import ReactDOM from "react-dom";
import "./index.css";
import Message from "./components/message";
import Login from "./components/Login";
import Home from "./components/home";
import {
  BrowserRouter,
  Switch,
  Route
} from "react-router-dom";


ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Switch>
        <Route exact path="/api/register">
          <Login />
        </Route>
        <Route exact path="/">
          <Home />
        </Route>
      </Switch>
      {/* <Message
      messageId={1}
      imgProfileUrl="https://i.ibb.co/Npv3bk7/image-2021-02-18-164910.png"
      senderName="Pepe"
      messageText="Life sucks, is banana a bread, gonna kill a ship later"
      />
      <Message
      messageId={2}
      imgProfileUrl="https://p.kindpng.com/picc/s/407-4074666_create-your-own-peepo-peepo-transparent-hd-png.png"
      senderName="CowBoy Peepo"
      messageText="What a nice tea"
      />
      <Message
      messageId={3}
      imgProfileUrl="https://ih1.redbubble.net/image.839413284.0422/flat,750x,075,f-pad,750x1000,f8f8f8.u14.jpg"
      senderName="Happy Peepo"
      messageText="Life is great, is banana a bread, is bread a banana"
      /> */}
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById("root")
);