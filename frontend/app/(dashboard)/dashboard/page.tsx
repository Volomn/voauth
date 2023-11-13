import { Navbar } from "./navbar";
import { EmptyNotes } from "./empty-notes";
import { Notes } from "./notes";
import { auth } from "@/lib/auth";

const baseUrl = process.env.APP_BASE_URL; //localhost:7005
export function appUrl(url: string) {
  return new URL(url, baseUrl);
}

async function fetchNotes() {
  const session = await auth();
  const response = await fetch(appUrl("/api/notes/"), {
    cache: "no-store",
    headers: {
      Authorization: `Bearer ${session?.user.accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error("Unable to fetch notes");
  }

  return response.json();
}

export default async function Dashboard() {
  const notes = await fetchNotes();

  return (
    <section className="min-h-screen flex flex-col">
      <Navbar />
      {notes.length > 0 ? <Notes notes={notes} /> : <EmptyNotes />}
    </section>
  );
}
