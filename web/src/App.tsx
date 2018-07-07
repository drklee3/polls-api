import "bulma/css/bulma.css";
import * as React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import NavBar from "./components/NavBar/NavBar";
import HomePage from "./pages/index";
import PollPage from "./pages/poll";

class App extends React.Component {
  public render() {
    return (
      <Router>
        <div>
          <NavBar />
          <Route exact={true} path="/" component={HomePage} />
          <Route path="/poll/:pollId" component={PollPage} />
        </div>

      </Router>
    );
  }
}

export default App;
