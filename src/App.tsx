import React from 'react';

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';

import TodoList from './components/TodoList';

function App() {
  return (
    <Container maxWidth="sm">
      <Grid container justify="center" alignItems="center">
        <Typography variant="h2">Todo List</Typography>
      </Grid>
      <TodoList />
    </Container>
  );
}

export default App;
