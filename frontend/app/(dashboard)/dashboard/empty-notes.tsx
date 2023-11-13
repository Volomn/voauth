"use client";
import { Button, Stack, Text } from "@mantine/core";
import EmptyNotesIcon from "@/public/assets/icons/empty-notes.svg";
import { Add } from "iconsax-react";
import { useCreateEmptyNote } from "@/api/hooks/notes";

export function EmptyNotes() {
  const { createEmptyNote, isLoading } = useCreateEmptyNote();

  return (
    <section className="flex-grow flex justify-center items-center gap-4">
      <Stack gap="md" align="center">
        <EmptyNotesIcon />
        <Text className="font-secondary">No items</Text>
        <Button
          size="lg"
          leftSection={<Add size="32" color="#FFF" />}
          onClick={createEmptyNote}
          loading={isLoading}
        >
          Create note
        </Button>
      </Stack>
    </section>
  );
}
