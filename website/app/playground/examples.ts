'use server';

import { readFile } from 'fs/promises';
import path from 'path';

export interface ICodingExample {
  name: string;
  code: string;
}

const rootPath = path.join(__dirname, '..', '..', '..', '..', '..');

const examples = [
  {
    name: 'Hello, World!',
    codePath: path.join(rootPath, 'examples/hello_world.ds'),
  },
  {
    name: 'FizzBuzz',
    codePath: path.join(rootPath, 'examples/fizz_buzz.ds'),
  },
  {
    name: 'Numerical Integration',
    codePath: path.join(rootPath, 'examples/integral.ds'),
  },
  {
    name: 'Non Linear Equation',
    codePath: path.join(rootPath, 'examples/non_linear_eqq.ds'),
  },
];

export async function getExamples(): Promise<ICodingExample[]> {
  return await Promise.all(
    examples.map(async (example) => {
      const codeBuffer = await readFile(example.codePath);
      return {
        name: example.name,
        code: codeBuffer.toString(),
      };
    })
  );
}
