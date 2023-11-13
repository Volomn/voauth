// Import styles of packages that you've installed.
// All packages except `@mantine/hooks` require styles imports
import "./globals.css";

import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";
import { ColorSchemeScript } from "@mantine/core";

import { montserrat, sora } from "@/fonts/fonts";
import Providers from "@/app/providers/providers";

export const metadata = {
  title: "Voauth",
  description: "A guide to implementing an oauth2 compliant platform",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
        <link rel="icon" href="/favicon.svg" sizes="any" />
      </head>
      <body
        className={`${montserrat.variable} ${sora.variable} ${montserrat.className}`}
      >
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
