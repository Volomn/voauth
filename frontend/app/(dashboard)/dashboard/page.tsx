import { Navbar } from "./navbar";
import { EmptyNotes } from "./empty-notes";
import { Notes } from "./notes";
export default function Dashboard() {
  return (
    <section className="h-screen flex flex-col">
      <Navbar />
      <Notes />
      {/* <EmptyNotes /> */}
    </section>
  );
}
