import { Button, Group, PasswordInput, Text, TextInput } from "@mantine/core";
import AuthLayout from "../layout";
import Link from "next/link";

export default function Register() {
  return (
    <>
      <Text size="20px" fw={500}>
        Create your account
      </Text>

      <form className="flex flex-col gap-6">
        <TextInput
          label="Email"
          labelProps={{ className: "mb-2" }}
          placeholder="example@example.com"
          size="lg"
        />
        <PasswordInput label="Password" placeholder="*******" size="lg" />
        <PasswordInput
          label="Confirm Password"
          placeholder="*******"
          size="lg"
        />

        <Button size="lg">Sign Up</Button>
      </form>

      <Group gap="xs">
        <Text>Have an account?</Text>
        <Link href="/login" className="text-[#656acb] font-medium">
          Sign in
        </Link>
      </Group>
    </>
  );
}
