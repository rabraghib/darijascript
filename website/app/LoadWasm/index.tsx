import './wasm_exec.js';
import './wasmTypes.d.ts';

import React, { useEffect } from 'react';

async function loadWasm(writeOutput: (output: string) => void): Promise<void> {
  const goWasm = new window.Go();
  const result = await WebAssembly.instantiateStreaming(
    fetch('darijascript-bin-web.wasm'),
    goWasm.importObject
  );

  let outputBuf = '';
  const decoder = new TextDecoder('utf-8');
  window.fs.writeSync = function (fd, buf) {
    const output = decoder.decode(buf);
    writeOutput(output);
    return buf.length;
  };
  goWasm.run(result.instance);
}

export const LoadWasm: React.FC<
  React.PropsWithChildren<{
    writeOutput: (output: string) => void;
  }>
> = (props) => {
  const [isLoading, setIsLoading] = React.useState(true);

  useEffect(() => {
    loadWasm(props.writeOutput).then(() => {
      setIsLoading(false);
    });
  }, [props.writeOutput]);

  // if (isLoading) {
  //   return <div>loading WebAssembly...</div>;
  // } else {
  // }
  return <React.Fragment>{props.children}</React.Fragment>;
};
