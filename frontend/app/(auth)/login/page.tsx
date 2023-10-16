import { Group, Text } from "@mantine/core";
import Link from "next/link";
import { LoginForm } from "./form";

export default function Login() {
  return (
    <>
      <Text size="20px" fw={500}>
        Sign in
      </Text>
      <LoginForm />
      <Group gap="xs">
        <Text>No account?</Text>
        <Link href="/register" className="text-[#656acb] font-medium">
          Sign up
        </Link>
      </Group>
    </>
  );
}
