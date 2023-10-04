// Import styles of packages that you've installed.
// All packages except `@mantine/hooks` require styles imports
import "./globals.css";
import "@mantine/core/styles.css";

import { MantineProvider, ColorSchemeScript, NavLink } from "@mantine/core";

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
        <link rel="icon" href="/favicon.ico" sizes="any" />
      </head>
      <body>
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
          {children}
        </MantineProvider>
      </body>
    </html>
  );
}
