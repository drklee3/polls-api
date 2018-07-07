import * as React from "react";
import {getPolls, Poll, Status} from "../../polls";

interface State {
  readonly polls?: Poll[];
  readonly status: Status;
}

export default class extends React.Component {
  public readonly state: State = {
    polls: undefined,
    status: Status.Loading,
  };

  public componentDidMount() {
    getPolls()
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
    const {polls, status} = this.state;

    if (status === Status.Loading) {
      return <div>loading polls...</div>;
    }

    if (status === Status.Error || polls === undefined) {
      return <div>error loading polls :(</div>;
    }

    console.log(polls); // tslint:disable-line

    return (
      <div className="columns">
        <div className="column">
          {
            polls.map(poll => (
              <div className="card">
                <div className="card-content">
                  <p className="title is-5">
                    {poll.Title}
                  </p>
                  <p className="content">
                    <ul>    
                      {
                        Object.keys(poll.content.choices).map(choiceID => {
                          const choice = poll.content.choices[choiceID];
                          
                          return (
                            <li key={choiceID}>
                              {`${choice.name} - ${choice.count}`}
                            </li>
                          );
                        })
                      }
                    </ul>
                  </p>
                </div>
              </div>
            ))
          }
        </div>
      </div>
    );
  }
}