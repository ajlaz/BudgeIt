import { Button } from '@material-ui/core';
import './App.css';
import Form from './components/Form';
import { useState } from 'react';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import axios from 'axios';



function App() {
  const [code, setCode] = useState()

  
  return (
    <div className="App">
      {code ?
      <Dashboard userId={code}  /> : 
      <Login setCode={setCode}/>
      }
      
    </div>
  );
}

export default App;
