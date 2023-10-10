import React, {useState} from 'react'
import Login from './Login'
import Register from './Register'

export default function LoginPage(props) {
    const [mode, setMode] = useState(true);

    function toggleMode() {
        setMode(!mode)
    }
  return (
    <div className='login-page'>
    {mode ? <div className='login-container'><Login setCode={props.setCode} toggleMode={toggleMode}/></div>  : <div className='login-container'><Register setCode={props.setCode} toggleMode={toggleMode}/> </div>}
    </div>
  )
}
