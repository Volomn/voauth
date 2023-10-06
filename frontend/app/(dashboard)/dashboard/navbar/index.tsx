import ProfileAvatar from "@/components/profile-avatar";
import { Group, TextInput } from "@mantine/core";
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

        <ProfileAvatar />
      </Group>
    </div>
  );
}
