import React, {useState} from 'react'
import axios from 'axios'

export default function Login(props) {
    function login() {
        return () => {
          axios.post('http://localhost:8080/login', {
            username: document.getElementById("username").value,
            password: document.getElementById("password").value
          }).then((res) => {
            props.setCode(res.data)
          }).catch(() => {
            document.getElementById("error").innerHTML = "Incorrect username or password"
          })
    
        }
      }
  return (
    <div>
        <input type='text' placeholder='username' id="username"></input>
        <input type='password' placeholder='password' id="password"></input>
        <button onClick={login()}>Log in</button>
        <p id="error" style={{color: "red"}}></p>
    </div>
  )
}
