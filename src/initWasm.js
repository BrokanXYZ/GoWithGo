import main from './wasm/main.wasm';

const initWasm = function(){
    if (WebAssembly) {
        /*// WebAssembly.instantiateStreaming is not currently available in Safari
        if (WebAssembly && !WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }  

        const go = new global.Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        });*/
        console.log("main",main)
    } else {
    console.log("WebAssembly is not supported in your browser")
    }
}

export default initWasm;