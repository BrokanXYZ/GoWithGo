import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import AlertDialog from './AlertDialog';

const useStyles = makeStyles((theme) => ({
  rootGrid: (actionBarHeight) => ({
    backgroundColor: theme.palette.primary.main,
    height: `${actionBarHeight}px`,
    fontSize: "16px"
  }),
  nextTurnGrid: {
    backgroundColor: theme.palette.secondary.main,
    padding: "8px 8px 2px 8px",
    borderRadius: "2px"
  }
}));

type ActionBarProps = {
  actionBarHeight: number,
  isBlackTurn: boolean,
  setIsBlackTurn: Function,
  canvasSize: number,
  createNewGame: Function
}

function ActionBar({actionBarHeight, isBlackTurn, setIsBlackTurn, canvasSize, createNewGame}: ActionBarProps) {

  const classes = useStyles(actionBarHeight);

  const [isOpen, setIsOpen] = React.useState<boolean>(false);

  // Constant determined by eyeballin' it
  const stoneRadius = canvasSize * 0.0389;
  const stoneSvgPath = 
  ` M 0, ${stoneRadius} 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;

  const handlePass = () => {
    const { isGameOver, playerOneScore, playerTwoScore } = passTurn(isBlackTurn);
    setIsBlackTurn(!isBlackTurn);
    console.log(`isGameOver: ${isGameOver}`);
    console.log(`playerOneScore: ${playerOneScore}`);
    console.log(`playerTwoScore: ${playerTwoScore}`);
  }

  return (
    <>
      <AlertDialog 
        isOpen={isOpen}
        setIsOpen={setIsOpen}
        title="New Game"
        body="Would you like to start a new game?"
        action={()=>createNewGame()}
      />
      <Grid 
        container 
        justify="space-around" 
        alignItems="center" 
        className={classes.rootGrid}
      >
        <Grid item>
          <Button color="secondary" variant="contained" onClick={()=>setIsOpen(true)}>
            New Game
          </Button>
        </Grid>
        <Grid item className={classes.nextTurnGrid}>
          <svg width={stoneRadius*2} height={stoneRadius*2}>
            <path d={stoneSvgPath} fill={isBlackTurn ? "black" : "white"}/>
          </svg>
        </Grid>
        <Grid item>
          <Button color="secondary" variant="contained" onClick={handlePass}>
            Pass
          </Button>
        </Grid>
      </Grid>
    </>
  );
}

export default ActionBar;
