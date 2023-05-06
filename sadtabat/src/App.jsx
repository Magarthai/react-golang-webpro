import Menubra from '../component/Menubra'
import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div className="app">
        <BrowserRouter>
        <Menubra />
        </BrowserRouter>
      </div>
    </>
  )
}

export default App
