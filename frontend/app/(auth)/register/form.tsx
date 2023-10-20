"use client";
import { useRegister } from "@/api/hooks/auth";
import {
  Button,
  TextInput,
  PasswordInput,
  LoadingOverlay,
} from "@mantine/core";
import {
  useForm,
  isEmail,
  matchesField,
  isNotEmpty,
  zodResolver,
} from "@mantine/form";
import { z } from "zod";

const signupValidator = z
  .object({
    firstName: z.string().min(1, "Enter first name"),
    lastName: z.string().min(1, "Enter last name"),
    email: z.string().min(1, "Email is required").email("Enter valid email"),
    password: z.string().min(1, "Enter password"),
    confirmPassword: z.string().min(1, "Enter password"),
  })
  .refine(({ password, confirmPassword }) => password === confirmPassword, {
    message: "Password do not match",
    path: ["confirmPassword"],
  });

export type TRegisterForm = z.infer<typeof signupValidator>;

export default function RegisterForm() {
  const { mutate: registerUser, isLoading } = useRegister();

  const registerForm = useForm({
    initialValues: {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
    validate: zodResolver(signupValidator),
  });

  function handleSubmit({
    firstName,
    lastName,
    email,
    password,
  }: TRegisterForm) {
    const payload = {
      firstName,
      lastName,
      email,
      password,
    };

    registerUser(payload);
  }

  return (
    <form
      className="flex flex-col gap-4 relative"
      onSubmit={registerForm.onSubmit(handleSubmit)}
    >
      <LoadingOverlay visible={isLoading} />
      <TextInput
        label="First name"
        labelProps={{ className: "mb-2" }}
        placeholder=""
        size="md"
        {...registerForm.getInputProps("firstName")}
      />
      <TextInput
        label="Last name"
        labelProps={{ className: "mb-2" }}
        placeholder=""
        size="md"
        {...registerForm.getInputProps("lastName")}
      />
      <TextInput
        label="Email"
        labelProps={{ className: "mb-2" }}
        placeholder=""
        size="md"
        {...registerForm.getInputProps("email")}
      />
      <PasswordInput
        label="Password"
        placeholder="*******"
        size="md"
        {...registerForm.getInputProps("password")}
      />
      <PasswordInput
        label="Confirm Password"
        placeholder="*******"
        size="md"
        {...registerForm.getInputProps("confirmPassword")}
      />

      <Button
        size="md"
        type="submit"
        className="bg-primary-01 hover:bg-primary-01"
        loading={isLoading}
      >
        Sign Up
      </Button>
    </form>
  );
}
