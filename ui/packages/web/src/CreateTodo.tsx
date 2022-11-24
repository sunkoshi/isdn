import React, { useState } from "react";
import { trpc } from "./trpc";

function CreateTodo() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const _createTodo = trpc.createTodo.useMutation();
  async function createTodo() {
    _createTodo.mutate({ title, description });
    console.log(_createTodo.isError);
  }
  return (
    <div>
      <h2>Create Todo</h2>
      <input
        type={"text"}
        value={title}
        onChange={(e) => {
          setTitle(e.target.value);
        }}
      />
      <input
        type={"text"}
        value={description}
        onChange={(e) => {
          setDescription(e.target.value);
        }}
      />
      <button onClick={createTodo}>Create</button>
      {_createTodo.error && (
        <p>Something went wrong! {_createTodo.error.message}</p>
      )}
    </div>
  );
}

export default CreateTodo;
