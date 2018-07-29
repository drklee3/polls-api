import axios from "axios";

export enum Status {
  Loading,
  Success,
  Error,
}

interface PollChoice {
  id: number;
  name: string;
  color: string;
  count: number;
}

interface PollContent {
  choices: {
    [key: string]: PollChoice
  };
  options: PollOptions;
}

interface PollOptions {
  restrictions: string;
  poll_type: string;
  randomize_choices: boolean;
  endtime?: string;
}

export interface Poll {
  ID: number;
  UUID: string;
  CreatedAt: string;
  UpdatedAt: string;
  Title: string;
  Archived: boolean;
  content: PollContent;
}

export function getPolls() {
  return new Promise((resolve, reject) => {
    axios
      .get("http://127.0.0.1:3001/polls")
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  });
}

export function getPoll(id: number) {
  return new Promise((resolve, reject) => {
    axios
      .get(`http://127.0.0.1:3001/polls/${id}`)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  });
}

export function submitPoll(id: number, data: Poll) {
  return new Promise((resolve, reject) => {
    axios
      .post(`http://127.0.0.1:3001/polls/${id}/vote`, data)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  }) 
}
