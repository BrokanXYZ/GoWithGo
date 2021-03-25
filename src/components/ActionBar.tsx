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
  isBlackTurn: boolean,
  canvasSize: number
}

function ActionBar({actionBarHeight, isBlackTurn, canvasSize}: ActionBarProps) {

  const classes = useStyles(actionBarHeight);

  // Constant determined by eyeballin' it
  const stoneRadius = canvasSize * 0.0389;
  const stoneSvgPath = 
  ` M 0, ${stoneRadius} 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;

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
      <Grid item>
        <svg width={stoneRadius*2} height={stoneRadius*2}>
          <path d={stoneSvgPath} fill={isBlackTurn ? "black" : "white"}/>
        </svg>
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
