import React from 'react';
import * as request from 'superagent';

export default class extends React.Component {

  constructor() {
    super();
    this.state = {loading: true};
  }

  componentDidMount() {
    this.setState({ loading: true });
    request
      .get(`/stacks/${this.props.routeParams.stackId}`)
      .end( (err, res) => {
        if(err) {
          this.setState({
            loading: false,
            error: err
          });
        }
        else {
          this.setState({
            stack: res.body,
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
        <h1>{this.state.stack.stackName}</h1>
        <StackResource stackId={this.state.stack.stackId} />
      </div>
    )
  }

}

class StackResource extends React.Component {

  constructor() {
    super();
    this.state = {loading: true};
  }

  componentDidMount() {
    this.setState({ loading: true });
    request
      .get(`/stacks/${this.props.stackId}/resources`)
      .end( (err, res) => {
        if(err) {
          this.setState({
            loading: false,
            error: err
          });
        }
        else {
          this.setState({
            resources: res.body,
            loading: false
          });
        }
      });
  }

  render() {
    if(this.state.loading || this.state.error){
      return <div></div>
    }
    return (
      <div>
        {this.state.resources.map( resource => (
          <div key={resource.id}>
              {resource.name}: {resource.type}
          </div>
        ))}
      </div>
    )
  }

}