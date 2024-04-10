'use client';

import CodeMirror from '@uiw/react-codemirror';
import { useEffect, useMemo, useRef, useState } from 'react';
import { Terminal } from '@xterm/xterm';

import '@xterm/xterm/css/xterm.css';
import { FitAddon } from '@xterm/addon-fit';
import DarijaScript, { FontSizeThemeExtension } from './codemirror/codemirror';
import { ICodingExample } from './examples';
import { githubDark } from '@uiw/codemirror-theme-github';

export default function PlaygroundPage({
  examples,
  selectedIndex,
  onClose = () => {},
  onSelectedIndexChange,
}: Readonly<{
  selectedIndex: number;
  examples: ICodingExample[];
  onClose?: () => void;
  onSelectedIndexChange: (index: number) => void;
}>) {
  const terminal = useMemo(() => {
    return new Terminal({
      convertEol: true,
      lineHeight: 1.5,
    });
  }, []);
  const terminalRef = useRef<HTMLDivElement>(null);
  const [code, setCode] = useState<string>(examples[selectedIndex]?.code ?? ``);

  useEffect(() => {
    if (examples[selectedIndex]) {
      setCode(examples[selectedIndex].code);
    }
  }, [examples, selectedIndex]);

  useEffect(() => {
    if (!terminalRef.current) return;
    terminal.open(terminalRef.current);
    const fitAddon = new FitAddon();
    terminal.loadAddon(fitAddon);
    fitAddon.fit();
  }, [terminal, terminalRef]);

  async function runCode() {
    terminal.clear();
    terminal.writeln('\x1b[32m$ darijascript run code.ds\x1b[0m');
    terminal.write('Running code...');
    const response = await fetch('/api/run-code', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        code,
      }),
    });
    const { output } = await response.json();
    terminal.write('\x1b[2K\r');
    terminal.writeln(`${output}`);
  }

  return (
    <div className="h-full w-full grid grid-cols-2">
      <div className="w-full h-full border-r border-slate-500">
        <PlaygroundSection
          title={
            <CodeEditorTitle
              examples={examples}
              selectedIndex={selectedIndex}
              onClose={onClose}
              onSelectedIndexChange={onSelectedIndexChange}
            />
          }
          actions={<CodeEditorActions onRun={runCode} />}
        >
          <AppCodeEditor
            value={code}
            onChange={(value) => {
              setCode(value);
            }}
          />
        </PlaygroundSection>
      </div>
      <PlaygroundSection title={<>Terminal</>}>
        <div className="h-full bg-black w-full" ref={terminalRef}></div>
      </PlaygroundSection>
    </div>
  );
}

function CodeEditorTitle({
  onClose,
  examples,
  selectedIndex,
  onSelectedIndexChange,
}: Readonly<{
  onClose: () => void;
  examples: ICodingExample[];
  selectedIndex: number;
  onSelectedIndexChange: (index: number) => void;
}>) {
  return (
    <div className="flex items-center space-x-2">
      <button
        onClick={onClose}
        className="border rounded p-1 text-center inline-flex items-center border-slate-500 hover:bg-slate-900 hover:bg-opacity-50"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="currentColor"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M15.75 19.5 8.25 12l7.5-7.5"
          />
        </svg>
      </button>
      <select
        value={selectedIndex}
        onChange={(e) => {
          const index = parseInt(e.target.value);
          onSelectedIndexChange(index);
        }}
        className="bg-slate-950 text-white border border-slate-500 rounded p-1.5 text-sm min-w-52"
      >
        {examples.map((example, index) => (
          <option key={example.name} value={index}>
            {example.name}
          </option>
        ))}
      </select>
    </div>
  );
}

function CodeEditorActions({
  onRun,
}: Readonly<{
  onRun?: () => void;
}>) {
  return (
    <div className="flex space-x-4">
      <button
        onClick={onRun}
        className="text-white bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700 hover:bg-gradient-to-br font-medium rounded text-sm px-3 py-1 text-center"
      >
        Run
      </button>
    </div>
  );
}

function AppCodeEditor({
  value,
  onChange,
}: Readonly<{
  value: string;
  onChange: (value: string) => void;
}>) {
  return (
    <CodeMirror
      value={value}
      lang="darijascript"
      className="h-full w-full"
      theme={githubDark}
      onChange={(value) => {
        onChange(value);
      }}
      data-enable-grammarly="false"
      extensions={[DarijaScript(), FontSizeThemeExtension]}
    />
  );
}

function PlaygroundSection({
  title,
  actions,
  children,
}: Readonly<{
  title: React.ReactNode;
  actions?: React.ReactNode;
  children: React.ReactNode;
}>) {
  return (
    <section className="h-full bg-slate-950 grid grid-rows-[auto,1fr] overflow-hidden">
      <div className="py-2 px-4 border-b border-slate-400 bg-black flex justify-between items-center h-14">
        <h2 className="text-lg font-medium">{title}</h2>
        {actions && <div className="flex justify-end space-x-4">{actions}</div>}
      </div>
      <div>{children}</div>
    </section>
  );
}
