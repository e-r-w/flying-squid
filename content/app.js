
import React from 'react'
import { render } from 'react-dom'
import ReactCSSTransitionGroup from 'react-addons-css-transition-group'
import { hashHistory, Router, Route, IndexRoute, Link } from 'react-router'
import Home from './pages/home';
import Nav from './pages/nav';
import Stack from './pages/Stack';

class App extends React.Component {
  render() {
    return (
      <div>

        <Nav />

        <ReactCSSTransitionGroup
          component="div"
          transitionName="page"
          transitionEnterTimeout={500}
          transitionLeaveTimeout={500}
        >
          {React.cloneElement(this.props.children, {
            key: this.props.location.pathname
          })}
        </ReactCSSTransitionGroup>

      </div>
    )
  }
}

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={Home}/>
      <Route path="stack/:stackId" component={Stack} />
    </Route>
  </Router>
), document.getElementById('app'));