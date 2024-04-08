import Image from 'next/image';

export default function Home() {
  return (
    <section className="container min-h-[=100%-8rem] mx-auto pb-16">
      <div className="px-6 pt-16 pb-6 mx-auto">
        <div className="flex flex-col w-full mb-2 text-center">
          <h1 className="mb-4 text-5xl font-bold tracking-tighter text-white lg:text-8xl md:text-7xl">
            DarijaScript
          </h1>
          <p className="mx-auto text-xl font-normal leading-relaxed text-gray-300 lg:w-2/3">
            The greatest language that ever was or will be!
          </p>
        </div>
      </div>
      <div className="flex flex-col max-w-6xl mx-auto items-center overflow-hidden justify-center py-8 rounded-lg p-3">
        <Image
          alt="hero"
          src="/code-example.webp"
          width={1920}
          height={1080}
          className="object-cover object-center w-full border-slate-600 border rounded-lg shadow-md shadow-slate-700"
        ></Image>
      </div>
    </section>
  );
}
