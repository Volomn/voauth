"use client";
import { useLogin } from "@/api/hooks/auth";
import { Button, PasswordInput, TextInput } from "@mantine/core";
import { zodResolver, useForm } from "@mantine/form";
import { signIn } from "next-auth/react";
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
    // signIn("credentials", {
    //   ...values,
    //   redirect: false,
    //   callbackUrl: "/dashboard",
    // });
  }

  return (
    <form
      className="flex flex-col gap-4"
      onSubmit={registerForm.onSubmit(handleSubmit)}
    >
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
      <Button size="md" type="submit" loading={isLoading}>
        Sign In
      </Button>
    </form>
  );
}
