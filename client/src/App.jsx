import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { Fragment } from 'react';
import List from './component/List';
export default function App() {

  return (
    <Router>
      <div className="container">
        <div className="row">
          <div className="col mt-3">
            <h1>Todo App</h1>
          </div>
          <hr className="mg-3"></hr>
        </div>
        <List />
      </div>
    </Router>
  );
}