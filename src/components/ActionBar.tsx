import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';


const useStyles = makeStyles((theme) => ({
  rootGrid: (actionBarHeight) => ({
    backgroundColor: theme.palette.primary.main,
    height: `${actionBarHeight}px`,
    fontSize: "16px"
  }),
}));

type ActionBarProps = {
  actionBarHeight: number,
  isBlackTurn: boolean
}

function ActionBar({actionBarHeight, isBlackTurn}: ActionBarProps) {

  const classes = useStyles(actionBarHeight);

  return (
    <Grid 
      container 
      justify="space-around" 
      alignItems="center" 
      className={classes.rootGrid}
    >
      <Grid item>
        <Button variant="contained">
          New Game
        </Button>
      </Grid>
      <Grid item style={{color: "white"}}>
        Turn: 
      </Grid>
      <Grid item>
        <Button variant="contained">
          Pass
        </Button>
      </Grid>
    </Grid>
  );
}

export default ActionBar;
