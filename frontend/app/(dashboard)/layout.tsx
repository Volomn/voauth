import { AppShell, AppShellMain, AppShellNavbar, Text } from "@mantine/core";
import NavItems from "./nav-items";
import { ReactNode } from "react";
import Image from "next/image";
import NavLogo from "@/public/assets/icons/logo-white.png";

export default function DashboardLayout({ children }: { children: ReactNode }) {
  return (
    <AppShell navbar={{ width: 300, breakpoint: "xs" }}>
      <AppShellNavbar p="xl" bg="brand.7">
        <Image src={NavLogo} width={97} height={22} alt="Voauth logo" />

        <NavItems />
      </AppShellNavbar>
      <AppShellMain>{children}</AppShellMain>
    </AppShell>
  );
}
