"use client";

import Link from "next/link";

export function Notes() {
  const notes = [1, 2, 3, 4];
  return (
    <section className="flex-grow p-5">
      <section className="grid grid-cols-4 grid-rows-[200px] gap-8">
        {notes.map((note) => (
          <Note key={note} index={note} />
        ))}
      </section>
    </section>
  );
}

function Note({ index }: { index: number }) {
  return (
    <Link href={`/notes/${index}`}>
      <div className="p-4 border rounded-md bg-[#FAFAFA] h-full flex flex-col justify-end">
        <span className="mt-auto">Note {index}</span>
      </div>
    </Link>
  );
}
