import React from 'react';
import { Link } from 'react-router'

export default class extends React.Component {
  render() {
    return (
      <div>
        <Link to="/" ><h1>Home</h1></Link>
      </div>
    )
  }
}