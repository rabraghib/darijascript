import type { Metadata } from 'next';
import Image from 'next/image';
import { Inter } from 'next/font/google';
import './globals.css';
import Link from 'next/link';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'DarijaScript',
  description: 'The greatest language that ever was or will be!',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="text-white bg-black">
          <Header />
          <div className="min-h-screen pt-16">{children}</div>
          <Footer />
        </div>
      </body>
    </html>
  );
}

function Header() {
  return (
    <header>
      <nav className="fixed overflow-hidden z-20 w-full bg-slate-900 bg-opacity-20 border-b border-slate-700 backdrop-blur-2xl">
        <div className="px-6 m-auto max-w-6xl">
          <div className="flex flex-wrap items-center justify-between py-3">
            <div className="items-center flex justify-between w-auto">
              <Link href="/">
                <Image
                  alt="DarijaScript Logo"
                  src="/vercel.svg"
                  className="invert"
                  width={100}
                  height={24}
                ></Image>
              </Link>
            </div>
            <div className="w-fit flex-nowrap justify-end items-center flex h-fit">
              <div className="tracking-wide text-sm text-gray-300 pr-2">
                <a
                  target="_blank"
                  href="https://github.com/rabraghib/darijascript"
                  className="block px-4 transition hover:text-indigo-400"
                >
                  <span>GitHub</span>
                </a>
              </div>

              <Link
                href="/playground"
                className="text-sm font-medium inline-flex items-center justify-center rounded px-4 py-2 shadow-lg bg-blue-600 hover:bg-blue-700 text-white"
              >
                Playground!
              </Link>
            </div>
          </div>
        </div>
      </nav>
    </header>
  );
}

function Footer() {
  return (
    <div className="bg-slate-900 bg-opacity-30 text-lg text-center p-4">
      Made with <span className="text-red-600">&hearts;</span> by{' '}
      <a
        target="_blank"
        href="https://rabraghib.me/"
        className="underline whitespace-nowrap"
      >
        Rabyâ Raghib
      </a>
    </div>
  );
}
