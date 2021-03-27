import React from 'react';
import GoBoard from './components/GoBoard';
import ActionBar from './components/ActionBar';
import initWasm from "./utilities/initWasm";
import useWindowDimensions from './hooks/useWindowDimensions';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

const useStyles = makeStyles((theme) => ({
  goBoardDiv: (actionBarHeight) => ({
    backgroundColor: theme.palette.primary.main,
    position: "absolute",
    top: "0px",
    bottom: `${actionBarHeight}px`,
    width: "100%"
  }),
  actionBarDiv: (actionBarHeight) => ({
    position: "absolute",
    height: `${actionBarHeight}px`,
    bottom: "0px",
    width: "100%"
  }),
}));

// TODO: Handle window resize
function App() {  
  const windowDimensions = useWindowDimensions();
  const [isWasmInitialized, setIsWasmInitialized] = React.useState<boolean>(false);
  const [isBlackTurn, setIsBlackTurn] = React.useState<boolean>(true);
  const [newGameFlag, setNewGameFlag] = React.useState<boolean>(false);

  const createNewGame = () => {setNewGameFlag(!newGameFlag)};

  const actionBarHeight = 100;
  let canvasSize: number;

  if(windowDimensions.height > windowDimensions.width){
    canvasSize = windowDimensions.width;
  }else{
    canvasSize = windowDimensions.height - actionBarHeight;
  }

  React.useEffect(()=>{
    initWasm(setIsWasmInitialized)
  }, []);

  const classes = useStyles(actionBarHeight);

  return (
    <main>
        <Grid 
          container 
          className={classes.goBoardDiv} 
          justify="center" 
          alignItems="center"
        >
          <Grid item>
            <GoBoard 
              canvasSize={canvasSize}
              isWasmInitialized={isWasmInitialized}
              isBlackTurn={isBlackTurn}
              setIsBlackTurn={setIsBlackTurn}
              newGameFlag={newGameFlag}
            />
          </Grid>
        </Grid>
        <div className={classes.actionBarDiv}>
          <ActionBar 
            actionBarHeight={actionBarHeight} 
            isBlackTurn={isBlackTurn}
            setIsBlackTurn={setIsBlackTurn}
            canvasSize={canvasSize}
            createNewGame={createNewGame}
          />
        </div>
    </main>
  );
}

export default App;
