import * as React from "react";
import {getPoll, Poll, Status} from "../../polls";

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

  public componentDidMount() {
    const id = parseInt(this.props.match.params.pollId, 10);

    getPoll(id)
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
    const {poll, status} = this.state;

    if (status === Status.Loading) {
      return <div>loading poll...</div>;
    }

    if (status === Status.Error || poll === undefined) {
      return <div>error loading poll :(</div>;
    }

    return (
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
    );
  }
}