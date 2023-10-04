"use client";
import { Button, Stack, Text } from "@mantine/core";
import EmptyNotesIcon from "@/public/assets/icons/empty-notes.svg";
import { Add } from "iconsax-react";

export function EmptyNotes() {
  return (
    <section className="flex-grow flex justify-center items-center gap-4">
      <Stack gap="md" align="center">
        <EmptyNotesIcon />
        <Text className="font-secondary">No items</Text>
        <Button size="lg" leftSection={<Add size="32" color="#FFF" />}>
          Create note
        </Button>
      </Stack>
    </section>
  );
}
