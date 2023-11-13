"use client";
import debounce from "lodash.debounce";
import { Box, TextInput, Textarea } from "@mantine/core";
import { z } from "zod";
import { useForm, zodResolver } from "@mantine/form";
import { useAddNote, useUpdateNote } from "@/api/hooks/notes";
import { Note } from "@/schema";
import { useEffect } from "react";

const notesSchema = z.object({
  title: z.string().min(1, "Enter a title"),
  content: z.string().min(1, "Enter content"),
  isPublic: z.boolean(),
});

export type NoteSchema = z.infer<typeof notesSchema>;
export default function NoteForm({ initialValues }: { initialValues: Note }) {
  const { mutate: updateNote, isLoading: updateNoteLoading } = useUpdateNote(
    initialValues.uuid
  );

  const noteForm = useForm({
    initialValues: {
      title: initialValues.title,
      content: initialValues.content,
      isPublic: initialValues.isPublic,
      isArchived: initialValues.isArchived,
      isFavorite: initialValues.isFavorite
    },
    validate: zodResolver(notesSchema),
  });

  // Debounce function
  //   let timeout: NodeJS.Timeout;
  // const debounceApiCall = ({
  //   title,
  //   content,
  // }: {
  //   title: string;
  //   content: string;
  // }) => {
  //   debounce(
  //     updateNote({ title, content, isPublic: noteForm.values.isPublic }),
  //     300
  //   );
  // };

  useEffect(
    function () {
      if (noteForm.values.title && noteForm.values.content) {
        const timeout = setTimeout(function () {
          updateNote(noteForm.values);
        }, 3000);
        return function () {
          return clearTimeout(timeout);
        };
      }
    },
    [noteForm.values, updateNote]
  );

  return (
    <Box className="p-5">
      <TextInput
        size="xl"
        placeholder="Title"
        className="border-none outline-none"
        classNames={{
          input: "border-none outline-none text-5xl",
        }}
        unstyled
        {...noteForm.getInputProps("title")}
      />

      <Textarea
        className="mt-5 leading-loose"
        placeholder="Note content..."
        unstyled
        classNames={{
          input:
            "border-none outline-none text-xl resize-none w-full min-h-[60vh] font-normal",
        }}
        {...noteForm.getInputProps("content")}
      />

      {updateNoteLoading && (
        <p className="text-red-700 font-semibold">Updating...</p>
      )}
    </Box>
  );
}
