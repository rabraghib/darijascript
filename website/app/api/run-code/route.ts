import { promisify } from 'util';
import { writeFile } from 'fs/promises';
import { exec } from 'child_process';
import { temporaryFile } from 'tempy';
import { createWriteStream } from 'fs';

const execute = promisify(exec);

export async function POST(req: Request) {
  const { code } = await req.json();
  if (code.includes('dakhel')) {
    return Response.json({
      output: 'Error: Cannot run code with dakhel call',
    });
  }
  await setupDarijaScript();
  return Response.json(await runCodeWithInput(code));
}

async function setupDarijaScript() {
  try {
    const { stderr } = await execute('./darijascript --version');
    if (stderr) {
      throw new Error(stderr);
    }
  } catch (error) {
    const res = await fetch(
      'https://github.com/rabraghib/darijascript/releases/latest/download/darijascript-bin'
    );
    const outFile = createWriteStream('./darijascript');
    res.body?.pipeTo(
      new WritableStream({
        write(chunk) {
          outFile.write(chunk);
        },
        close() {
          outFile.close();
        },
      })
    );
    await execute('chmod +x ./darijascript');
  }
}

async function runCodeWithInput(code: string) {
  try {
    const codeFilePath = temporaryFile({ extension: 'ds' });
    await writeFile(codeFilePath, code);

    const { stdout, stderr } = await execute(
      `darijascript run ${codeFilePath}`
    );

    if (stderr) {
      throw new Error(stderr);
    }

    return {
      output: stdout,
    };
  } catch (error) {
    return {
      output: `Error running code: ${error}`,
    };
  }
}
