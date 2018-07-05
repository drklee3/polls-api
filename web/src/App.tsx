import "bulma/css/bulma.css";
import * as React from "react";
import Polls from "./components/Polls/Polls";

import NavBar from "./components/NavBar/NavBar";

class App extends React.Component {
  public render() {
    return (
      <div>
        <NavBar />
        <div className="section">
          <div className="container">
            <Polls />
          </div>
        </div>
      </div>
    );
  }
}

export default App;
