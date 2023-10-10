import React, {useState} from 'react'
import axios from 'axios'

export default function Register(props) {
    function register() {
        return () => {
          axios.post('http://localhost:8080/register', {
            username: document.getElementById("username").value,
            password: document.getElementById("password").value
          }).then((res) => {
            props.setCode(res.data)
            document.getElementById("success").innerHTML = "Registration successful. <br> Return to Login"
          }).catch(() => {
            document.getElementById("error").innerHTML = "User already exists"
          })
    
        }
      }
  return (
    <div className='login-container'>
        <input type='text' placeholder='username' id="username"></input>
        <input type='password' placeholder='password' id="password"></input>
        <button onClick={register()}>Register</button>
        <a href="#" onClick={props.toggleMode}>Login</a>
        <p id="error" style={{color: "red"}}></p>
        <p id="success" style={{color: "green"}}></p>
    </div>
  )
}
