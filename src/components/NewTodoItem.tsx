import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';

type NewTodoItemProps = {
  addNewTodo: (id: string) => void
}

export default function NewTodoItem({ addNewTodo }: NewTodoItemProps){

  const [ name, setName ] = React.useState<string>("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newName = e.target.value;
    setName(newName);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if(e.key === 'Enter')
    {
      addNewTodo(name);
    }
  }

  return(
    <ListItem dense>
      <ListItemIcon>
        <Checkbox
          edge="start"
          disabled
          tabIndex={-1}
          disableRipple
        />
      </ListItemIcon>
      <TextField
        autoFocus={true}
        placeholder="New todo" 
        onBlur={()=>addNewTodo(name)} 
        value={name}
        onChange={handleChange}
        onKeyPress={handleKeyPress}
      />
      <ListItemSecondaryAction>
        <IconButton edge="end" aria-label="delete" disabled>
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  );
}