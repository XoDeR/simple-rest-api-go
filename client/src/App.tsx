import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Books from './Books';
import Detail from './Detail';
import AddModal from './AddModal';
import EditModal from './EditModal';
import { Link } from 'react-router-dom';

const App: React.FC = () => (

  <Router>
    <div>
      <nav>
        <ul>
          <li>
            <Link to="/">Books</Link>
          </li>
          <li>
            <Link to="/detail">Book Detail</Link>
          </li>
          <li>
            <Link to="/add">Add Book</Link>
          </li>
          <li>
            <Link to="/edit">Edit Book</Link>
          </li>
        </ul>
      </nav>
      <Routes>
        <Route path="/" element={<Books />} />
        <Route path="/detail" element={<Detail />} />
        <Route path="/add" element={<AddModal />} />
        <Route path="/edit" element={<EditModal />} />
      </Routes>
    </div>
  </Router>
)

export default App
