"use client";
import { useCreateEmptyNote } from "@/api/hooks/notes";
import ProfileAvatar from "@/components/profile-avatar";
import { Button, Group, TextInput } from "@mantine/core";

export function Navbar() {
  const { createEmptyNote, isLoading } = useCreateEmptyNote();

  return (
    <div className="p-4 border-b">
      <Group justify="space-between">
        <TextInput
          radius="lg"
          size="md"
          placeholder="Search for note"
          w="100%"
          maw="600px"
        />

        <Group>
          <Button
            variant="transparent"
            onClick={createEmptyNote}
            loading={isLoading}
          >
            New note
          </Button>
          <ProfileAvatar />
        </Group>
      </Group>
    </div>
  );
}
