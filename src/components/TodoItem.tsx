import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Checkbox from '@material-ui/core/Checkbox';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';

type TodoItemProps = {
  todoId: number,
  name: string,
  isChecked: boolean,
  handleToggle: (id: number) => () => void,
  handleDelete: (id: number) => () => void
}

export default function TodoItem({ todoId, name, isChecked, handleToggle, handleDelete}: TodoItemProps){

  return(
    <ListItem dense button onClick={handleToggle(todoId)}>
      <ListItemIcon>
        <Checkbox
          edge="start"
          checked={isChecked}
          tabIndex={-1}
          disableRipple
        />
      </ListItemIcon>
      <ListItemText primary={name} />
      <ListItemSecondaryAction>
        <IconButton edge="end" aria-label="delete" onClick={handleDelete(todoId)}>
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  );
}