
const initWasm = async function(){
    if (WebAssembly) {
        // WebAssembly.instantiateStreaming is not currently available in Safari
        if (WebAssembly && !WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (
                resp/*: Response | PromiseLike<Response>*/, 
                importObject/*: Record<string, Record<string, WebAssembly.ImportValue>> | undefined*/
            ) => {
            const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }  

        const go = new global.Go();

        const test = await fetch("/wasm/main.wasm");
        console.log("test",test)

        WebAssembly.instantiateStreaming(test, go.importObject).then((result) => {
            go.run(result.instance);
        });
    } else {
    console.log("WebAssembly is not supported in your browser")
    }
}

export default initWasm;