import React from 'react';

const SCALE = 0.1;
const OFFSET = 80;
export const canvasWidth = window.innerWidth ;
export const canvasHeight = window.innerHeight;

const stoneSpacing = 1.25;
const stoneRadius = 25;
const stoneSvgPath = 
  ` M 0, ${stoneRadius} 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*2},0 
    a ${stoneRadius},${stoneRadius} 0 1,0 ${stoneRadius*-2},0`;
const stone = new Path2D(stoneSvgPath);

enum Intersection {
  None = 0,
  White,
  Black
}

export function drawBoard(ctx: CanvasRenderingContext2D, board: number[][]){
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

export function useGoBoard(): [
  number[][], 
  React.Dispatch<React.SetStateAction<number[][]>>, 
  React.RefObject<HTMLCanvasElement>, 
  number, 
  number
]{
    // 0: nothing 
    // 1: white stone
    // 2: black stone
    const [board, setBoard] = React.useState<number[][]>([[1,1,1,1,1],[0,0,2,1,0]]);
    const canvasRef = React.useRef<HTMLCanvasElement>(null);

    React.useEffect(()=>{
        if(canvasRef)
        {
          console.log("draw");
          const canvasObj = canvasRef.current!;
          const ctx : CanvasRenderingContext2D = canvasObj.getContext('2d')!;
          // clear the canvas area before rendering the board
          ctx.clearRect( 0,0, canvasWidth, canvasHeight );
          drawBoard(ctx, board);
        }
    }, [board]);

    return [ board, setBoard, canvasRef, canvasWidth, canvasHeight ];
}