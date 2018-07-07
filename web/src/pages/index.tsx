import * as React from "react";
import Polls from "../components/Polls/Polls";

export default class extends React.Component {
  public render() {
    return (
      <div className="section">
        <div className="container">
          <Polls />
        </div>
      </div>
    );
  }
}