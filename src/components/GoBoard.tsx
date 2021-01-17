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
    const gridBuffer = 50;
    
    const rowSpacing = canvasHeight/(boardRows+1);
    const columnSpacing = canvasWidth/(boardColumns+1);

    const stoneRadius = 20;
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
            InitializeBoard(ctx, board);
        }
    }, [canvasRef]);

    function InitializeBoard(ctx: CanvasRenderingContext2D, board: number[][]){
      drawBoardIntersections(ctx, boardRows, boardColumns);
      drawBoardStones(ctx, board);
    };

    function drawBoardIntersections(ctx: CanvasRenderingContext2D, boardRows: number, boardColumns: number)
    {
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

    function drawBoardStones(ctx: CanvasRenderingContext2D, board: number[][])
    {
      board.forEach((row, i) => {
        row.forEach((intersection, j) => {
          switch(intersection)
          {
            case Intersection.None:
              break;
            case Intersection.White:
              ctx.fillStyle = 'white';
              ctx.setTransform(1, 0, 0, 1, 0, 0);
              ctx.translate(columnSpacing*j+columnSpacing-stoneRadius,rowSpacing*i+rowSpacing-stoneRadius);
              ctx.fill(stone);
              break;
            case Intersection.Black:
              ctx.fillStyle = 'black';
              ctx.setTransform(1, 0, 0, 1, 0, 0);
              ctx.translate(columnSpacing*j+columnSpacing-stoneRadius,rowSpacing*i+rowSpacing-stoneRadius);
              ctx.fill(stone);
              break;
            default:
              console.error(`Unhandled intersection of value: ${intersection}`);
              break;
          }
        })
      });
    }

    const handleCanvasClick=(event: React.MouseEvent)=>{
      const mousePositionX: number = event.clientX;
      const mousePositionY: number = event.clientY;
      const cellX = (mousePositionX-(gridBuffer/2))/columnSpacing;
      const cellY = (mousePositionY-(gridBuffer/2))/rowSpacing;

      if((cellX < 0 || cellX > boardColumns) || (cellY < 0 || cellY > boardRows))
      {
        console.log("out of bounds");
      }
      else
      {
        console.log(`(${Math.floor(cellX)},${Math.floor(cellY)})`);
      }

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
