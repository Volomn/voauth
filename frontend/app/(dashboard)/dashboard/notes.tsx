// "use client";

// import { useFetchNotes } from "@/api/hooks/notes";
// import { LoadingOverlay, Skeleton } from "@mantine/core";
import Link from "next/link";
// import { EmptyNotes } from "./empty-notes";
import { Note } from "@/schema";

export function Notes({ notes }: { notes: Note[] }) {
  // const { data, isLoading } = useFetchNotes();

  // if (isLoading) {
  //   return <Skeleton className="h-full w-full" />;
  // }
  return (
    <section className="p-5 relative overflow-y-auto">
      <section className="grid grid-cols-4 gap-8 h-full">
        {notes.map((note) => (
          <Note key={note.uuid} note={note} />
        ))}
      </section>
    </section>
  );
}

function Note({ note }: { note: Note }) {
  return (
    <Link href={`/notes/${note.uuid}`}>
      <div className="p-4 border rounded-md bg-[#FAFAFA] h-48 flex flex-col justify-end">
        <span className="mt-auto">{note.title}</span>
      </div>
    </Link>
  );
}
