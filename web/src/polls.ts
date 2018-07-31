import axios from "axios";

const API_BASE_URL = "http://127.0.0.1:3001";

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

/**
 * Gets a list of all available polls
 *
 * @export
 * @returns {Promise.<Poll[]>} Array of polls
 */
export function getPolls() {
  return new Promise((resolve, reject) => {
    axios
      .get(`${API_BASE_URL}/polls`)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  });
}

/**
 * Gets a single poll
 *
 * @export
 * @param {number} id        Poll ID
 * @returns {Promise.<Poll>} Poll data
 */
export function getPoll(id: number) {
  return new Promise((resolve, reject) => {
    axios
      .get(`${API_BASE_URL}/polls/${id}`)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  });
}

/**
 * Submits poll choices
 *
 * @export
 * @param {number} id        Poll ID
 * @param {number[]} choices Poll choices to vote for
 * @returns {Promise.<Poll>} Poll data after submission
 */
export function submitPoll(id: number, choices: number[]) {
  return new Promise((resolve, reject) => {
    axios
      .post(`${API_BASE_URL}/polls/${id}/vote`, choices)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  }) 
}

/**
 * Creates a new poll
 *
 * @export
 * @param {Poll} poll        Poll data to create a poll
 * @returns {Promise.<Poll>} Poll data after creation
 */
export function createPoll(poll: Poll) {
  return new Promise((resolve, reject) => {
    axios
      .post(`${API_BASE_URL}/polls`, poll)
      .then(response => {
        resolve(response.data);
      })
      .catch(reject);
  })
}
