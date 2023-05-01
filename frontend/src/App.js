import './App.css';
import HomePage from "./HomePage";
import {BrowserRouter, Route} from "react-router-dom";
import Login from "./Login";

function App() {
  return (
      <BrowserRouter>
          <Route path="/" component={HomePage} />
          <Route path ="/login" component={Login} />
      </BrowserRouter>
  )
}

export default App;
