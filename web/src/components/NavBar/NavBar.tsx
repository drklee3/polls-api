import * as React from 'react';

export default class extends React.Component {
  public render() {
    return(
      <div className="uk-container">
        <nav className="uk-navbar-transparent" uk-navbar="true">
          <div className="uk-navbar-left">
            <ul className="uk-navbar-nav">
              <li className="uk-active"><a href="#">Home</a></li>
              <li>
                <a href="#">Parent</a>
                <div className="uk-navbar-dropdown">
                  <ul className="uk-nav uk-navbar-dropdown-nav">
                    <li className="uk-active"><a href="#">Active</a></li>
                    <li><a href="#">Item</a></li>
                    <li><a href="#">Item</a></li>
                  </ul>
                </div>
              </li>
              <li><a href="#">Item</a></li>
            </ul>
          </div>
        </nav>
      </div>
    )
  }
}