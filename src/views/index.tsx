import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter as Router, Route, Switch } from 'react-router-dom';

import ApolloProvider from './components/ApolloProvider/ApolloProvider';
import AuthProvider from './components/Auth/AuthProvider';

const App = () => (
  <AuthProvider>
    <ApolloProvider>
      <Router>
        // Remove Ghpr // Remove
        <Switch>{/* <Route exact path="/" component={Ghpr} /> */}</Switch>
      </Router>
    </ApolloProvider>
  </AuthProvider>
);

ReactDOM.render(<App />, document.getElementById('react-init'));
