import React from 'react';
import Coord from '../interfaces/Coord';

// Path2D for a Heart SVG
const heartSVG = "M0 200 v-200 h200 a100,100 90 0,1 0,200 a100,100 90 0,1 -200,0 z"
const SVG_PATH = new Path2D(heartSVG);

// Scaling Constants for Canvas
const SCALE = 0.1;
const OFFSET = 80;
export const canvasWidth = window.innerWidth ;
export const canvasHeight = window.innerHeight;

export function draw(ctx: any, location: any){
  console.log("attempting to draw")
  ctx.fillStyle = 'red';
  ctx.shadowColor = 'blue';
  ctx.shadowBlur = 15;
  ctx.save();
  ctx.scale(SCALE, SCALE);
  ctx.translate(location.x / SCALE - OFFSET, location.y / SCALE - OFFSET);
  ctx.rotate(225 * Math.PI / 180);
  ctx.fill(SVG_PATH);
  // .restore(): Canvas 2D API restores the most recently saved canvas state
  ctx.restore();  
};

export function useCanvas(): [
  Coord[], 
  React.Dispatch<React.SetStateAction<Coord[]>>, 
  React.RefObject<HTMLCanvasElement>, 
  number, 
  number
]{
    const canvasRef = React.useRef<HTMLCanvasElement>(null);
    const [coordinates, setCoordinates] = React.useState<Coord[]>([]);

    React.useEffect(()=>{
        if(canvasRef)
        {
          const canvasObj = canvasRef.current!;
          const ctx : CanvasRenderingContext2D = canvasObj.getContext('2d')!;
          // clear the canvas area before rendering the coordinates held in state
          ctx.clearRect( 0,0, canvasWidth, canvasHeight );

          // draw all coordinates held in state
          coordinates.forEach((coordinate)=>{draw(ctx, coordinate)});
        }
    });

    return [ coordinates, setCoordinates, canvasRef, canvasWidth, canvasHeight ];
}