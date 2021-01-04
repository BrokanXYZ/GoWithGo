import React from 'react';

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

function App() {
  return (
    <Container maxWidth="sm">
      <Grid container justify="center" alignItems="center">
        <Grid item xs={12}>
          <Typography variant="h2">GoWithGo!</Typography>
        </Grid>
        <Grid item xs={12}>
          <Button>Do Something</Button>
        </Grid>
      </Grid>
    </Container>
  );
}

export default App;
