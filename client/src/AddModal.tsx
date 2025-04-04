import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const AddModal: React.FC = () => {
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [desc, setDesc] = useState("");
  const navigate = useNavigate();

  const addBook = async () => {
    console.log(`Add Book: Title = ${title}, Author = ${author}, Description = ${desc}`);

    const book = JSON.stringify({ title, author, desc });

    fetch("http://localhost:4000/books", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: book,
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        navigate("/"); // Navigate back to the Books list after adding
      })
      .catch((error) => {
        console.error(`${error.message}!`);
      });
  };
  return (
    <div>
      <h1>Add Book</h1>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          addBook();
        }}
      >
        <div>
          <label>
            Title:
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </label>
        </div>
        <div>
          <label>
            Author:
            <input
              type="text"
              value={author}
              onChange={(e) => setAuthor(e.target.value)}
              required
            />
          </label>
        </div>
        <div>
          <label>
            Description:
            <textarea
              value={desc}
              onChange={(e) => setDesc(e.target.value)}
              required
            />
          </label>
        </div>
        <button type="submit">Add Book</button>
      </form>
    </div>
  );
};
export default AddModal;