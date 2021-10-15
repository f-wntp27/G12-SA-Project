
import Navbar from "./components/Navbar";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import MecRec from "./components/MecRec";
import Home from "./components/Home";

function App() {
  return (
    <Router>
    <div>
    <Navbar/>
        <Route exact path="/" component={Home} />
        <Route exact path="/MecdicalRecord" component={MecRec} />
    </div>
    </Router>
  );
}

export default App;
