import * as React from "react";
import {getPolls, Poll} from "./polls";

interface State {
  readonly polls?: Poll[];
}

export default class extends React.Component {
  public readonly state: State = {
    polls: undefined,
  };

  public componentDidMount() {
    getPolls()
      .then((data: Poll) => {
        this.setState({polls: data});
      });
  }

  public render() {
    const {polls} = this.state;

    if (polls === undefined) {
      return <div>loading polls...</div>;
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