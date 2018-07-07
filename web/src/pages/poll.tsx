import * as React from "react";
import Poll from "../components/Poll/Poll";

interface Props {
  readonly match?: any;
}

export default class extends React.Component<Props> {
  public render() {
    console.log(this.props); // tslint:disable-line
    return (
      <div className="section">
        <div className="container">
          <Poll match={this.props.match} />
        </div>
      </div>
    );
  }
}
