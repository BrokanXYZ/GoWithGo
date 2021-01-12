import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  canvas: {
    backgroundColor: "lightblue",
    //border: "5px",
    //borderColor: "black"
  },
}));

enum Intersection {
  None = 0,
  White,
  Black
}

function GoBoard() {

    const classes = useStyles();

    const canvasWidth = 500;
    const canvasHeight = 500;

    const boardRows = 9;
    const boardColumns = 9;

    const stoneRadius = 25;
    const stoneSpacing = 1;
    const stoneSvgPath = 
      ` M 0, ${stoneRadius} 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;
    const stone = new Path2D(stoneSvgPath);

    // 0: nothing 
    // 1: white stone
    // 2: black stone
    const [board, setBoard] = React.useState<number[][]>(
        new Array(boardRows).fill(
            new Array(boardColumns).fill(1)
        )
    );
    const canvasRef = React.useRef<HTMLCanvasElement>(null);
    let canvasElement: HTMLCanvasElement | null = null;
    let ctx: CanvasRenderingContext2D;

    React.useEffect(()=>{
        if(canvasRef)
        {
            canvasElement = canvasRef.current!;
            ctx = canvasElement.getContext('2d')!;
            drawBoardIntersections(ctx, boardRows, boardColumns);
        }
    }, [canvasRef]);
    
    React.useEffect(()=>{
      if(canvasRef)
      {
        console.log("draw");
        
        // clear the canvas area before rendering the board
        //ctx.clearRect( 0,0, canvasWidth, canvasHeight );
        drawBoard(ctx, board);
      }
    }, [board]);

    function drawBoardIntersections(ctx: CanvasRenderingContext2D, boardRows: number, boardColumns: number)
    {
        const rowSpacing = canvasHeight/(boardRows+1);
        const columnSpacing = canvasWidth/(boardColumns+1);

        for(let i=1; i<=boardRows; i++)
        {
            ctx.beginPath();
            ctx.moveTo(columnSpacing, rowSpacing*i);
            ctx.lineTo(columnSpacing*boardColumns, rowSpacing*i);
            ctx.stroke();

            ctx.beginPath();
            ctx.moveTo(columnSpacing*i, rowSpacing);
            ctx.lineTo(columnSpacing*i, rowSpacing*boardRows);
            ctx.stroke();
        }
    }

    function drawBoard(ctx: CanvasRenderingContext2D, board: number[][]){
      board.forEach((row, i) => {
        row.forEach((intersection, j) => {
          switch(intersection)
          {
            case Intersection.None:
              break;
            case Intersection.White:
              ctx.fillStyle = 'white';
              ctx.setTransform(1, 0, 0, 1, 0, 0);
              ctx.translate(stoneRadius*2*j*stoneSpacing,stoneRadius*2*i*stoneSpacing);
              ctx.fill(stone);
              break;
            case Intersection.Black:
              ctx.fillStyle = 'black';
              ctx.setTransform(1, 0, 0, 1, 0, 0);
              ctx.translate(stoneRadius*2*j*stoneSpacing,stoneRadius*2*i*stoneSpacing);
              ctx.fill(stone);
              break;
            default:
              console.error(`Unhandled intersection of value: ${intersection}`);
              break;
          }
        })
      });
    };

    const handleCanvasClick=(event: React.MouseEvent)=>{
      // on each click get current mouse location 
      //const currentCoord: Coord = { x: event.clientX, y: event.clientY };
      // add the newest mouse location to an array in state 
      //setCoordinates([...coordinates, currentCoord]);
    };

    const handleClearCanvas=(event: React.MouseEvent)=>{
      //setCoordinates([]);
    };

  return (
    <canvas 
        className={classes.canvas}
        ref={canvasRef}
        width={canvasWidth}
        height={canvasHeight}
        onClick={handleCanvasClick} 
    />
  );
}

export default GoBoard;
