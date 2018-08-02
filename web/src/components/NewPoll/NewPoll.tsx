import * as React from "react";
import {createPoll, Poll, Status} from "../../polls";

interface State {
  readonly poll?: Poll;
  readonly status: Status;
}

interface Props {
  readonly match: any;
}

export default class extends React.Component<Props> {
  public readonly state: State = {
    poll: undefined,
    status: Status.Loading,
  };

  public handleSubmit(event: any) {
    event.preventDefault();

    createPoll(event)
      .then((data: Poll) => {
        this.setState({
          polls: data,
          status: Status.Success,
        });
      })
      .catch(err => {
        console.error("Error loading polls:", err); // tslint:disable-line no-console
        this.setState({
          status: Status.Error,
        });
      });
  }

  public render() {
    return (
      <form action="">
        <div className="field">
          <label className="label">Poll Title</label>
          <div className="control">
            <input className="input" type="text" placeholder="A Poll" />
          </div>
        </div>

        <div className="field">
          <div className="control">
            <label className="checkbox">
              <input type="checkbox" />
              Allow multiple selections
            </label>
          </div>
        </div>

        <div className="field">
          <label className="label">Option</label>
          <div className="control">
            <input className="input" type="text" placeholder="option" />
          </div>
        </div>

        <div className="field is-grouped">
          <div className="control">
            <button className="button is-link">Submit</button>
          </div>
          <div className="control">
            <button className="button is-text">Cancel</button>
          </div>
        </div>
      </form>
    );
  }
}