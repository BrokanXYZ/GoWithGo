import React from 'react';

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

import { useCanvas } from './hooks/useCanvas';
import Coord from './interfaces/Coord';

const useStyles = makeStyles((theme) => ({
  canvas: {
    backgroundColor: "cornsilk",
    border: "5px",
    borderColor: "black"
  },
}));

function App() {

  const classes = useStyles();
  const [ coordinates, setCoordinates, canvasRef, canvasWidth, canvasHeight ] = useCanvas();

  const handleCanvasClick=(event: React.MouseEvent)=>{
    // on each click get current mouse location 
    const currentCoord: Coord = { x: event.clientX, y: event.clientY };
    // add the newest mouse location to an array in state 
    setCoordinates([...coordinates, currentCoord]);
  };

  const handleClearCanvas=(event: React.MouseEvent)=>{
    setCoordinates([]);
  };

  return (
    <main>
        <canvas 
          className={classes.canvas}
          ref={canvasRef}
          width={canvasWidth}
          height={canvasHeight}
          onClick={handleCanvasClick} 
        />
        <Button onClick={handleClearCanvas}>clear</Button>
    </main>
  );
}

export default App;
