declare global {
    var Go: any
    var tryPlaceStone: (col: number, row: number, isBlackTurn: boolean) => 
        {board: number[][], error: string}
    var newGoGame: (boardSize: number) => {board: number[][], error: string}
    var passTurn: (isBlackTurn: boolean) => 
        {isGameOver: boolean, playerOneScore: number, playerTwoScore: number}
}

const initWasm = async function(setIsWasmInitialized: React.Dispatch<React.SetStateAction<boolean>>){
    if (WebAssembly) {
        // WebAssembly.instantiateStreaming is not currently available in Safari
        if (WebAssembly && !WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (
                resp: Response | PromiseLike<Response>, 
                importObject: Record<string, Record<string, WebAssembly.ImportValue>> | undefined
            ) => {
            const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }  

        const go = new Go();
        const wasmMain = await fetch("/wasm/main.wasm");

        WebAssembly.instantiateStreaming(wasmMain, go.importObject).then((result) => {
            go.run(result.instance);
            setIsWasmInitialized(true);
        });
    } else {
    console.log("WebAssembly is not supported in your browser")
    }
}

export default initWasm;