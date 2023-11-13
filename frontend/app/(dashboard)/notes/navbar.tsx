"use client";
import ProfileAvatar from "@/components/profile-avatar";
import { Button, Group, Text } from "@mantine/core";
import { ArrowLeft } from "iconsax-react";
import Link from "next/link";

import ArchiveIcon from "@/public/assets/icons/archive.svg";
import BinIcon from "@/public/assets/icons/bin.svg";
import StarIcon from "@/public/assets/icons/star.svg";
import { modals } from "@mantine/modals";
import { useDeleteNote } from "@/api/hooks/notes";
import { Note } from "@/schema";
import { revalidateHomepage } from "@/app/actions";

export function Navbar({ note }: { note?: Note | undefined }) {
  const { mutate: deleteNote, isLoading: deleteNoteLoading } = useDeleteNote(
    note?.uuid
  );

  function confirmDelete() {
    modals.openConfirmModal({
      title: "Delete note?",
      centered: true,
      children: (
        <Text size="sm">
          Are you sure you want to delete this note? This action is not
          reversible.
        </Text>
      ),
      labels: { confirm: "Delete note", cancel: "No, cancel" },
      confirmProps: { color: "red", loading: deleteNoteLoading },
      closeOnConfirm: false,
      onConfirm: function () {
        deleteNote();
        revalidateHomepage();
      },
    });
  }

  function confirmArchive() {
    modals.openConfirmModal({
      title: "Archive note?",
      centered: true,
      children: (
        <Text size="sm">
          Are you sure you want to archive this note? You can unarchive it later
        </Text>
      ),
      labels: { confirm: "Archive note", cancel: "No, cancel" },
      confirmProps: { loading: deleteNoteLoading },
      closeOnConfirm: false,
      onConfirm: function () {},
    });
  }

  return (
    <div className="p-4 border-b">
      <Group justify="space-between">
        <Link href="/dashboard">
          <Group gap="sm">
            <ArrowLeft />
            <Text>Back</Text>
          </Group>
        </Link>

        <Group align="center">
          {!!note && (
            <>
              <span className="cursor-pointer">
                <StarIcon />
              </span>
              <span className="cursor-pointer" onClick={confirmArchive}>
                <ArchiveIcon />
              </span>
              <span
                className="text-red-600 cursor-pointer"
                onClick={confirmDelete}
              >
                <BinIcon />
              </span>
            </>
          )}
          <Button>Share</Button>
          <ProfileAvatar />
        </Group>
      </Group>
    </div>
  );
}
