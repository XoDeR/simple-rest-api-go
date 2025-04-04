import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";

const Books: React.FC = () => {
  const navigate = useNavigate();

  const handleAddBook = () => {
    navigate("/add"); // Navigate to the AddModal route
  };

  const [books, setBooks] = useState<any[]>([])

  useEffect(() => {
    getBooks();
  }, [])

  const getBooks = () => {
    fetch("http://localhost:4000/books", {
      method: "GET"
    }).then(res => res.json())
      .then(res => {
        setBooks(res)
      }).catch(error => {
        console.log(`${error.message}`)
      });
  }

  return (
    <div>
      {books.map((book) => (
        <div key={book.id}>
          <p>{book.title}</p>
          <p>{book.author}</p>
          <p>{book.desc}</p>
          <button>Detail</button>
        </div>
      ))}
      <button onClick={handleAddBook}>Add</button>
    </div>
  )
}

export default Books;