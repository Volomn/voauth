import { useMutation, useQuery } from "react-query";
import { _axiosInstance, axiosInstance } from "..";
import { useSession } from "next-auth/react";
import axios, { AxiosResponse } from "axios";
import { Note } from "@/schema";
import { NoteSchema } from "@/app/(dashboard)/notes/[id]/form";
import { useRouter } from "next/navigation";
import { showNotification } from "@mantine/notifications";
import { modals } from "@mantine/modals";
import { revalidateHomepage } from "@/app/actions";

function useAxiosInstance() {
  const { data: session } = useSession();

  return axios.create({
    baseURL: process.env.NEXT_PUBLIC_APP_BASE_URL,
    headers: {
      Authorization: `Bearer ${session?.user.accessToken}`,
    },
  });
}

export function useFetchNotes() {
  const { data: session } = useSession();
  const axiosInstance = useAxiosInstance();
  return useQuery({
    queryKey: ["notes"],
    queryFn: function (): Promise<AxiosResponse<Note[]>> {
      return axiosInstance.get("/api/notes/");
    },
    enabled: !!session?.user.accessToken,
  });
}

export function useCreateEmptyNote() {
  const { mutateAsync: createNewNote, isLoading } = useAddNote();

  async function createNote() {
    try {
      await createNewNote({
        title: "New Note",
        content: "Enter content....",
        isPublic: true,
      });
      showNotification({ message: "New note created", color: "blue" });
      revalidateHomepage();
    } catch (error) {
      showNotification({ message: "Unable to create note", color: "red" });
    }
  }

  return {
    createEmptyNote: createNote,
    isLoading: isLoading,
  };
}

export function useAddNote() {
  const axiosInstance = useAxiosInstance();

  return useMutation({
    mutationFn: function (payload: NoteSchema): Promise<AxiosResponse<Note>> {
      return axiosInstance.post("/api/notes/", payload);
    },
    onSuccess: function (response) {
      const uuid = response.data.uuid;
      console.log({ uuid });
    },
  });
}

export function useUpdateNote(id: string) {
  const axiosInstance = useAxiosInstance();

  return useMutation({
    mutationFn: function ({ title, content }: NoteSchema) {
      return axiosInstance.put(`/api/notes/${id}`, { title, content });
    },
  });
}

export function useDeleteNote(id: string | undefined) {
  const axiosInstance = useAxiosInstance();
  const router = useRouter();

  return useMutation({
    mutationFn: function () {
      return axiosInstance.delete(`/api/notes/${id}`);
    },
    onSuccess: function () {
      showNotification({
        message: "Note deleted",
        color: "green",
      });
      router.push("/dashboard");
    },
    onError: function () {
      showNotification({
        title: "Operation failed",
        message: "Unable to delete note",
        color: "red",
      });
    },
    onSettled: function () {
      modals.closeAll();
    },
  });
}
