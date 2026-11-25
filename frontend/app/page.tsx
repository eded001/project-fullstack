"use client";

import Image from "next/image";

export default function Home() {
  return (
    <main className="w-full">
      <header className="w-full bg-[#E61E26] p-5 flex justify-between">
        <Image
          src="/logo-a-verdade.svg"
          alt="A Verdade"
          width={250}
          height={100}
          priority
          className="h-auto w-64"
        />

        <h1 className="font-typewriter text-5xl text-white">System</h1>
      </header>
    </main>
  );
}
