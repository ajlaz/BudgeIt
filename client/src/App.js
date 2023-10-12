import './App.css';
import React, { useState } from 'react';
import LoginPage from './components/LoginPage';
import Dashboard from './components/Dashboard';
import Header from './components/Header';




function App() {
  const [code, setCode] = useState(3)

  
  return (
    <div className="App">
      {code ?
      <div className='app-body'>
      <Header />
      <Dashboard userId={code}  />
      </div> : 
      <LoginPage setCode={setCode}/>
      }
      
    </div>
  );
}

export default App;
