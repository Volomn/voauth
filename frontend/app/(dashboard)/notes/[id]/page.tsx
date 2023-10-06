import { Box, TextInput, Textarea } from "@mantine/core";
import { Navbar } from "../navbar";

export default function Note() {
  return (
    <main className="min-h-screen flex flex-col">
      <Navbar />
      <Box className="p-5">
        <TextInput
          size="xl"
          placeholder="Title"
          className="border-none outline-none"
          classNames={{
            input: "border-none outline-none text-5xl",
          }}
          unstyled
        />

        <Textarea
          className="mt-5 leading-loose"
          placeholder="Note content..."
          unstyled
          classNames={{
            input:
              "border-none outline-none text-xl resize-none w-full min-h-[60vh] font-normal",
          }}
        />
      </Box>
    </main>
  );
}
