import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  canvas: {
    backgroundColor: theme.palette.secondary.main,
    borderRadius: "2px"
  },
}));

enum Intersection {
  None = 0,
  BlackStone,
  WhiteStone
}

type GoBoardProps = {
  isWasmInitialized: boolean,
  canvasSize: number,
  isBlackTurn: boolean,
  setIsBlackTurn: Function,
  newGameFlag: boolean
}

function GoBoard({isWasmInitialized, canvasSize, isBlackTurn, setIsBlackTurn, newGameFlag}: GoBoardProps) {

    const classes = useStyles();

    const canvasWidth = canvasSize;
    const canvasHeight = canvasSize;
    const boardSize = 9;

    // Space b/w edge of board and canvas
    const gridBuffer = 50;
    
    const rowSpacing = canvasHeight/(boardSize+1);
    const columnSpacing = canvasWidth/(boardSize+1);

    // Constant determined by eyeballin' it
    const stoneRadius = canvasSize * 0.0389;
    const stoneSvgPath = 
      ` M 0, ${stoneRadius} 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
        a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;
    const stone = new Path2D(stoneSvgPath);

    const [board, setBoard] = React.useState<number[][]>([]);

    // Create new game when...
    //    WASM has been initialized OR 
    //    flag is triggered
    React.useEffect(()=>{
      if(isWasmInitialized) {
        const {board, error} = newGoGame(boardSize);

        if(error !== null) {
          console.error(error);
        } else {
          setIsBlackTurn(true);
          setBoard(board);
        }
      }
    }, [isWasmInitialized, newGameFlag])

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

    const handleCanvasClick=(event: React.MouseEvent, isBlackTurn: boolean)=>{

      let mousePositionX: number = event.clientX;
      let mousePositionY: number = event.clientY;

      if(canvasRef && canvasRef.current) {
        if(canvasRef.current.offsetLeft) {
          mousePositionX -= canvasRef.current.offsetLeft;
        }
        if(canvasRef.current.offsetTop) {
          mousePositionY -= canvasRef.current.offsetTop;
        }
      }

      const cellCol = (mousePositionX-(gridBuffer/2))/columnSpacing;
      const cellRow = (mousePositionY-(gridBuffer/2))/rowSpacing;

      if((cellCol < 0 || cellCol > boardSize) || (cellRow < 0 || cellRow > boardSize))
      {
        console.log("Cursor out of bounds");
      }
      else
      {
        const { board, error } = tryPlaceStone(cellCol, cellRow, isBlackTurn);
        if(error)
        {
          console.log(`Unable to place stone\nReason: ${error}`);
        }
        else
        {
          setIsBlackTurn(!isBlackTurn);
          setBoard(board);
        }
      }
    };

  return (
    <canvas 
        className={classes.canvas}
        ref={canvasRef}
        width={canvasWidth}
        height={canvasHeight}
        onClick={(event)=>handleCanvasClick(event, isBlackTurn)} 
    />
  );
}

export default GoBoard;
