import React from 'react';
import GoBoard from './components/GoBoard';
import initWasm from "./utilities/initWasm";

function App() {

  const [isWasmInitialized, setIsWasmInitialized] = React.useState<boolean>(false);

  React.useEffect(()=>{
    initWasm(setIsWasmInitialized)
  }, []);

  return (
    <main>
        <GoBoard isWasmInitialized={isWasmInitialized}/>
    </main>
  );
}

export default App;
