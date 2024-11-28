import {Route, Routes} from 'react-router-dom'
import './App.css'
import {Home} from "./pages/root/Home.tsx";
import {Login} from "./pages/auth/Login.tsx";

function App() {

  return (
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
      </Routes>
  )
}

export default App
