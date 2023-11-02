import { showNotification } from "@mantine/notifications";
import { AxiosResponse, AxiosError } from "axios";
import { useMutation } from "react-query";
import { axiosInstance } from "..";
import { TRegisterForm } from "@/app/(auth)/register/form";
import { TLoginForm } from "@/app/(auth)/login/form";
import { useRouter } from "next/navigation";

export function useRegister() {
  return useMutation({
    mutationFn: (payload: Omit<TRegisterForm, "confirmPassword">) => {
      return axiosInstance.post("/users/", payload);
    },
    onSuccess: (response: AxiosResponse) => {
      if (response.status === 201) {
        showNotification({
          message: response.data.msg || "Account created successfully",
          color: "green",
        });
      }
    },
    onError: (error: AxiosError) => {
      console.log({ error });
      const responseText = error.response as { data: { msg: string } };
      showNotification({
        title: "Unable to register",
        message: responseText.data.msg || "Sign up failed",
        color: "red",
      });
    },
  });
}

export function useLogin() {
  const router = useRouter();
  return useMutation({
    mutationFn: (payload: TLoginForm) => {
      return axiosInstance.post("/auth/", payload);
    },
    onSuccess: (response: AxiosResponse) => {
      if (response.status === 200) {
        showNotification({
          message: response.data.msg || "Login successful",
          color: "green",
        });

        router.push("/dashboard");
      }
    },
    onError: (error: AxiosError) => {
      const responseText = error.response as { data: { msg: string } };
      showNotification({
        message: responseText.data.msg || "Unable to login",
        color: "red",
      });
    },
  });
}
