'use client';

import { ICodingExample } from './examples';
import CodeMirror, { EditorView } from '@uiw/react-codemirror';
import Playground from './playground';
import { githubDark } from '@uiw/codemirror-theme-github';
import DarijaScript, { FontSizeThemeExtension } from './codemirror/codemirror';
import { useEffect, useState } from 'react';

export default function PlaygroundManager({
  examples,
}: {
  examples: ICodingExample[];
}) {
  const [selectedExample, setSelectedExample] = useState<number | null>(null);

  useEffect(() => {
    window.scrollTo(0, 0);
  }, [selectedExample]);

  if (selectedExample !== null) {
    return (
      <Playground
        examples={examples}
        selectedIndex={selectedExample}
        onClose={() => setSelectedExample(null)}
        onSelectedIndexChange={(index) => setSelectedExample(index)}
      />
    );
  }
  return (
    <div className="grid grid-rows-[auto,1fr] px-6 py-10 max-w-6xl mx-auto gap-6 items-center justify-center min-h-full w-full">
      <div className="text-center">
        <h1 className="text-4xl font-bold">Playground</h1>
        <p className="text-lg text-gray-600">Choose an example to run</p>
      </div>
      <div className="grid lg:grid-cols-2 gap-6">
        {examples.map((example, index) => (
          <ExampleCard
            key={example.name}
            name={example.name}
            code={example.code}
            onTryClick={() => setSelectedExample(index)}
          />
        ))}
      </div>
    </div>
  );
}

function ExampleCard({
  name,
  code,
  onTryClick,
}: Readonly<{
  name: string;
  code: string;
  onTryClick?: () => void;
}>) {
  'use client';
  return (
    <div className="bg-gray-800 p-4 rounded-lg">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-bold">{name}</h2>
        <button
          className="text-white bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700 hover:bg-gradient-to-br font-medium rounded text-sm px-3 py-1 text-center"
          onClick={onTryClick}
        >
          Try it!
        </button>
      </div>
      <CodeMirror
        value={code}
        lang="darijascript"
        className="h-44 w-full mt-4"
        theme={githubDark}
        readOnly={true}
        data-enable-grammarly="false"
        extensions={[
          DarijaScript(),
          FontSizeThemeExtension,
          EditorView.lineWrapping,
        ]}
      />
    </div>
  );
}
