import Image from "next/image";
import GridLines from "@/public/assets/gradient-lines.png";
export default function Home() {
  return (
    <main className="flex h-screen flex-col items-center justify-center p-24 bg-gradient-to-b from-[#7C69D8] to-[#D7603C] relative">
      <div className="absolute w-full h-full left-0 top-0">
        <Image src={GridLines} alt="" className="w-screen h-screen" />
      </div>

      <div className="flex flex-col gap-10 typewriter">
        <h1 className="text-7xl text-white font-bold">Voauth</h1>
        <h4 className="text-3xl font-light text-white animate-pulse text-center">
          Coming Soon...
        </h4>
      </div>
    </main>
  );
}
