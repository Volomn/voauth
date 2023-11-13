"use client";
import { SessionProvider } from "next-auth/react";
import { ReactNode } from "react";
import ReactQueryProvider from "./react-query";
import { MantineProvider } from "@mantine/core";
import { ModalsProvider } from "@mantine/modals";
import { Notifications } from "@mantine/notifications";

export default function Providers({ children }: { children: ReactNode }) {
  return (
    <MantineProvider
      theme={{
        primaryColor: "brand",
        colors: {
          brand: [
            "#eff0fb",
            "#dbdcf0",
            "#b3b6e3",
            "#898dd6",
            "#656acb",
            "#4f54c4",
            "#1B2063",
            "#151957",
            "#0c0f3d",
            "#080a31",
          ],
        },
      }}
    >
      <ModalsProvider>
        <Notifications position="top-right" />
        <SessionProvider>
          <ReactQueryProvider>{children}</ReactQueryProvider>
        </SessionProvider>
      </ModalsProvider>
    </MantineProvider>
  );
}
