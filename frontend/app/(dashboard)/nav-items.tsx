"use client";
import { Button, NavLink, Stack, Text } from "@mantine/core";
import { modals } from "@mantine/modals";
import { usePathname, useRouter } from "next/navigation";
import { signOut } from "next-auth/react";
import classes from "./nav-items.module.css";

import NotesIcon from "@/public/assets/icons/note.svg";
import ArchiveIcon from "@/public/assets/icons/archive.svg";
import BinIcon from "@/public/assets/icons/bin.svg";
import StarIcon from "@/public/assets/icons/star.svg";

export default function NavItems() {
  const pathname = usePathname();
  const router = useRouter();
  function handleLogout() {
    return modals.openConfirmModal({
      title: "Please confirm your action",
      children: <Text size="sm">Are you sure you want to logout?</Text>,
      labels: { confirm: "Confirm", cancel: "Cancel" },
      onCancel: () => console.log("Cancel"),
      onConfirm: () => {
        signOut();
        router.push("/");
      },
    });
  }

  return (
    <Stack mt="xl" className="h-full">
      <NavLink
        label="All Notes"
        active={pathname === "/dashboard"}
        onClick={() => router.push("/dashboard")}
        leftSection={<NotesIcon />}
        classNames={classes}
      />
      <NavLink
        label="Favorite"
        active={pathname === "/favourites"}
        onClick={() => router.push("/favourites")}
        leftSection={<StarIcon />}
        classNames={classes}
      />
      <NavLink
        label="Archived"
        active={pathname === "/archive"}
        onClick={() => router.push("/archive")}
        leftSection={<ArchiveIcon />}
        classNames={classes}
      />
      <NavLink
        label="Bin"
        active={pathname === "/bin"}
        onClick={() => router.push("/bin")}
        leftSection={<BinIcon />}
        classNames={classes}
      />

      <Button mt="auto" onClick={handleLogout}>
        Logout
      </Button>
    </Stack>
  );
}
