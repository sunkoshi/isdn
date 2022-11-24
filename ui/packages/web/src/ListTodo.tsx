import { trpc } from "./trpc";

function ListTodo() {
  const todos = trpc.getTodos.useQuery({});
  const todoWithId = trpc.getTodoById.useQuery({ id: 2 });
  if (!todos.data) return <div>Loading...</div>;
  return (
    <div>
      <div>
        With ID
        <h1>{todoWithId.data?.title}</h1>
        <p>{todoWithId.data?.description}</p>
      </div>
      {todos.data.map((item, index) => {
        return (
          <div key={index}>
            <h1>{item.title}</h1>
            <p>{item.description}</p>
          </div>
        );
      })}
    </div>
  );
}

export default ListTodo;
