import React from "react";
import ReactDOM from "react-dom";
import { HashRouter as Router, Route, Switch } from "react-router-dom";

const App = () => {
  return (
    <React.Fragment> // Remove
      Ghpr // Remove
      <Router>
        <Switch>
          {/* <Route exact path="/" component={Ghpr} /> */}
        </Switch>
      </Router>
    </React.Fragment> // Remove
  );
};

ReactDOM.render(<App />, document.getElementById('react-init'));
