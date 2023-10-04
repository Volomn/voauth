"use client";
import { Avatar, Box, Group, TextInput } from "@mantine/core";
export function Navbar() {
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

        <Box pl={8} className="border-l">
          <Avatar size="md" />
        </Box>
      </Group>
    </div>
  );
}
