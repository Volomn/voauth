import { Button, Group, PasswordInput, Text, TextInput } from "@mantine/core";
import Link from "next/link";
import RegisterForm from "./form";

export default function Register() {
  return (
    <>
      <Text size="20px" fw={500}>
        Create your account
      </Text>

      <RegisterForm />

      <Group gap="xs">
        <Text>Have an account?</Text>
        <Link href="/login" className="text-[#656acb] font-medium">
          Sign in
        </Link>
      </Group>
    </>
  );
}
