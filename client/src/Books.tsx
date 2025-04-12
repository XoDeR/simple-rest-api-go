import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import { DataTable } from './books-data-table/data-table';
import { BookEntry, columns } from './books-data-table/columns';

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

  const oldVersion = false;

  return (
    <>
      <div className="container mx-auto py-10">
        <DataTable columns={columns} data={books.map((book: any): BookEntry => ({
          id: book.id,
          title: book.title,
          author: book.author,
          desc: book.desc,
        }))} />
      </div>
      {oldVersion && (
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
      )}
    </>
  )
}

export default Books;