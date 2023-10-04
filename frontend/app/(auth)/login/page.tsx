import { Button, Group, PasswordInput, Text, TextInput } from "@mantine/core";
import Link from "next/link";
import AuthLayout from "../layout";

export default function Login() {
  return (
    <>
      <Text size="20px" fw={500}>
        Sign in
      </Text>

      <form className="flex flex-col gap-6">
        <TextInput
          label="Email"
          labelProps={{ className: "mb-2" }}
          placeholder="example@example.com"
          size="lg"
        />
        <PasswordInput label="Password" placeholder="*******" size="lg" />
        <Button size="lg">Sign In</Button>
      </form>

      <Group gap="xs">
        <Text>No account?</Text>
        <Link href="/register" className="text-[#656acb] font-medium">
          Sign up
        </Link>
      </Group>
    </>
  );
}
