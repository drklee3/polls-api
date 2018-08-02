import * as React from "react";
import NewPoll from "../components/NewPoll/NewPoll";

interface Props {
  readonly match?: any;
}

export default class extends React.Component<Props> {
  public render() {
    console.log(this.props); // tslint:disable-line
    return (
      <div className="section">
        <div className="container">
          <NewPoll match={this.props.match} />
        </div>
      </div>
    );
  }
}
