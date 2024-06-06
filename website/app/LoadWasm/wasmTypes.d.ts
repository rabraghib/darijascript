declare global {
  export interface Window {
    Go: any;
    runDarijaScript: (code: string) => void;
    fs: {
      writeSync: (fd: number, buf: Uint8Array) => number;
    };
  }
}

export {};
