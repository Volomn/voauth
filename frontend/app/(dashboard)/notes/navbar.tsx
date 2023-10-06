"use client";
import ProfileAvatar from "@/components/profile-avatar";
import { Button, Group, Text } from "@mantine/core";
import { ArrowLeft } from "iconsax-react";
import Link from "next/link";

import ArchiveIcon from "@/public/assets/icons/archive.svg";
import BinIcon from "@/public/assets/icons/bin.svg";
import StarIcon from "@/public/assets/icons/star.svg";

export function Navbar() {
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
          <span className="cursor-pointer">
            <StarIcon />
          </span>
          <span className="cursor-pointer">
            <ArchiveIcon />
          </span>
          <span className="text-red-600 cursor-pointer">
            <BinIcon />
          </span>
          <Button>Share</Button>
          <ProfileAvatar />
        </Group>
      </Group>
    </div>
  );
}
