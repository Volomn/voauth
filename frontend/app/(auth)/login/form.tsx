"use client";
import { useLogin } from "@/api/hooks/auth";
import {
  Button,
  LoadingOverlay,
  PasswordInput,
  TextInput,
} from "@mantine/core";
import { zodResolver, useForm } from "@mantine/form";
import { z } from "zod";

const loginValidator = z.object({
  email: z.string().min(1, "Email is required").email("Enter valid email"),
  password: z.string().min(1, "Enter password"),
});
export type TLoginForm = z.infer<typeof loginValidator>;

export function LoginForm() {
  const { mutate: login, isLoading } = useLogin();

  const registerForm = useForm({
    initialValues: {
      email: "",
      password: "",
    },
    validate: zodResolver(loginValidator),
  });

  function handleSubmit(values: TLoginForm) {
    login(values);
  }

  return (
    <form
      className="flex flex-col gap-4 relative"
      onSubmit={registerForm.onSubmit(handleSubmit)}
    >
      <LoadingOverlay visible={isLoading} />
      <TextInput
        label="Email"
        labelProps={{ className: "mb-2" }}
        placeholder="example@example.com"
        size="md"
        {...registerForm.getInputProps("email")}
      />
      <PasswordInput
        label="Password"
        placeholder="*******"
        size="md"
        {...registerForm.getInputProps("password")}
      />
      <Button
        size="md"
        type="submit"
        className="bg-primary-01 hover:bg-primary-01"
        loading={isLoading}
      >
        Sign In
      </Button>
    </form>
  );
}
