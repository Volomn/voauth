import { Suspense } from "react";
import { auth } from "@/lib/auth";
import { Navbar } from "../navbar";
import NoteForm from "./form";
import { appUrl } from "../../dashboard/page";
import { Note } from "@/schema";

async function fetchNote(id: string) {
  const session = await auth();
  const response = await fetch(appUrl(`/api/notes/${id}`), {
    cache: "no-store",
    headers: {
      Authorization: `Bearer ${session?.user.accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error("Unable to fetch note");
  }

  return response.json();
}

export default async function Note({
  params: { id },
}: {
  params: { id: string };
}) {
  const note: Note = await fetchNote(id);
  return (
    <main className="min-h-screen flex flex-col">
      <Navbar note={note} />
      <Suspense
        fallback={
          <h1 className="text-slate-600 text-3xl">Loading content...</h1>
        }
      >
        <NoteForm initialValues={note} />
      </Suspense>
    </main>
  );
}
