import React from 'react';
import GoBoard from './components/GoBoard';
import ActionBar from './components/ActionBar';
import initWasm from "./utilities/initWasm";
import useWindowDimensions from './hooks/useWindowDimensions';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  goBoardDiv: (actionBarHeight) => ({
    backgroundColor: "lightblue",
    //textAlign: "center",
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
// TODO: Center board and handle offset in click event
function App() {  
  const windowDimensions = useWindowDimensions();
  const [isWasmInitialized, setIsWasmInitialized] = React.useState<boolean>(false);
  const [isBlackTurn, setIsBlackTurn] = React.useState<boolean>(true);

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
        <div className={classes.goBoardDiv}>
          <GoBoard 
            canvasSize={canvasSize}
            isWasmInitialized={isWasmInitialized}
            isBlackTurn={isBlackTurn}
            setIsBlackTurn={setIsBlackTurn}
          />
        </div>
        <div className={classes.actionBarDiv}>
          <ActionBar 
            actionBarHeight={actionBarHeight} 
            isBlackTurn={isBlackTurn}
          />
        </div>
    </main>
  );
}

export default App;
