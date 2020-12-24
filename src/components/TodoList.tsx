import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';

import TodoItem from './TodoItem';
import NewTodoItem from './NewTodoItem';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
  },
}));

type Todo = {
  id: number,
  name: string,
  isChecked: boolean,
}

type TodoListProps = {
}

export default function TodoList({}: TodoListProps){

  const classes = useStyles();

  const [ nextTodoId, setNextTodoId ] = React.useState<number>(1);
  const [ todoItems, setTodoItems ] = React.useState<Todo[]>([]);

  const handleToggle = (id: number) => () => {

    let newTodoItems = todoItems.map(item=>{
      if(item.id === id)
      {
        let newItem = {...item};
        newItem.isChecked = !item.isChecked;
        return newItem;
      }
      else
      {
        return item;
      }
    });

    setTodoItems(newTodoItems);
  };

  const handleDelete = (id: number) => () => {
    let newTodoItems = todoItems.filter(item=>item.id!==id);
    setTodoItems(newTodoItems);
  };

  const addNewTodo = (name: string) => {
    if(name !== "")
    {
      const newNextTodoId = nextTodoId + 1;
      let newTodoItems = [...todoItems];
      newTodoItems.push({ id: nextTodoId, name: name, isChecked: false});
      setTodoItems(newTodoItems);
      setNextTodoId(newNextTodoId);
    }
  };


  return(
    <List className={classes.root}>
      {todoItems.map((todoItem: Todo) => {
        return (
          <TodoItem 
            key={todoItem.id}
            todoId={todoItem.id} 
            isChecked={todoItem.isChecked} 
            name={todoItem.name} 
            handleToggle={handleToggle}
            handleDelete={handleDelete}
          />
        );
      })}
      <NewTodoItem
        key={nextTodoId}
        addNewTodo={addNewTodo}
      />
    </List>
  );
}