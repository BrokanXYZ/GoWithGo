import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import MatchAction from '../interfaces/MatchAction';

const useStyles = makeStyles((theme) => ({
  canvas: {
    backgroundColor: "lightblue",
  },
}));

enum Intersection {
  None = 0,
  BlackStone,
  WhiteStone
}

function GoBoard() {

    const classes = useStyles();

    const canvasWidth = 500;
    const canvasHeight = 500;
    const boardSize = 9;

    // Space b/w edge of board and canvas
    const gridBuffer = 50;
    
    const rowSpacing = canvasHeight/(boardSize+1);
    const columnSpacing = canvasWidth/(boardSize+1);

    const stoneRadius = 20;
    const stoneSvgPath = 
      ` M 0, ${stoneRadius} 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;
    const stone = new Path2D(stoneSvgPath);

    const [isBlackTurn, setIsBlackTurn] = React.useState<boolean>(true);
    const [matchHistory, setMatchHistory] = React.useState<MatchAction[]>([]); 

    const [board, setBoard] = React.useState<number[][]>((): number[][] => {
      let initBoard = [];
      for(let i=0; i<boardSize; i++)
      {
        let row = [];
        for(let j=0; j<boardSize; j++)
        {
          row.push(0);
        }
        initBoard.push(row);
      }
      return initBoard
    });

    const canvasRef = React.useRef<HTMLCanvasElement>(null);
    let canvasElement: HTMLCanvasElement | null = null;
    let ctx: CanvasRenderingContext2D;

    React.useEffect(()=>{
        if(canvasRef)
        {
            canvasElement = canvasRef.current!;
            ctx = canvasElement.getContext('2d')!;
            DrawBoard(ctx, board);
        }
    }, [canvasRef, board]);

    function DrawBoard(ctx: CanvasRenderingContext2D, board: number[][]){
      ctx.clearRect(0, 0, canvasWidth, canvasHeight);
      drawBoardLines(ctx, boardSize);
      drawBoardStones(ctx, board);
    };

    function drawBoardLines(ctx: CanvasRenderingContext2D, boardSize: number)
    {
        for(let i=1; i<=boardSize; i++)
        {
            ctx.beginPath();
            ctx.moveTo(columnSpacing, rowSpacing*i);
            ctx.lineTo(columnSpacing*boardSize, rowSpacing*i);
            ctx.stroke();

            ctx.beginPath();
            ctx.moveTo(columnSpacing*i, rowSpacing);
            ctx.lineTo(columnSpacing*i, rowSpacing*boardSize);
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
            case Intersection.WhiteStone:
              ctx.fillStyle = 'white';
              ctx.translate(columnSpacing*j+columnSpacing-stoneRadius,rowSpacing*i+rowSpacing-stoneRadius);
              ctx.fill(stone);
              resetContextTranslation(ctx);
              break;
            case Intersection.BlackStone:
              ctx.fillStyle = 'black';
              ctx.translate(columnSpacing*j+columnSpacing-stoneRadius,rowSpacing*i+rowSpacing-stoneRadius);
              ctx.fill(stone);
              resetContextTranslation(ctx);
              break;
            default:
              console.error(`Unhandled intersection of value: ${intersection}`);
              break;
          }
        })
      });
    }

    function resetContextTranslation(ctx: CanvasRenderingContext2D)
    {
      ctx.setTransform(1, 0, 0, 1, 0, 0);
    }

    const handleCanvasClick=(event: React.MouseEvent)=>{
      const mousePositionX: number = event.clientX;
      const mousePositionY: number = event.clientY;
      const cellX = (mousePositionX-(gridBuffer/2))/columnSpacing;
      const cellY = (mousePositionY-(gridBuffer/2))/rowSpacing;

      if((cellX < 0 || cellX > boardSize) || (cellY < 0 || cellY > boardSize))
      {
        console.log("Cursor out of bounds");
      }
      else
      {
        const { updatedBoard, error } = placeStone(cellX, cellY, [...board], isBlackTurn);
        if(error)
        {
          console.log(`Unable to place stone\nReason: ${error}`);
        }
        else
        {
          setIsBlackTurn(!isBlackTurn);
          setBoard(updatedBoard);
        }
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
