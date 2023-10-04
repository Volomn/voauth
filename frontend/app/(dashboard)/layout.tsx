import { AppShell, AppShellMain, AppShellNavbar, Text } from "@mantine/core";
import NavItems from "./nav-items";
import { ReactNode } from "react";

export default function DashboardLayout({ children }: { children: ReactNode }) {
  return (
    <AppShell navbar={{ width: 300, breakpoint: "xs" }}>
      <AppShellNavbar p="xl" bg="brand.7">
        <Text fz="32px" fw={600} c="white">
          Voauth
        </Text>

        <NavItems />
      </AppShellNavbar>
      <AppShellMain>{children}</AppShellMain>
    </AppShell>
  );
}
