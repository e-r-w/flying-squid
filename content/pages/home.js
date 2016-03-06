import React from 'react';
import { Link } from 'react-router'
import * as request from 'superagent';


export default class extends React.Component {

  constructor() {
    super();
    this.state = {stacks: [], loading: true};
  }

  componentDidMount() {
    this.setState({ loading: true });
    request
      .get('/stacks')
      .end( (err, res) => {
        if(err) {
          this.setState({
            loading: false,
            error: err
          });
        }
        else {
          this.setState({
            stacks: res.body,
            loading: false
          });
        }
      });
  }

  render() {
    if(this.state.loading || this.state.error){
      return (
        <div className="page-container">
          <div className="container mt20">
            <h1>{this.state.loading ? 'loading' : JSON.stringify(this.state.error)}</h1>
          </div>
        </div>
      );
    }
    return (
      <div className="page-container">
        <div className="container mt20">
          {this.state.stacks.map( stack => (
            <div key={stack.stackId} className="col-sm-12 col-lg-4 mb-20">
              <Link to={`/stack/${stack.stackId}`} className={`stack-panel center stack-link ${this.panelType(stack.status)}`}>
                <div>
                  <span className="glyphicon glyphicon-cloud cloud-icon"></span>
                </div>
                <h2 className="stack-name">{stack.stackName}</h2>
                <h4>{stack.status}</h4>
                <h5>{stack.environment} - {stack.slice}</h5>
                <div>creation on: {stack.creationTime}</div>
                <div>for: {stack.lineOfBusiness}</div>
              </Link>
            </div>
          ))}
        </div>
      </div>
    )
  }

  panelType(status){
    switch (status) {
      case 'DELETE_FAILED':
      case 'UPDATE_FAILED':
      case 'CREATE_FAILED':
      case 'UPDATE_ROLLBACK_FAILED':
        return 'stack-panel--error';
      case 'CREATE_IN_PROGRESS':
      case 'UPDATE_IN_PROGRESS':
      case 'UPDATE_ROLLBACK_COMPLETE':
        return 'stack-panel--warning';
      default:
        return '';
    }
  }
}