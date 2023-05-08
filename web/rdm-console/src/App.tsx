import './App.css'
import { Routes, Route, Outlet, Link } from "react-router-dom";
import Home from './Home';
import TopNav from './TopNav';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="*" element={<NoMatch />} />
      </Route>
    </Routes>
  )
}

const Layout = () => {
  return(
    <>
      <TopNav />
      <Outlet /> 
    </>
  )
}

function NoMatch() {
  return (
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/">Go to the home page</Link>
      </p>
    </div>
  );
}

export default App
